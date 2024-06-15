package service

type DiscordService interface {
	UpdateMessage(message string) error
}
