package config

import (
	"os"

	"github.com/ichiro-its/discord-pr-bot/constants"
)

type Config struct {
	DiscordBotToken  string
	DiscordChannelID string
	DiscordMessageID string
}

func LoadConfig() *Config {
	return &Config{
		DiscordBotToken:  os.Getenv(constants.DiscordBotTokenEnv),
		DiscordChannelID: os.Getenv(constants.DiscordChannelIdEnv),
		DiscordMessageID: os.Getenv(constants.DiscordMessageIdEnv),
	}
}
