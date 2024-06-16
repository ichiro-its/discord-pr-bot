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
	expectedDiscordChannelID := "test-discord-channel-id"
	expectedDiscordMessageID := "test-discord-message-id"

	os.Setenv(constants.DiscordBotTokenEnv, expectedDiscordBotToken)
	os.Setenv(constants.DiscordChannelIdEnv, expectedDiscordChannelID)
	os.Setenv(constants.DiscordMessageIdEnv, expectedDiscordMessageID)

	// Ensure environment variables are unset after the test
	defer func() {
		os.Unsetenv(constants.DiscordBotTokenEnv)
		os.Unsetenv(constants.DiscordChannelIdEnv)
		os.Unsetenv(constants.DiscordMessageIdEnv)
	}()

	// Call LoadConfig and verify the returned values
	cfg := LoadConfig()

	assert.NotNil(t, cfg, "Config should not be nil")
	assert.Equal(t, expectedDiscordBotToken, cfg.DiscordBotToken, "DiscordBotToken should match")
	assert.Equal(t, expectedDiscordChannelID, cfg.DiscordChannelID, "DiscordChannelID should match")
	assert.Equal(t, expectedDiscordMessageID, cfg.DiscordMessageID, "DiscordMessageID should match")
}
