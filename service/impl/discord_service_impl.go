package impl

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/ichiro-its/discord-pr-bot/service"
)

type DiscordServiceImpl struct {
	session *discordgo.Session
}

func (d *DiscordServiceImpl) GetMessages(channelId string) ([]*discordgo.Message, error) {
	messages, err := d.session.ChannelMessages(channelId, 100, "", "", "")
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (d *DiscordServiceImpl) SendMessage(channelId string, content string) error {
	_, err := d.session.ChannelMessageSend(channelId, content)
	return err
}

func (d *DiscordServiceImpl) UpdateMessage(channelId string, messageId string, content string) error {
	_, err := d.session.ChannelMessageEdit(channelId, messageId, content)
	return err
}

func (d *DiscordServiceImpl) DeleteMessage(channelId string, messageId string) error {
	return d.session.ChannelMessageDelete(channelId, messageId)
}

func NewDiscordService(token string) (service.DiscordService, error) {
	if token == "" {
		return nil, fmt.Errorf("discord bot token is empty")
	}
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("error creating Discord session: %+v", err)
	}

	return &DiscordServiceImpl{
		session: session,
	}, nil
}
