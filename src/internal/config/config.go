package config

import "github.com/kelseyhightower/envconfig"

type EnvConfig struct {
	BotToken       string `split_words:"true" required:"true"`
	GuildID        string `split_words:"true"`                //"Test guild ID. If not passed - bot registers commands globally"
	RemoveCommands bool   `default:"true" split_words:"true"` //"Remove all commands after shutting down or not"
}

func New() *EnvConfig {
	var envConfig EnvConfig
	envconfig.MustProcess("", &envConfig)

	return &envConfig
}
