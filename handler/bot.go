package handler

import (
	"fmt"
	"time"

	"github.com/ichiro-its/discord-pr-bot/config"
	"github.com/ichiro-its/discord-pr-bot/constants"
	"github.com/ichiro-its/discord-pr-bot/entity"
	"github.com/ichiro-its/discord-pr-bot/service"
	"github.com/ichiro-its/discord-pr-bot/service/impl"
)

type Bot struct {
	discordService service.DiscordService
	githubService  service.GithubService
	channelID      string
	messageID      string
	githubOrg      string
}

func NewBot(config *config.Config) (*Bot, error) {
	discordService, err := impl.NewDiscordService(config.DiscordBotToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create discord service: %+v", err)
	}

	githubService, err := impl.NewGithubService(config.GithubToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create github service: %+v", err)
	}

	return &Bot{
		discordService: discordService,
		githubService:  githubService,
		channelID:      config.DiscordChannelID,
		messageID:      config.DiscordMessageID,
		githubOrg:      config.GithubOrg,
	}, nil
}

func (b *Bot) Process() error {
	// Get open pull requests
	pullRequests, err := b.githubService.GetOpenPullRequests(b.githubOrg)
	if err != nil {
		return fmt.Errorf("failed to get open pull requests: %+v", err)
	}
	// TODO: support listing all open pull requests
	pullRequests = pullRequests[max(0, len(pullRequests)-10):]

	var message string
	if len(pullRequests) == 0 {
		message = "Congratulations! No open pull requests.\n"
	} else {
		message = "Open pull requests:\n" + constructGroupedPrMessage(pullRequests)
	}
	message += "\n_Result updated at: " + time.Now().Format(constants.StandardTimeLayout) + "WIB_"

	// Update message in Discord
	err = b.discordService.UpdateMessage(b.channelID, b.messageID, message)
	if err != nil {
		return fmt.Errorf("failed to update message: %+v", err)
	}

	return nil
}

func constructGroupedPrMessage(pullRequests []entity.PullRequest) string {
	repoMap := make(map[string][]entity.PullRequest)
	for _, pr := range pullRequests {
		repoMap[string(pr.Repository.Name)] = append(repoMap[string(pr.Repository.Name)], pr)
	}

	message := ""
	for repo, prs := range repoMap {
		message += fmt.Sprintf("- **%s**\n", repo)
		for _, pr := range prs {
			message += fmt.Sprintf(" - [%s](<%s>) (%s)\n", pr.Title, pr.Url.URL, pr.Author.Login)
		}
	}
	return message
}
