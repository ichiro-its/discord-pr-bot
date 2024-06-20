package service

import "github.com/bwmarrin/discordgo"

type DiscordService interface {
	GetMessages(channelId string) ([]*discordgo.Message, error)
	SendMessage(channelId string, content string) error
	UpdateMessage(channelId string, messageId string, content string) error
	DeleteMessage(channelId string, messageId string) error
}
