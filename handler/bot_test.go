package handler

import (
	"errors"
	"net/url"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/ichiro-its/discord-pr-bot/config"
	"github.com/ichiro-its/discord-pr-bot/constants"
	"github.com/ichiro-its/discord-pr-bot/entity"
	"github.com/ichiro-its/discord-pr-bot/mocks/service"
	"github.com/shurcooL/githubv4"
	"github.com/stretchr/testify/assert"
)

const (
	testDiscordBotToken = "test-discord-bot-token"
	testDiscordBotId    = "test-discord-bot-id"
	testChannelId       = "test-channel-id"
	testMessageId       = "test-message-id"

	testGithubToken = "test-github-token"
	testGithubOrg   = "test-github-org"

	testRepo     = "test-repo"
	boldTestRepo = "**" + testRepo + "**"
	prTypeClosed = "closed"
)

func TestNewBot(t *testing.T) {
	// Create a new Bot instance
	bot, _ := NewBot(&config.Config{
		DiscordBotToken:  testDiscordBotToken,
		DiscordBotId:     testDiscordBotId,
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
	type fields struct {
		discordService *service.DiscordServiceMock
		githubService  *service.GithubServiceMock
		botId          string
		channelID      string
		messageID      string
		githubOrg      string
	}
	type args struct {
		botParam *entity.BotParam
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		setupMocks  func(fields fields)
		expectedErr error
	}{
		{
			name: "No pull requests with no existing message",
			fields: fields{
				discordService: new(service.DiscordServiceMock),
				githubService:  new(service.GithubServiceMock),
				botId:          testDiscordBotId,
				channelID:      testChannelId,
				messageID:      testMessageId,
				githubOrg:      testGithubOrg,
			},
			args: args{
				botParam: &entity.BotParam{
					Repository: testRepo,
					PrType:     prTypeClosed,
				},
			},
			setupMocks: func(f fields) {
				f.githubService.On("GetOpenPullRequests", testGithubOrg, testRepo).
					Return([]*entity.PullRequest{}, nil)

				f.discordService.On("GetMessages", testChannelId).
					Return([]*discordgo.Message{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "No pull requests with existing message",
			fields: fields{
				discordService: new(service.DiscordServiceMock),
				githubService:  new(service.GithubServiceMock),
				botId:          testDiscordBotId,
				channelID:      testChannelId,
				messageID:      testMessageId,
				githubOrg:      testGithubOrg,
			},
			args: args{
				botParam: &entity.BotParam{
					Repository: testRepo,
					PrType:     prTypeClosed,
				},
			},
			setupMocks: func(f fields) {
				f.githubService.On("GetOpenPullRequests", testGithubOrg, testRepo).
					Return([]*entity.PullRequest{}, nil)

				f.discordService.On("GetMessages", testChannelId).
					Return([]*discordgo.Message{
						{
							ID:      "user-message-id",
							Content: boldTestRepo,
							Author: &discordgo.User{
								ID: "test-user-id",
							},
						},
						{
							ID:      testMessageId,
							Content: boldTestRepo,
							Author: &discordgo.User{
								ID: testDiscordBotId,
							},
						},
					}, nil)

				f.discordService.On("DeleteMessage", testChannelId, testMessageId).
					Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "Multiple pull requests with no existing message",
			fields: fields{
				discordService: new(service.DiscordServiceMock),
				githubService:  new(service.GithubServiceMock),
				botId:          testDiscordBotId,
				channelID:      testChannelId,
				messageID:      testMessageId,
				githubOrg:      testGithubOrg,
			},
			args: args{
				botParam: &entity.BotParam{
					Repository: testRepo,
					PrType:     "opened",
				},
			},
			setupMocks: func(f fields) {
				pullRequests := []*entity.PullRequest{
					{
						Title:      "PR1",
						Url:        mustParseURL("https://github.com/org/repo/pull/1"),
						Author:     entity.Author{Login: "user1"},
						Repository: entity.Repository{Name: testRepo},
					},
					{
						Title:      "PR2",
						Url:        mustParseURL("https://github.com/org/repo/pull/2"),
						Author:     entity.Author{Login: "user2"},
						Repository: entity.Repository{Name: testRepo},
					},
				}

				f.githubService.On("GetOpenPullRequests", testGithubOrg, testRepo).
					Return(pullRequests, nil)

				f.discordService.On("GetMessages", testChannelId).
					Return([]*discordgo.Message{}, nil)

				f.discordService.On("SendMessage", testChannelId, "**test-repo**\n- [PR1](<https://github.com/org/repo/pull/1>) (user1)\n- [PR2](<https://github.com/org/repo/pull/2>) (user2)\n").
					Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "Multiple pull requests with existing message",
			fields: fields{
				discordService: new(service.DiscordServiceMock),
				githubService:  new(service.GithubServiceMock),
				botId:          testDiscordBotId,
				channelID:      testChannelId,
				messageID:      testMessageId,
				githubOrg:      testGithubOrg,
			},
			args: args{
				botParam: &entity.BotParam{
					Repository: testRepo,
					PrType:     constants.GithubPrTypeOpened,
				},
			},
			setupMocks: func(f fields) {
				pullRequests := []*entity.PullRequest{
					{
						Title:      "PR1",
						Url:        mustParseURL("https://github.com/org/repo/pull/1"),
						Author:     entity.Author{Login: "user1"},
						Repository: entity.Repository{Name: testRepo},
					},
					{
						Title:      "PR2",
						Url:        mustParseURL("https://github.com/org/repo/pull/2"),
						Author:     entity.Author{Login: "user2"},
						Repository: entity.Repository{Name: testRepo},
					},
				}

				f.githubService.On("GetOpenPullRequests", testGithubOrg, testRepo).
					Return(pullRequests, nil)

				f.discordService.On("GetMessages", testChannelId).
					Return([]*discordgo.Message{
						{
							ID:      testMessageId,
							Content: boldTestRepo,
							Author: &discordgo.User{
								ID: testDiscordBotId,
							},
						},
					}, nil)

				f.discordService.On("DeleteMessage", testChannelId, testMessageId).
					Return(nil)

				f.discordService.On("SendMessage", testChannelId, "**test-repo**\n- [PR1](<https://github.com/org/repo/pull/1>) (user1)\n- [PR2](<https://github.com/org/repo/pull/2>) (user2)\n").
					Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "GitHub service error",
			fields: fields{
				discordService: new(service.DiscordServiceMock),
				githubService:  new(service.GithubServiceMock),
				botId:          testDiscordBotId,
				channelID:      testChannelId,
				messageID:      testMessageId,
				githubOrg:      testGithubOrg,
			},
			args: args{
				botParam: &entity.BotParam{
					Repository: testRepo,
					PrType:     constants.GithubPrTypeOpened,
				},
			},
			setupMocks: func(f fields) {
				f.githubService.On("GetOpenPullRequests", testGithubOrg, testRepo).
					Return(nil, errors.New("github error"))

				f.discordService.On("GetMessages", testChannelId).
					Return([]*discordgo.Message{}, nil)
			},
			expectedErr: errors.New("failed to get pull requests or messages: github error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bot := &Bot{
				discordService: tt.fields.discordService,
				githubService:  tt.fields.githubService,
				botId:          tt.fields.botId,
				channelID:      tt.fields.channelID,
				messageID:      tt.fields.messageID,
				githubOrg:      tt.fields.githubOrg,
			}
			tt.setupMocks(tt.fields)

			err := bot.Process(tt.args.botParam)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

			tt.fields.discordService.AssertExpectations(t)
			tt.fields.githubService.AssertExpectations(t)
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
