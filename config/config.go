package config

import (
	"os"

	"github.com/ichiro-its/discord-pr-bot/constants"
)

type Config struct {
	DiscordBotToken  string
	DiscordBotId     string
	DiscordChannelID string
	DiscordMessageID string

	GithubToken string
	GithubOrg   string
}

func LoadConfig() *Config {
	return &Config{
		DiscordBotToken:  os.Getenv(constants.DiscordBotTokenEnv),
		DiscordBotId:     os.Getenv(constants.DiscordBotIdEnv),
		DiscordChannelID: os.Getenv(constants.DiscordChannelIdEnv),
		DiscordMessageID: os.Getenv(constants.DiscordMessageIdEnv),
		GithubToken:      os.Getenv(constants.GithubTokenEnv),
		GithubOrg:        os.Getenv(constants.GithubOrgEnv),
	}
}
