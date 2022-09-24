package dadjoke

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		//https://discord.com/developers/docs/interactions/application-commands#slash-commands
		{
			Name:        "dad-joke",
			Description: "Does this really need a description?",
		},
	}

	jokes = getJokes()

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

type StructJokes struct {
	Jokes []string `json:"jokes"`
}

func getJokes() []string {
	jokesFile, err := os.ReadFile("jokes.json")
	if err != nil {
		log.Fatal("Unable to read jokesFile!")
	}

	jokeArray := StructJokes{}

	err = json.Unmarshal([]byte(jokesFile), &jokeArray)
	if err != nil {
		log.Printf("Unable to Unmarshal : %v", err)
	}

	return jokeArray.Jokes

}

func dadJoke(i *discordgo.InteractionCreate, j []string) string {
	rand.Seed(time.Now().Unix())
	selectedDadJoke := j[rand.Intn(len(j))]
	result := string(selectedDadJoke)
	return result
}
