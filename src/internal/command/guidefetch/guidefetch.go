package guidefetch

import (
	"errors"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func GenerateCommandOptions(g []guide) ([]*discordgo.ApplicationCommandOption, error) {
	var subCommands []*discordgo.ApplicationCommandOption

	if g == nil {
		return nil, errors.New("guides slice is empty")
	}

	for _, s := range g {
		subCommands = append(subCommands, &discordgo.ApplicationCommandOption{
			Name:        s.SubCommandName,
			Description: s.Description,
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		})
	}

	return subCommands, nil
}

func GetCommands(co []*discordgo.ApplicationCommandOption) []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		//https://discord.com/developers/docs/interactions/application-commands#slash-commands
		{
			Name:        "fetch-guide",
			Description: "Provides a link to materials for a given Destiny activity",
			Options:     co,
		},
	}
}

func GetCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"fetch-guide": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			response := getInteractionResponse(i)
			err := s.InteractionRespond(i.Interaction, response)
			if err != nil {
				log.Printf("Failed to respond to interaction: %v", err)
			}
		},
	}
}

func getInteractionResponse(i *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "",
		},
	}

	for _, g := range Guides {
		if i.ApplicationCommandData().Options[0].Name == g.SubCommandName {
			if g.GHLink != "" {
				response.Data.Content = guideGithub(i, g.Name, g.GHLink, g.GDriveLink)
				return response
			}
			response.Data.Content = guideMessage(i, g.Name, g.GDriveLink)
			return response
		}
	}

	response.Data.Content = "Oops, your command didn't return a guide message!\n"
	log.Printf("fetch-guide was unablet to source guide: %v", i.ApplicationCommandData().Options[0].Name)
	return response
}

func guideMessage(i *discordgo.InteractionCreate, activity string, link string) string {
	result := fmt.Sprintf("%s, here is your requested **%s** supplementary material!\n\n[Google Drive Link](%s)", i.Member.Mention(), activity, link)
	return result
}

/**
Testing Github links in tandem with Google Drive links
Arrow brackets are used to escape the github link to prevent previews
*/

func guideGithub(i *discordgo.InteractionCreate, activity string, ghubLink string, gdriveLink string) string {
	result := fmt.Sprintf("%s, here is your requested **%s** supplementary material!\n\n[Github Link](<%s>)\n[Google Drive Link](%s)", i.Member.Mention(), activity, ghubLink, gdriveLink)
	return result
}
