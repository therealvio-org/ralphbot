package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/ralphbot/internal/config"
	"github.com/ralphbot/internal/dadjoke"
	"github.com/ralphbot/internal/discord"
	"github.com/ralphbot/internal/guidefetch"
)

var (
	env = config.New()
)

var s *discordgo.Session

func init() {
	var err error
	s, err = discordgo.New("Bot " + env.BotToken)
	if err != nil {
		log.Fatalf("Error creating new Discord session: %v", err)
	}
}

func registerCommands(commands []*discordgo.ApplicationCommand, handler map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) []*discordgo.ApplicationCommand {

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := handler[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, env.GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	return registeredCommands
}

func main() {
	discord.CheckGuildId(env.GuildID, s)

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Adding commands...")
	registerCommands(guidefetch.Commands, guidefetch.CommandHandlers)
	registerCommands(dadjoke.Commands, dadjoke.CommandHandlers)

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C/Cmd+C to exit")
	<-stop

	if env.RemoveCommands {
		log.Println("Removing commands...")
		// We need to fetch the commands, since deleting requires the command ID.
		// We are doing this from commands defined in registerCommand() runs, because using
		// this will delete all the commands, which might not be desirable, so we
		// are deleting only the commands that we added.

		registeredCommands, err := s.ApplicationCommands(s.State.User.ID, env.GuildID)
		if err != nil {
			log.Fatalf("Could not fetch registered commands: %v", err)
		}

		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, env.GuildID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	log.Println("Shutting down gracefully...")
}
