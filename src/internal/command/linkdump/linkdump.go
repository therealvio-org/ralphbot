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

type LinksStruct struct {
	Links []LinkDetails `json:"links"`
}

//go:embed links.json
var linkFile []byte

func getLinksFromJSON(b []byte) (LinksStruct, error) {
	var linkStruct LinksStruct
	err := json.Unmarshal(b, &linkStruct)
	if err != nil {
		log.Printf("Unable to Unmarshal : %v", err)
		return LinksStruct{}, err
	}

	return linkStruct, nil
}

func makeLinkDumpMessage(l LinksStruct) string {
	var linkSlice []string
	for _, l := range l.Links {
		//For some reason, markdown for links doesn't work here
		linkElem := fmt.Sprintf("%s - <%s>", l.Name, l.URL)
		linkSlice = append(linkSlice, linkElem)
	}
	content := strings.Join(linkSlice, "\n")

	return content
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
	links, err := getLinksFromJSON(linkFile)
	if err != nil {
		return nil, err
	}

	message := makeLinkDumpMessage(links)

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

			_, err = s.ChannelMessageSend(channel.ID, message)
			if err != nil {
				log.Printf("Failed to send message to DM channel: %v", err)
			}
		},
	}, nil
}
