package guidefetch

import (
	"errors"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func generateCommandOptions(g []Guide) ([]*discordgo.ApplicationCommandOption, error) {
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

func GetCommands() ([]*discordgo.ApplicationCommand, error) {
	commandOptions, err := generateCommandOptions(guides)
	if err != nil {
		return nil, fmt.Errorf("unable to generated command options: %v", err)
	}
	return []*discordgo.ApplicationCommand{
		//https://discord.com/developers/docs/interactions/application-commands#slash-commands
		{
			Name:        "fetch-guide",
			Description: "Provides a link to materials for a given Destiny activity",
			Options:     commandOptions,
		},
	}, nil
}

func GetCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {

	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"fetch-guide": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			content := ""
			switch i.ApplicationCommandData().Options[0].Name {
			case crypt.SubCommandName:
				content = guideMessage(i, crypt.Name, crypt.GDriveUrl)
			case garden.SubCommandName:
				content = guideMessage(i, garden.Name, garden.GDriveUrl)
			case kingsfall.SubCommandName:
				content = guideMessage(i, kingsfall.Name, kingsfall.GDriveUrl)
			case pit.SubCommandName:
				content = guideMessage(i, pit.Name, pit.GDriveUrl)
			case ron.SubCommandName:
				content = guideMessage(i, ron.Name, ron.GDriveUrl)
			case spire.SubCommandName:
				content = guideMessage(i, spire.Name, spire.GDriveUrl)
			case vault.SubCommandName:
				content = guideMessage(i, vault.Name, vault.GDriveUrl)
			case vow.SubCommandName:
				content = guideGithub(i, vow.Name, vow.GHUrl, vow.GDriveUrl)
			case wish.SubCommandName:
				content = guideMessage(i, wish.Name, wish.GDriveUrl)
			default:
				content = "Oops, something has gone wrong!\n"
				log.Printf("fetch-guide has ran into `default` in switch statement! Value: %v", i.ApplicationCommandData().Options[0].Name)
			}
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: content,
				},
			})
			if err != nil {
				log.Printf("Failed to respond to interaction: %v", err)
			}
		},
	}
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
