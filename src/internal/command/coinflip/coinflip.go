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
		"Can you believe it? **{{.Side}}**, only two weeks away!",
		"Now that **{{.Side}}** is gone",
		"**{{.Side}}**, HOW!?",
		"Uh oh, **{{.Side}}**",
		"**{{.Side}}**, I believe",
		"When's Guild **{{.Side}}** night?",
		"Wow, they have it! **{{.Side}}**",
		"Where's your 4 **{{.Side}}**?",
		"I can shoot four **{{.Side}}**",
		"I'm gonna **{{.Side}}**",
		"Jeffery Epstein didn't kill himself. Oh and I got **{{.Side}}**",
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
	// coin sides are always strings, though we need to use a struct to satisfy template execution
	type coin struct {
		Side string
	}

	c := &coin{}
	c.Side = side
	buf := new(bytes.Buffer)
	template, err := template.New("phrase").Parse(phrase)
	if err != nil {
		log.Printf("Failed to parse phrase template: %v", err)
	}
	err = template.Execute(buf, c)
	if err != nil {
		log.Printf("Failed to execute phrase template: %v", err)
	}

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
