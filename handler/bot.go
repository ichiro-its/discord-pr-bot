package handler

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/ichiro-its/discord-pr-bot/config"
	"github.com/ichiro-its/discord-pr-bot/constants"
	"github.com/ichiro-its/discord-pr-bot/entity"
	"github.com/ichiro-its/discord-pr-bot/service"
	"github.com/ichiro-its/discord-pr-bot/service/impl"
)

type Bot struct {
	discordService service.DiscordService
	githubService  service.GithubService
	botId          string
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
		botId:          config.DiscordBotId,
		channelID:      config.DiscordChannelID,
		messageID:      config.DiscordMessageID,
		githubOrg:      config.GithubOrg,
	}, nil
}

func (b *Bot) Process(repo string) error {
	// Get open pull requests
	pullRequests, err := b.githubService.GetOpenPullRequests(b.githubOrg, repo)
	if err != nil {
		return fmt.Errorf("failed to get open pull requests: %+v", err)
	}
	messages, err := b.discordService.GetMessages(b.channelID)
	if err != nil {
		return fmt.Errorf("failed to get messages: %+v", err)
	}

	messageId := b.getRepoMessageMessageId(messages, repo)
	if len(pullRequests) == 0 {
		return b.deleteMessage(messageId)
	}

	content := b.constructPrMessageContent(pullRequests)
	if messageId != "" {
		return b.updateMessage(messageId, content)
	}

	return b.sendMessage(content)
}

func (b *Bot) constructPrMessageContent(pullRequests []*entity.PullRequest) string {
	message := fmt.Sprintf("**%s**\n", pullRequests[0].Repository.Name)
	for _, pr := range pullRequests {
		prMessage := fmt.Sprintf("- [%s](<%s>) (%s)\n", pr.Title, pr.Url.URL, pr.Author.Login)
		if len(message)+len(prMessage) > constants.DiscordMessageLengthLimit {
			break
		}
		message += prMessage
	}
	return message
}

func (b *Bot) getRepoMessageMessageId(messages []*discordgo.Message, repo string) string {
	for _, message := range messages {
		if message.Author.ID != b.botId {
			continue
		}
		s := strings.Split(message.Content, "**")
		if s[1] == repo {
			return message.ID
		}
	}
	return ""
}

func (b *Bot) deleteMessage(messageId string) error {
	if messageId == "" {
		return nil
	}
	err := b.discordService.DeleteMessage(b.channelID, messageId)
	if err != nil {
		return fmt.Errorf("failed to delete message, id: %s: %+v", messageId, err)
	}
	return nil
}

func (b *Bot) sendMessage(content string) error {
	err := b.discordService.SendMessage(b.channelID, content)
	if err != nil {
		return fmt.Errorf("failed to send message: %+v", err)
	}
	return nil
}

func (b *Bot) updateMessage(messageId string, content string) error {
	err := b.discordService.UpdateMessage(b.channelID, messageId, content)
	if err != nil {
		return fmt.Errorf("failed to update message, id: %s: %+v", messageId, err)
	}
	return nil
}
