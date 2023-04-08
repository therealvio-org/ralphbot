package coinflip

import (
	"bytes"
	"log"
	"math/rand"
	"text/template"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	sides = []string{"**Heads**", "**Tails**"}

	phrases = []string{
		"It's **{{.Side}}**",
		"What the **{{.Side}}** doing?",
		"**{{.Side}}**, you hate to see it",
	}
)

func coinFlip() string {
	selectedSide := sides[rand.Intn(len(sides))]
	return selectedSide
}

func selectPhrase(phrases []string) string {
	rand.NewSource(time.Now().UnixNano())
	selectedPhrase := phrases[rand.Intn(len(phrases))]
	return selectedPhrase
}

func makePhrase(side string, phrase string) string {
	type coin struct {
		Side string
	}

	c := &coin{}
	c.Side = side
	buf := new(bytes.Buffer)
	template, _ := template.New("phrase").Parse(phrase)
	template.Execute(buf, c)

	return buf.String()
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
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"coin-flip": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: makePhrase(coinFlip(), selectPhrase(phrases)),
				},
			})
			if err != nil {
				log.Printf("Failed to respond to interaction: %v", err)
			}
		},
	}, nil
}
