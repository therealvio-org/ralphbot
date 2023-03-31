package linkdump

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type LinkDetails struct {
	Name string
	URL  string
}

type Links struct {
	Links []LinkDetails `json:"Links"`
}

//go:embed links.json
var linkFile []byte

func getLinks(b []byte) (string, error) {
	var linkStruct Links
	err := json.Unmarshal(b, &linkStruct)
	if err != nil {
		log.Printf("Unable to Unmarshal : %v", err)
		return "", err
	}

	var linkSlice []string
	for _, l := range linkStruct.Links {
		//For some reason, markdown for links doesn't work here
		linkElem := fmt.Sprintf("%s - <%s>", l.Name, l.URL)
		linkSlice = append(linkSlice, linkElem)
	}
	content := strings.Join(linkSlice, "\n")

	return content, nil
}

func GetCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		{
			Name:        "link-dump",
			Description: "Sends a link dump via DM of helpful Destiny 2-related web apps",
		},
	}
}

func GetCommandHandlers() (map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate), error) {
	links, err := getLinks(linkFile)
	if err != nil {
		return nil, err
	}

	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"link-dump": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			channel, err := s.UserChannelCreate(i.Member.User.ID)
			if err != nil {
				log.Printf("Failed to create DM channel for %v: %v", i.User.ID, err)
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   discordgo.MessageFlagsEphemeral,
					Content: "Check your DMs for a message from me! :smirk:",
				},
			})
			if err != nil {
				log.Printf("Failed to respond to interaction: %v", err)
			}

			_, err = s.ChannelMessageSend(channel.ID, links)
			if err != nil {
				log.Printf("Failed to send message to DM channel: %v", err)
			}
		},
	}, nil
}
