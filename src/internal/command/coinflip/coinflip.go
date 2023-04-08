package coinflip

import (
	"log"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	sides = []string{"heads", "tails"}
)

func coinFlip() string {
	rand.NewSource(time.Now().UnixNano())
	selectedSide := sides[rand.Intn(len(sides))]
	return selectedSide
}

func GetCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		{
			Name:        "coin-flip",
			Description: "Let ralphbot choose for you",
		},
	}
}

func GetCommandHandlers() (map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate), error) {
	side := coinFlip()

	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"coin-flip": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: side,
				},
			})
			if err != nil {
				log.Printf("Failed to respond to interaction: %v", err)
			}
		},
	}, nil
}
