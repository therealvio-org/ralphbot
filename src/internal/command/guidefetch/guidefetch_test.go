package guidefetch

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func TestGenerateCommandOptions(t *testing.T) {

	backrooms := &Guide{
		Name:           "The Backrooms",
		SubCommandName: "raid-backrooms",
		Description:    "The Backrooms Raid",
		GDriveUrl:      "https://drive.google.com/drive/folders/123456abcdef?usp=sharing",
		GHUrl:          "https://github.com/butterdog/destiny-guides/tree/main/raids/the-backrooms",
	}

	wax := &Guide{
		Name:           "Wax Museum",
		SubCommandName: "raid-wax",
		Description:    "The Wax Museum Raid",
		GDriveUrl:      "https://drive.google.com/drive/folders/1d_WEa84KuX1_9hPTwgFhl651IwywHeOg?usp=sharing",
	}

	commandOptions = generateCommandOptions([]Guide{
		*backrooms,
		*wax,
	})

	expected := []*discordgo.ApplicationCommandOption{
		{
			Name:        backrooms.SubCommandName,
			Description: backrooms.Description,
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		},
		{
			Name:        wax.SubCommandName,
			Description: wax.Description,
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		},
	}

	assert.ElementsMatch(t, expected, commandOptions)
}
