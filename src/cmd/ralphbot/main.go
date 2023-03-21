package main

import (
	"log"

	"ralphbot/internal/config"
	"ralphbot/internal/discord"

	"github.com/bwmarrin/discordgo"
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
	// check if the session is running
	ds.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	discord.CheckGuildId(ds, env.GuildID)

	// launch! 🚀
	err = discord.StartBotService(ds, env)
	if err != nil {
		log.Fatalf("Error starting new discord session: %v", err)
	}
}
