package handler

import (
	"net/url"
	"testing"
	"time"

	"github.com/ichiro-its/discord-pr-bot/config"
	"github.com/ichiro-its/discord-pr-bot/constants"
	"github.com/ichiro-its/discord-pr-bot/entity"
	"github.com/ichiro-its/discord-pr-bot/mocks/service"
	"github.com/shurcooL/githubv4"
	"github.com/stretchr/testify/assert"
)

type TableTest struct {
	name             string
	mockGithubResult []entity.PullRequest // Simulated result from mock GithubService
	expectedMessage  string               // Expected message sent to DiscordService
	mockGithubError  error                // Simulated error from mock GithubService
	mockDiscordError error                // Simulated error from mock DiscordService

}

const (
	testDiscordBotToken = "test-discord-bot-token"
	testChannelId       = "test-channel-id"
	testMessageId       = "test-message-id"

	testGithubToken = "test-github-token"
	testGithubOrg   = "test-org"
)

func TestNewBot(t *testing.T) {
	// Create a new Bot instance
	bot, _ := NewBot(&config.Config{
		DiscordBotToken:  testDiscordBotToken,
		DiscordChannelID: testChannelId,
		DiscordMessageID: testMessageId,
		GithubToken:      testGithubToken,
		GithubOrg:        testGithubOrg,
	})

	// Verify that the Bot instance is created correctly
	assert.NotNil(t, bot)
	assert.NotNil(t, bot.discordService)
	assert.NotNil(t, bot.githubService)
	assert.Equal(t, testChannelId, bot.channelID)
	assert.Equal(t, testMessageId, bot.messageID)
	assert.Equal(t, testGithubOrg, bot.githubOrg)
}

func TestBotProcess(t *testing.T) {
	tests := []*TableTest{
		{
			name:             "No open pull requests",
			mockGithubResult: []entity.PullRequest{},
			expectedMessage:  "Congratulations! No open pull requests.\n\n_Result updated at: " + time.Now().Format(constants.StandardTimeLayout) + "WIB_",
			mockGithubError:  nil,
			mockDiscordError: nil,
		},
		{
			name: "Open pull requests",
			mockGithubResult: []entity.PullRequest{
				{
					Title: "Fix issue #123",
					Url:   mustParseURL("https://github.com/org/repo/pull/123"),
					Author: entity.Author{
						Login: "mockuser",
					},
					Repository: entity.Repository{
						Name: "repo",
					},
					CreatedAt: mustParseTime("2021-01-01 00:00:01"),
				},
				{
					Title: "Add feature X",
					Url:   mustParseURL("https://github.com/org/repo/pull/124"),
					Author: entity.Author{
						Login: "mockuser2",
					},
					Repository: entity.Repository{
						Name: "repo2",
					},
					CreatedAt: mustParseTime("2021-01-01 00:00:00"),
				},
			},
			expectedMessage:  "Open pull requests:\n- **repo**\n - [Fix issue #123](<https://github.com/org/repo/pull/123>) (mockuser)\n- **repo2**\n - [Add feature X](<https://github.com/org/repo/pull/124>) (mockuser2)\n\n_Result updated at: " + time.Now().Format(constants.StandardTimeLayout) + "WIB_",
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
			mockGithubResult: []entity.PullRequest{},
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
				mockGithubService.On("GetOpenPullRequests", testGithubOrg).Return(tt.mockGithubResult, tt.mockGithubError).Once()
			} else {
				mockGithubService.On("GetOpenPullRequests", testGithubOrg).Return(tt.mockGithubResult, tt.mockGithubError)
			}

			// Mock DiscordService
			mockDiscordService := &service.DiscordServiceMock{}
			if tt.mockDiscordError != nil {
				mockDiscordService.On("UpdateMessage", testChannelId, testMessageId, tt.expectedMessage).Return(tt.mockDiscordError).Once()
			} else {
				mockDiscordService.On("UpdateMessage", testChannelId, testMessageId, tt.expectedMessage).Return(nil).Once()
			}

			// Create Bot instance
			bot := &Bot{
				discordService: mockDiscordService,
				githubService:  mockGithubService,
				channelID:      testChannelId,
				messageID:      testMessageId,
				githubOrg:      testGithubOrg,
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

func mustParseURL(s string) githubv4.URI {
	u, err := url.Parse(s)
	if err != nil {
		panic("mustParseURL: parsing URL failed: " + err.Error())
	}
	return githubv4.URI{URL: u}
}

func mustParseTime(s string) githubv4.DateTime {
	t, err := time.Parse(constants.StandardTimeLayout, s)
	if err != nil {
		panic("mustParseTime: parsing time failed: " + err.Error())
	}
	return githubv4.DateTime{Time: t}
}
