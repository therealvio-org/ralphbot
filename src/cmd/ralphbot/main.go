package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/ralphbot/internal/config"
	"github.com/ralphbot/internal/discord"
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
