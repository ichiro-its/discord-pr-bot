package handler

import (
	"testing"
	"time"

	"github.com/ichiro-its/discord-pr-bot/config"
	"github.com/ichiro-its/discord-pr-bot/constants"
	"github.com/ichiro-its/discord-pr-bot/mocks/service"
	"github.com/stretchr/testify/assert"
)

type TableTest struct {
	name             string
	mockGithubResult []string // Simulated result from mock GithubService
	expectedMessage  string   // Expected message sent to DiscordService
	mockGithubError  error    // Simulated error from mock GithubService
	mockDiscordError error    // Simulated error from mock DiscordService

}

const (
	testDiscordBotToken = "test-discord-bot-token"
	testChannelID       = "test-channel-id"
	testMessageID       = "test-message-id"
)

func TestNewBot(t *testing.T) {
	// Create a new Bot instance
	bot, _ := NewBot(&config.Config{
		DiscordBotToken:  testDiscordBotToken,
		DiscordChannelID: testChannelID,
		DiscordMessageID: testMessageID,
	})

	// Verify that the Bot instance is created correctly
	assert.NotNil(t, bot)
	assert.NotNil(t, bot.discordService)
	assert.NotNil(t, bot.githubService)
	assert.Equal(t, testChannelID, bot.channelID)
	assert.Equal(t, testMessageID, bot.messageID)
}

func TestBotProcess(t *testing.T) {
	tests := []*TableTest{
		{
			name:             "No open pull requests",
			mockGithubResult: []string{},
			expectedMessage:  "Congratulations! No open pull requests.\n\n_Result updated at: " + time.Now().Format(constants.StandardTimeLayout) + "WIB_",
			mockGithubError:  nil,
			mockDiscordError: nil,
		},
		{
			name:             "Open pull requests",
			mockGithubResult: []string{"https://github.com/example/pr1", "https://github.com/example/pr2"},
			expectedMessage:  "Open pull requests:\nhttps://github.com/example/pr1\nhttps://github.com/example/pr2\n\n_Result updated at: " + time.Now().Format(constants.StandardTimeLayout) + "WIB_",
			mockGithubError:  nil,
			mockDiscordError: nil,
		},
		{
			name:             "Error retrieving pull requests",
			mockGithubResult: nil,
			expectedMessage:  "",
			mockGithubError:  assert.AnError,
			mockDiscordError: nil,
		},
		{
			name:             "Error updating message",
			mockGithubResult: []string{},
			expectedMessage:  "Congratulations! No open pull requests.\n\n_Result updated at: " + time.Now().Format(constants.StandardTimeLayout) + "WIB_",
			mockGithubError:  nil,
			mockDiscordError: assert.AnError,
		},
	}

	// Iterate over test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock GithubService
			mockGithubService := &service.GithubServiceMock{}
			if tt.mockGithubError != nil {
				mockGithubService.On("GetOpenPullRequestUrls").Return(tt.mockGithubResult, tt.mockGithubError).Once()
			} else {
				mockGithubService.On("GetOpenPullRequestUrls").Return(tt.mockGithubResult, tt.mockGithubError)
			}

			// Mock DiscordService
			mockDiscordService := &service.DiscordServiceMock{}
			if tt.mockDiscordError != nil {
				mockDiscordService.On("UpdateMessage", testChannelID, testMessageID, tt.expectedMessage).Return(tt.mockDiscordError).Once()
			} else {
				mockDiscordService.On("UpdateMessage", testChannelID, testMessageID, tt.expectedMessage).Return(nil).Once()
			}

			// Create Bot instance
			bot := &Bot{
				discordService: mockDiscordService,
				githubService:  mockGithubService,
				channelID:      testChannelID,
				messageID:      testMessageID,
			}

			// Call Process method
			bot.Process()

			mockGithubService.AssertExpectations(t)
			if tt.mockGithubError == nil {
				mockDiscordService.AssertExpectations(t)
			}
		})
	}
}
