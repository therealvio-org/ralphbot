package discord

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

//Registers commands for `ralphbot` service.
//If `GUILD_ID` is passed, then the command is set as a guild command, and will not be registered globally. However, it will be immediately registered exclusively
//to the Discord server (guild). If `GUILD_ID` is not passed, then the command is registered globally.
func RegisterCommand(s *discordgo.Session, id string, commands []*discordgo.ApplicationCommand, handler map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) ([]*discordgo.ApplicationCommand, error) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := handler[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, id, v)
		if err != nil {
			return nil, fmt.Errorf("cannot create command: '%v', error: %v", v.Name, err)
		}
		log.Printf("Registered command: %v", v.Name)
		registeredCommands[i] = cmd
	}

	return registeredCommands, nil
}
