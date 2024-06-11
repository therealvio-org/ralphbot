package guide

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
			Name:        "guide",
			Description: "Provides a link to materials for a given Destiny activity",
			Options:     co,
		},
	}
}

func GetCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"guide": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			resp := getInteractionResponse(i)
			err := s.InteractionRespond(i.Interaction, resp)
			if err != nil {
				log.Printf("Failed to respond to interaction: %v", err)
			}
		},
	}
}

func getInteractionResponse(i *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	if i.Type != discordgo.InteractionApplicationCommand {
		log.Printf("fetchguide - interaction type does not match: discordgo.InteractionApplicationCommand - Type: %v Data:%v", i.Type, i.Data)
		return nil
	}

	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "",
		},
	}

	for _, g := range Guides {
		if i.ApplicationCommandData().Options[0].Name == g.SubCommandName {
			response.Data.Content = guideMessage(i, &g)
			return response
		}
	}

	response.Data.Content = "Oops, your command didn't return a guide message!\n"
	log.Printf("fetch-guide was unable to source guide: %v", i.ApplicationCommandData().Options[0].Name)
	return response
}

func guideMessage(i *discordgo.InteractionCreate, g *guide) string {
	if g.GHLink != "" {
		result := fmt.Sprintf("%s, here is your requested **%s** supplementary material!\n\n[Github Link](<%s>)\n[Google Drive Link](%s)", i.Member.Mention(), g.Name, g.GHLink, g.GDriveLink)
		return result
	}
	result := fmt.Sprintf("%s, here is your requested **%s** supplementary material!\n\n[Google Drive Link](%s)", i.Member.Mention(), g.Name, g.GDriveLink)
	return result
}
