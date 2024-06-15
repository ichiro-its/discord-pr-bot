package main

import (
	"log"

	"github.com/ichiro-its/discord-pr-bot/service"
	"github.com/ichiro-its/discord-pr-bot/service/impl"
)

func main() {
	discordService := impl.NewDiscordService()
	githubService := impl.NewGithubService()

	pullRequestUrls, err := githubService.GetOpenPullRequestUrls()
	if err != nil {
		log.Fatalf("Failed to get open pull request urls: %+v", err)
	}

	if len(pullRequestUrls) == 0 {
		updateMessage(discordService, "No open pull requests")
		return
	}

	message := "Open pull requests:\n"
	for _, pullRequestUrl := range pullRequestUrls {
		message += pullRequestUrl + "\n"
	}

	updateMessage(discordService, "No open pull requests")
}

func updateMessage(discordService service.DiscordService, message string) {
	err := discordService.UpdateMessage(message)
	if err != nil {
		log.Fatalf("Failed to update message: %+v", err)
	}
}
