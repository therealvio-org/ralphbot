package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type DiscordAPI interface {
	Guild(guildID string, options ...discordgo.RequestOption) (st *discordgo.Guild, err error)
}

// If GUILD_ID was passed, it checks to see if the GUILD_ID environment variable matches the discord server's (guild) id that
// `ralphbot` is connected to.
// If it is, it indicates in the logs that its commands will be registered as `guild` commands, rather than `global` commands.
// i.e. immediate registration for the server, vs a live command
//
// The intention for this is to help facilitate local testing on a private server. In production, commands should be `global`.
// Return `true` if `id` matches guildId, otherwise return `false`
func CheckGuildId(dAPI DiscordAPI, id string) bool {
	doFail := false
	if id != "" {
		log.Printf("Checking for GuildID variable validity against current server (if it matches, we're okay)...")

		g, err := dAPI.Guild(id)
		if err != nil {
			log.Fatalf("Cannot retrieve Guild Id from server: %v", err)
		}

		if id == g.ID {
			log.Printf("GuildID is valid...")
			log.Printf("GuildID is defined - ralphbot is running in development mode, Guild commands are available...")
			return true
		} else {
			log.Printf("GuildID is invalid...")
			doFail = true
		}
	} else {
		doFail = true
	}

	if doFail {
		log.Printf("GuildID is undefined - ralphbot is running in production mode, only Global commands are available...")
		return false
	}

	return false
}
