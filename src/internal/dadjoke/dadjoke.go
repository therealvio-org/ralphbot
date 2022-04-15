package dadjoke

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
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
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: dadJoke(i, jokes),
				},
			})
		},
	}
)

type StructJokes struct {
	Jokes []string `json:"jokes"`
}

func getJokes() []string {
	jokesFile, err := ioutil.ReadFile("jokes.json")

	if err != nil {
		log.Fatal("Unable to open jokesFile!")
	}

	jokeArray := StructJokes{}

	json.Unmarshal([]byte(jokesFile), &jokeArray)

	return jokeArray.Jokes //json.Unmarshal([]byte(jokesFile), StructJokes)

}

func dadJoke(i *discordgo.InteractionCreate, j []string) string {
	rand.Seed(time.Now().Unix())
	selectedDadJoke := j[rand.Intn(len(j))]
	result := string(selectedDadJoke)
	return result
}
