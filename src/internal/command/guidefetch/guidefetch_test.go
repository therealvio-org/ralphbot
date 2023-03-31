package guidefetch

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func TestGenerateCommandOptions(t *testing.T) {
	backrooms := &guide{
		Name:           "The Backrooms",
		SubCommandName: "raid-backrooms",
		Description:    "The Backrooms Raid",
		GDriveUrl:      "https://drive.google.com/drive/folders/123456abcdef?usp=sharing",
		GHUrl:          "https://github.com/butterdog/destiny-guides/tree/main/raids/the-backrooms",
	}

	wax := &guide{
		Name:           "Wax Museum",
		SubCommandName: "raid-wax",
		Description:    "The Wax Museum Raid",
		GDriveUrl:      "https://drive.google.com/drive/folders/1d_WEa84KuX1_9hPTwgFhl651IwywHeOg?usp=sharing",
	}

	cases := []struct {
		name     string
		input    []guide
		expected []*discordgo.ApplicationCommandOption
	}{
		{
			name:     "when no guides are supplied, return a slice of zero ApplicationCommandOption, and an error",
			input:    []guide{},
			expected: nil,
		},
		{
			name: "when one guide is supplied, return a slice of one ApplicationCommandOption",
			input: []guide{
				*backrooms,
			},
			expected: []*discordgo.ApplicationCommandOption{
				{
					Name:        backrooms.SubCommandName,
					Description: backrooms.Description,
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
			},
		},
		{
			name: "when two guides are supplied, return a slice of two ApplicationCommandOption",
			input: []guide{
				*backrooms,
				*wax,
			},
			expected: []*discordgo.ApplicationCommandOption{
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
			},
		},
	}

	for _, test := range cases {
		result, err := generateCommandOptions(test.input)
		assert.NoError(t, err)
		assert.Len(t, result, len(test.expected))
		assert.ElementsMatch(t, test.expected, result)
	}
}
