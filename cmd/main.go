package main

import (
	"github.com/ichiro-its/discord-pr-bot/config"
	"github.com/ichiro-its/discord-pr-bot/handler"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cfg := config.LoadConfig()
	bot, err := handler.NewBot(cfg)
	if err != nil {
		panic(err)
	}
	err = bot.Process()
	if err != nil {
		panic(err)
	}
}
