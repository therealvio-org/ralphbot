package discord

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

// This is more of a regression test. NewDiscord() shouldn't be doing more than setting up a new session
// additionally, if Discord.New() requires extra parameters, or returns more than just the session, this test will catch it outside of the assertions
func TestNewDiscord(t *testing.T) {
	dgs := &discordgo.Session{}

	ds, err := NewDiscord("mockSecret123")

	assert.NoError(t, err)
	assert.IsType(t, dgs, ds.Session)
}
