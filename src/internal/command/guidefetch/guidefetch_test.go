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
		result, err := GenerateCommandOptions(test.input)
		assert.NoError(t, err)
		assert.Len(t, result, len(test.expected))
		assert.ElementsMatch(t, test.expected, result)
	}
}

func TestGetCommands(t *testing.T) {
	backrooms := &discordgo.ApplicationCommandOption{
		Name:        "The Backrooms",
		Description: "The Backrooms Raid",
		Type:        discordgo.ApplicationCommandOptionSubCommand,
	}

	wax := &discordgo.ApplicationCommandOption{
		Name:        "The Wax Museum",
		Description: "The Wax Museum Raid",
		Type:        discordgo.ApplicationCommandOptionSubCommand,
	}

	cases := []struct {
		name   string
		input  []*discordgo.ApplicationCommandOption
		expect []*discordgo.ApplicationCommand
	}{
		{
			name: "when one command option is provided, the command only has one option",
			input: []*discordgo.ApplicationCommandOption{
				backrooms,
			},
			expect: []*discordgo.ApplicationCommand{{
				Name:        "fetch-guide",
				Description: "Provides a link to materials for a given Destiny activity",
				Options: []*discordgo.ApplicationCommandOption{
					backrooms,
				},
			}},
		},
		{
			name: "when two command options are provided, the command has two options",
			input: []*discordgo.ApplicationCommandOption{
				backrooms,
				wax,
			},
			expect: []*discordgo.ApplicationCommand{{
				Name:        "fetch-guide",
				Description: "Provides a link to materials for a given Destiny activity",
				Options: []*discordgo.ApplicationCommandOption{
					backrooms,
					wax,
				},
			}},
		},
	}

	for _, test := range cases {
		result := GetCommands(test.input)
		assert.Equal(t, test.expect, result)
		assert.IsType(t, test.expect, result)
	}
}
