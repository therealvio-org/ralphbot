package dadjoke

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"

	"ralphbot/internal/common"

	"github.com/bwmarrin/discordgo"
)

//go:embed jokes.json
var jokesFile []byte

type JokeStruct struct {
	Jokes []string `json:"jokes"`
}

func getJokes(b []byte) ([]string, error) {

	var jokeArray JokeStruct
	err := json.Unmarshal(b, &jokeArray)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal json: %v", err)
	}
	return jokeArray.Jokes, nil
}

func GetCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		//https://discord.com/developers/docs/interactions/application-commands#slash-commands
		{
			Name:        "dad-joke",
			Description: "Does this really need a description?",
		},
	}
}

func GetCommandHandlers() (map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate), error) {
	jokes, err := getJokes(jokesFile)
	if err != nil {
		return nil, fmt.Errorf("unable to execute getJokes: %v", err)
	}

	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"dad-joke": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: common.SelectRandomString(jokes), //selectDadJoke(jokes),
				},
			})
			if err != nil {
				log.Printf("Failed to respond to interaction: %v", err)
			}
		},
	}, nil
}
