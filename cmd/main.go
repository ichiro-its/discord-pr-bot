package main

import (
	"encoding/json"
	"os"

	"github.com/ichiro-its/discord-pr-bot/config"
	"github.com/ichiro-its/discord-pr-bot/entity"
	"github.com/ichiro-its/discord-pr-bot/handler"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	if len(os.Args) < 2 {
		panic("missing parameter argument")
	}
	cfg := config.LoadConfig()
	bot, err := handler.NewBot(cfg)
	if err != nil {
		panic(err)
	}
	botParam := &entity.BotParam{}
	err = json.Unmarshal([]byte(os.Args[1]), botParam)
	if err != nil {
		panic(err)
	}
	err = bot.Process(botParam)
	if err != nil {
		panic(err)
	}
}
