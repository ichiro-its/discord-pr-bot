package service

type DiscordService interface {
	UpdateMessage(channelId string, messageId string, content string) error
}
