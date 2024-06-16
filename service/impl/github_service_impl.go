package impl

import (
	"github.com/ichiro-its/discord-pr-bot/service"
)

type GithubServiceImpl struct {
}

func (*GithubServiceImpl) GetOpenPullRequestUrls() (pullRequestUrls []string, err error) {
	return nil, nil
}

func NewGithubService() (service.GithubService, error) {
	return &GithubServiceImpl{}, nil
}
