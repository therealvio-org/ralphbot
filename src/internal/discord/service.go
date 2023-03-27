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

type DiscordSession struct {
	*discordgo.Session
}

type DiscordHandlerOutput struct {
	handler func()
}

// Starts a new Discord session
func NewDiscord(authToken string) (*DiscordSession, error) {
	s, err := discordgo.New("Bot " + authToken)

	if err != nil {
		return nil, fmt.Errorf("error in creating discord session: %v", err)
	}

	return &DiscordSession{s}, nil
}

func CheckOnline(ds *DiscordSession) *DiscordHandlerOutput {
	return &DiscordHandlerOutput{ds.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})}
}

// Starts the `ralphbot` service, to be used after pre-flight checks
// This should be responsible for the running service, command registration, e.t.c.
func StartBotService(ds *DiscordSession, env *config.EnvConfig) error {
	err := ds.Open()
	if err != nil {
		err = fmt.Errorf("cannot open the session: %v", err)
		return err
	}
	defer ds.Close()

	_, err = registerCommand(ds, env.GuildID, guidefetch.Commands, guidefetch.CommandHandlers)
	if err != nil {
		err = fmt.Errorf("unable to register command %v error: %v", "guidefetch", err)
		return err
	}
	_, err = registerCommand(ds, env.GuildID, dadjoke.Commands, dadjoke.CommandHandlers)
	if err != nil {
		err = fmt.Errorf("unable to register command %v error: %v", "dadjoke", err)
		return err
	}
	_, err = registerCommand(ds, env.GuildID, linkdump.Commands, linkdump.CommandHandlers)
	if err != nil {
		err = fmt.Errorf("unable to register command %v error: %v", "linkdump", err)
		return err
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C/Cmd+C to exit")
	<-stop

	if env.RemoveCommands {
		deregisterCommand(ds, env)
	}

	log.Println("Shutting down gracefully...")
	return nil
}
