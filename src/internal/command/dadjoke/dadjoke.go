package dadjoke

import (
	_ "embed"
	"encoding/json"
	"log"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

//go:embed jokes.json
var jokesFile []byte

var (
	Commands = []*discordgo.ApplicationCommand{
		//https://discord.com/developers/docs/interactions/application-commands#slash-commands
		{
			Name:        "dad-joke",
			Description: "Does this really need a description?",
		},
	}

	jokes = getJokes(jokesFile)

	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"dad-joke": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: dadJoke(i, jokes),
				},
			})
			if err != nil {
				log.Printf("Failed to respond to interaction: %v", err)
			}
		},
	}
)

type JokeStruct struct {
	Jokes []string `json:"jokes"`
}

func getJokes(b []byte) []string {

	var jokeArray JokeStruct
	err := json.Unmarshal(b, &jokeArray)
	if err != nil {
		log.Printf("Unable to Unmarshal : %v", err)
	}

	return jokeArray.Jokes

}

func dadJoke(i *discordgo.InteractionCreate, j []string) string {
	selectedDadJoke := j[rand.Intn(len(j))]
	result := string(selectedDadJoke)
	return result
}
