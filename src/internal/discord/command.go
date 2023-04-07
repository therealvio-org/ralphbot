package discord

import (
	"fmt"
	"log"

	"ralphbot/internal/config"

	"github.com/bwmarrin/discordgo"
)

// Registers commands for `ralphbot` service.
// If `GUILD_ID` is passed, then the command is set as a guild command, and will not be registered globally. However, it will be immediately registered exclusively
// to the Discord server (guild). If `GUILD_ID` is not passed, then the command is registered globally.
func registerCommand(ds *DiscordSession, gId string, commands []*discordgo.ApplicationCommand, handler map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) ([]*discordgo.ApplicationCommand, error) {
	ds.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := handler[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := ds.ApplicationCommandCreate(ds.State.User.ID, gId, v)
		if err != nil {
			return nil, fmt.Errorf("cannot create command: '%v', error: %v", v.Name, err)
		}
		log.Printf("Registered command: %v", v.Name)
		registeredCommands[i] = cmd
	}

	return registeredCommands, nil
}

func deregisterCommand(ds *DiscordSession, env *config.EnvConfig) {
	log.Println("Removing commands...")
	// We need to fetch the commands, since deleting requires the command ID.
	// We are doing this from commands defined in registerCommand() runs, because using
	// this will delete all the commands, which might not be desirable, so we
	// are deleting only the commands that we added.

	registeredCommands, err := ds.ApplicationCommands(ds.State.User.ID, env.GuildID)
	if err != nil {
		log.Fatalf("Could not fetch registered commands: %v", err)
	}

	for _, v := range registeredCommands {
		err := ds.ApplicationCommandDelete(ds.State.User.ID, env.GuildID, v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
}
