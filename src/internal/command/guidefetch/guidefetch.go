package guidefetch

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		//https://discord.com/developers/docs/interactions/application-commands#slash-commands
		{
			Name:        "fetch-guide",
			Description: "Provides a link to materials for a given Destiny activity",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "raid-kingsfall",
					Description: "King's Fall Raid",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
				{
					Name:        "raid-vow",
					Description: "Vow of the Disciple Raid",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
				{
					Name:        "raid-vault",
					Description: "Vault of Glass Raid",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
				{
					Name:        "raid-crypt",
					Description: "Deep Stone Crypt Raid",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
				{
					Name:        "raid-garden",
					Description: "Garden of Salvation Raid",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
				{
					Name:        "raid-lastwish",
					Description: "Last Wish Raid",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
				{
					Name:        "dungeon-pit",
					Description: "Pit of Heresy Dungeon",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
			},
		},
	}

	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"fetch-guide": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			content := ""
			switch i.ApplicationCommandData().Options[0].Name {
			case "raid-kingsfall":
				content = guideMessage(i, "Kings Fall", "https://drive.google.com/drive/folders/1tsOVCy2SwP0rLUDQUJaDIFh5y-O0DoKn")
			case "raid-vow":
				content = guideGithub(i, "Vow of the Disciple", "https://github.com/therealvio/destiny-guides/tree/main/raids/vow-of-the-disciple", "https://drive.google.com/drive/folders/1ZAPIXYlSs7yTQEdznQAqz2rOnnvpzwr7?usp=sharing")
			case "raid-vault":
				content = guideMessage(i, "Vault of Glass", "https://drive.google.com/drive/folders/1HLx6nVIji_3OcwnzaLeSoksspa4pfdjD?usp=sharing")
			case "raid-crypt":
				content = guideMessage(i, "Deep Stone Crypt", "https://drive.google.com/drive/folders/1YKU4_-hInHQ3rVEAvqIjdJaT25oQvmYc?usp=sharing")
			case "raid-garden":
				content = guideMessage(i, "Garden of Salvation", "https://drive.google.com/drive/folders/1pPdtAptJMaaDYRv2i-8bfaL6l3I0WTsT?usp=sharing")
			case "raid-lastwish":
				content = guideMessage(i, "Last Wish", "https://drive.google.com/drive/folders/1d_WEa84KuX1_9hPTwgFhl651IwywHeOg?usp=sharing")
			case "dungeon-pit":
				content = guideMessage(i, "Pit of Heresy", "https://drive.google.com/drive/folders/17lB7m9KQMwzBb6UHfoBt9ZEA82haD2Fd?usp=sharing")
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

func guideMessage(i *discordgo.InteractionCreate, activity string, link string) string {
	result := fmt.Sprintf("%s, here is your requested %s supplementary material!\n %s", i.Member.Mention(), activity, link)
	return result
}

/**
Testing Github links in tandem with Google Drive links
Arrow brackets are used to escape the github link to prevent previews
*/

func guideGithub(i *discordgo.InteractionCreate, activity string, ghubLink string, gdriveLink string) string {
	result := fmt.Sprintf("%s, here is your requested %s supplementary material!\n\n Github Link: <%s>\n\nGoogle DriveLink: %s", i.Member.Mention(), activity, ghubLink, gdriveLink)
	return result
}
