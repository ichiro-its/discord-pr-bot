package impl

import (
	"github.com/ichiro-its/discord-pr-bot/service"
)

type DiscordServiceImpl struct {
}

func (*DiscordServiceImpl) UpdateMessage(message string) error {
	return nil
}

func NewDiscordService() service.DiscordService {
	return &DiscordServiceImpl{}
}
