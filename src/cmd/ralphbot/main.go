package main

import (
	"fmt"
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

func main() {
	discord.CheckGuildId(s, env.GuildID)

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	_, err = discord.RegisterCommands(s, env.GuildID, guidefetch.Commands, guidefetch.CommandHandlers)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	_, err = discord.RegisterCommands(s, env.GuildID, dadjoke.Commands, dadjoke.CommandHandlers)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

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
