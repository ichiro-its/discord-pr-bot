package handler

import (
	"fmt"
	"time"

	"github.com/ichiro-its/discord-pr-bot/config"
	"github.com/ichiro-its/discord-pr-bot/constants"
	"github.com/ichiro-its/discord-pr-bot/service"
	"github.com/ichiro-its/discord-pr-bot/service/impl"
)

type Bot struct {
	discordService service.DiscordService
	githubService  service.GithubService
	channelID      string
	messageID      string
}

func NewBot(config *config.Config) (*Bot, error) {
	discordService, err := impl.NewDiscordService(config.DiscordBotToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create discord service: %+v", err)
	}

	githubService, err := impl.NewGithubService()
	if err != nil {
		return nil, fmt.Errorf("failed to create github service: %+v", err)
	}

	return &Bot{
		discordService: discordService,
		githubService:  githubService,
		channelID:      config.DiscordChannelID,
		messageID:      config.DiscordMessageID,
	}, nil
}

func (b *Bot) Process() error {
	// Get open pull request urls
	pullRequestUrls, err := b.githubService.GetOpenPullRequestUrls()
	if err != nil {
		return fmt.Errorf("failed to get open pull request urls: %+v", err)
	}

	var message string
	if len(pullRequestUrls) == 0 {
		message = "Congratulations! No open pull requests.\n"
	} else {
		message = "Open pull requests:\n"
		for _, pullRequestUrl := range pullRequestUrls {
			message += pullRequestUrl + "\n"
		}
	}
	message += "\n_Result updated at: " + time.Now().Format(constants.StandardTimeLayout) + "WIB_"

	// Update message in Discord
	err = b.discordService.UpdateMessage(b.channelID, b.messageID, message)
	if err != nil {
		return fmt.Errorf("failed to update message: %+v", err)
	}

	return nil
}
