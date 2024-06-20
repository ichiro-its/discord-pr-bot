package main

import (
	"os"

	"github.com/ichiro-its/discord-pr-bot/config"
	"github.com/ichiro-its/discord-pr-bot/handler"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	if len(os.Args) < 2 {
		panic("missing repo argument")
	}
	cfg := config.LoadConfig()
	bot, err := handler.NewBot(cfg)
	if err != nil {
		panic(err)
	}
	err = bot.Process(os.Args[1])
	if err != nil {
		panic(err)
	}
}
