package config

import (
	"os"
	"testing"

	"github.com/ichiro-its/discord-pr-bot/constants"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Set up environment variables for testing
	expectedDiscordBotToken := "test-discord-bot-token"
	expectedDiscordBotId := "test-discord-bot-id"
	expectedDiscordChannelID := "test-discord-channel-id"
	expectedDiscordMessageID := "test-discord-message-id"
	expectedGithubToken := "test-github-token"
	expectedGithubOrg := "test-github-org"

	os.Setenv(constants.DiscordBotTokenEnv, expectedDiscordBotToken)
	os.Setenv(constants.DiscordBotIdEnv, expectedDiscordBotId)
	os.Setenv(constants.DiscordChannelIdEnv, expectedDiscordChannelID)
	os.Setenv(constants.DiscordMessageIdEnv, expectedDiscordMessageID)
	os.Setenv(constants.GithubTokenEnv, expectedGithubToken)
	os.Setenv(constants.GithubOrgEnv, expectedGithubOrg)

	// Ensure environment variables are unset after the test
	defer func() {
		os.Unsetenv(constants.DiscordBotTokenEnv)
		os.Unsetenv(constants.DiscordBotIdEnv)
		os.Unsetenv(constants.DiscordChannelIdEnv)
		os.Unsetenv(constants.DiscordMessageIdEnv)
		os.Unsetenv(constants.GithubTokenEnv)
		os.Unsetenv(constants.GithubOrgEnv)
	}()

	// Call LoadConfig and verify the returned values
	cfg := LoadConfig()

	assert.NotNil(t, cfg, "Config should not be nil")
	assert.Equal(t, expectedDiscordBotToken, cfg.DiscordBotToken, "DiscordBotToken should match")
	assert.Equal(t, expectedDiscordBotId, cfg.DiscordBotId, "DiscordBotId should match")
	assert.Equal(t, expectedDiscordChannelID, cfg.DiscordChannelID, "DiscordChannelID should match")
	assert.Equal(t, expectedDiscordMessageID, cfg.DiscordMessageID, "DiscordMessageID should match")
	assert.Equal(t, expectedGithubToken, cfg.GithubToken, "GithubToken should match")
	assert.Equal(t, expectedGithubOrg, cfg.GithubOrg, "GithubOrg should match")
}
