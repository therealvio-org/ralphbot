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

	// pre-flight checks ✅
	discord.CheckGuildId(ds, env.GuildID)
	discord.CheckOnline(ds)

	// launch! 🚀
	err = discord.StartBotService(ds, env)
	if err != nil {
		log.Fatalf("Error starting new discord session: %v", err)
	}
}
