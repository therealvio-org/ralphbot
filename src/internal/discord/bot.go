package discord

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"ralphbot/internal/command/dadjoke"
	"ralphbot/internal/command/guidefetch"
	"ralphbot/internal/command/linkdump"
	"ralphbot/internal/config"

	"github.com/bwmarrin/discordgo"
)

// Starts the `ralphbot` service
func StartBotService(s *discordgo.Session, env *config.EnvConfig) {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer s.Close()

	_, err = registerCommand(s, env.GuildID, guidefetch.Commands, guidefetch.CommandHandlers)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	_, err = registerCommand(s, env.GuildID, dadjoke.Commands, dadjoke.CommandHandlers)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	_, err = registerCommand(s, env.GuildID, linkdump.Commands, linkdump.CommandHandlers)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C/Cmd+C to exit")
	<-stop

	if env.RemoveCommands {
		deregisterCommand(s, env)
	}

	log.Println("Shutting down gracefully...")
}
