package discord

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

// NOTE: mockGuild is named after the method we're mocking, not a particular package
type mockGuild func(guildID string, options ...discordgo.RequestOption) (st *discordgo.Guild, err error)

func (m mockGuild) Guild(guildID string, options ...discordgo.RequestOption) (st *discordgo.Guild, err error) {
	return m(guildID, options...)
}

func TestCheckGuildId(t *testing.T) {
	testGuildId := "123456789"
	cases := []struct {
		name  string
		input struct {
			api func(t *testing.T) mockGuild
			id  string
		}
		expect bool
	}{
		{
			name: "when no guildID is provided, return false",
			input: struct {
				api func(t *testing.T) mockGuild
				id  string
			}{
				api: func(t *testing.T) mockGuild {
					return mockGuild(func(guildID string, options ...discordgo.RequestOption) (st *discordgo.Guild, err error) {
						return &discordgo.Guild{
							ID: testGuildId,
						}, nil
					})
				},
				id: "",
			},
			expect: false,
		},
		{
			name: "when guildID is provided, and matches the session guildID, return true",
			input: struct {
				api func(t *testing.T) mockGuild
				id  string
			}{
				api: func(t *testing.T) mockGuild {
					return mockGuild(func(guildID string, options ...discordgo.RequestOption) (st *discordgo.Guild, err error) {
						return &discordgo.Guild{
							ID: testGuildId,
						}, nil
					})
				},
				id: testGuildId,
			},
			expect: true,
		},
		{
			name: "when guildID is provided, and does not match the session guildID, return false",
			input: struct {
				api func(t *testing.T) mockGuild
				id  string
			}{
				api: func(t *testing.T) mockGuild {
					return mockGuild(func(guildID string, options ...discordgo.RequestOption) (st *discordgo.Guild, err error) {
						return &discordgo.Guild{
							ID: testGuildId,
						}, nil
					})
				},
				id: "987654321",
			},
			expect: false,
		},
	}

	for _, test := range cases {
		result := CheckGuildId(test.input.api(t), test.input.id)
		assert.Equal(t, test.expect, result)
	}
}
