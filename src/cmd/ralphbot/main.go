package main

import (
	"log"

	"ralphbot/internal/config"
	"ralphbot/internal/discord"
)

var (
	env = config.New()
)

func main() {
	ds, err := discord.NewDiscord(env.BotToken)
	if err != nil {
		log.Fatalf("Error executing NewDiscord(): %v", err)
	}

	// pre-flight checks go here
	discord.CheckGuildId(ds, env.GuildID)

	// launch! 🚀
	discord.StartBotService(ds, env)
}
