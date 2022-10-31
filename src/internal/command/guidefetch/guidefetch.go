package guidefetch

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	commandOptions = generateCommandOptions(guides)

	Commands = []*discordgo.ApplicationCommand{
		//https://discord.com/developers/docs/interactions/application-commands#slash-commands
		{
			Name:        "fetch-guide",
			Description: "Provides a link to materials for a given Destiny activity",
			Options:     commandOptions,
		},
	}

	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"fetch-guide": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			content := ""
			switch i.ApplicationCommandData().Options[0].Name {
			case kingsfall.SubCommandName:
				content = guideMessage(i, kingsfall.Name, kingsfall.GDriveUrl)
			case vow.SubCommandName:
				content = guideGithub(i, vow.Name, vow.GHUrl, vow.GDriveUrl)
			case vault.SubCommandName:
				content = guideMessage(i, vault.Name, vault.GDriveUrl)
			case crypt.SubCommandName:
				content = guideMessage(i, crypt.Name, crypt.GDriveUrl)
			case garden.SubCommandName:
				content = guideMessage(i, garden.Name, garden.GDriveUrl)
			case wish.SubCommandName:
				content = guideMessage(i, wish.Name, wish.GDriveUrl)
			case pit.SubCommandName:
				content = guideMessage(i, pit.Name, pit.GDriveUrl)
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
)

func generateCommandOptions(g []Guide) []*discordgo.ApplicationCommandOption {
	var subCommands []*discordgo.ApplicationCommandOption

	for _, s := range g {
		subCommands = append(subCommands, &discordgo.ApplicationCommandOption{
			Name:        s.SubCommandName,
			Description: s.Description,
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		})
	}

	return subCommands
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
	result := fmt.Sprintf("%s, here is your requested **%s** supplementary material!\n\n [Github Link](<%s>)\n[Google Drive Link](%s)", i.Member.Mention(), activity, ghubLink, gdriveLink)
	return result
}
