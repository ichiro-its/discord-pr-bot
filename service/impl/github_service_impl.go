package impl

import (
	"context"
	"fmt"

	"github.com/ichiro-its/discord-pr-bot/entity"
	"github.com/ichiro-its/discord-pr-bot/service"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type GithubServiceImpl struct {
	client *githubv4.Client
}

type searchQuery struct {
	Repository struct {
		PullRequests struct {
			Nodes []*entity.PullRequest
		} `graphql:"pullRequests(states: OPEN, last: 100)"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}

func (g *GithubServiceImpl) GetOpenPullRequests(org string, repo string) ([]*entity.PullRequest, error) {
	var query searchQuery
	err := g.client.Query(context.Background(), &query, map[string]interface{}{
		"owner": githubv4.String(org),
		"name":  githubv4.String(repo),
	})
	if err != nil {
		return nil, fmt.Errorf("failed github query: %+v", err)
	}
	return query.Repository.PullRequests.Nodes, nil
}

func NewGithubService(token string) (service.GithubService, error) {
	if token == "" {
		return nil, fmt.Errorf("github token is empty")
	}
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	return &GithubServiceImpl{
		client: githubv4.NewClient(httpClient),
	}, nil
}
