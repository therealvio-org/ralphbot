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

var s *discordgo.Session

func init() {
	var err error
	s, err = discordgo.New("Bot " + env.BotToken)
	if err != nil {
		log.Fatalf("Error creating new Discord session: %v", err)
	}
}

func main() {
	discord.CheckGuildId(s, env.GuildID)
	discord.StartBotService(s, env)
}
