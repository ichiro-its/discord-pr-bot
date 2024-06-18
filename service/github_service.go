package service

import "github.com/ichiro-its/discord-pr-bot/entity"

type GithubService interface {
	GetOpenPullRequests(org string) ([]entity.PullRequest, error)
}
