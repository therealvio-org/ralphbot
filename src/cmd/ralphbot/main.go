package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	guidefetch "github.com/ralphbot/internal"
)

//Bot Parameters
var (
	GuildID        = flag.String("guild", os.Getenv("GUILD_ID"), "Test guild ID. If not passed - bot registers commands globally")
	BotToken       = flag.String("token", os.Getenv("BOT_TOKEN"), "Bot access token")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

var s *discordgo.Session

func init() {
	flag.Parse()
}

func init() {
	var err error
	s, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := guidefetch.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func checkGuildId(id string) {
	doFail := false
	if *GuildID != "" {
		log.Printf("Checking for GuildID variable validity against current server (if it matches, we're okay)...")

		g, err := s.Guild(id)
		if err != nil {
			log.Fatalf("Cannot retrieve Guild Id from server: %v", err)
		}

		if *GuildID == g.ID {
			log.Printf("GuildID is valid...")
			log.Printf("GuildID is defined - ralphbot is running in development mode, Guild commands are available...")
		} else {
			log.Printf("GuildID is invalid...")
			doFail = true
		}
	} else {
		doFail = true
	}
	if doFail {
		log.Printf("GuildID is undefined - ralphbot is running in production mode, only Global commands are available...")
	}
}

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	checkGuildId(*GuildID)

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(guidefetch.Commands))
	for i, v := range guidefetch.Commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, *GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C/Cmd+C to exit")
	<-stop

	if *RemoveCommands {
		log.Println("Removing commands...")
		// // We need to fetch the commands, since deleting requires the command ID.
		// // We are doing this from the returned commands on line 375, because using
		// // this will delete all the commands, which might not be desirable, so we
		// // are deleting only the commands that we added.
		// registeredCommands, err := s.ApplicationCommands(s.State.User.ID, *GuildID)
		// if err != nil {
		// 	log.Fatalf("Could not fetch registered commands: %v", err)
		// }

		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, *GuildID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	log.Println("Shutting down gracefully...")
}
