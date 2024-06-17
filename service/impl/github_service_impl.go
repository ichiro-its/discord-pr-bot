package impl

import (
	"context"
	"fmt"
	"sort"

	"github.com/ichiro-its/discord-pr-bot/entity"
	"github.com/ichiro-its/discord-pr-bot/service"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type GithubServiceImpl struct {
	client *githubv4.Client
}

type searchQuery struct {
	Search struct {
		Edges []struct {
			Node struct {
				PullRequest entity.PullRequest `graphql:"... on PullRequest"`
			}
		}
		PageInfo struct {
			EndCursor   githubv4.String
			HasNextPage githubv4.Boolean
		}
	} `graphql:"search(query: $query, type: ISSUE, last: 100)"`
}

func (g *GithubServiceImpl) GetOpenPullRequests(org string) ([]entity.PullRequest, error) {
	var query searchQuery
	err := g.client.Query(context.Background(), &query, map[string]interface{}{
		"query": githubv4.String(fmt.Sprintf("org:%s is:pr is:open", org)),
	})
	if err != nil {
		return nil, fmt.Errorf("failed github query: %+v", err)
	}

	var pullRequests []entity.PullRequest
	for _, edge := range query.Search.Edges {
		pr := edge.Node.PullRequest
		if pr.Mergeable == githubv4.MergeableStateMergeable {
			pullRequests = append(pullRequests, edge.Node.PullRequest)
		}
	}
	sort.Slice(pullRequests, func(i, j int) bool {
		return pullRequests[i].CreatedAt.Time.Before(pullRequests[j].CreatedAt.Time)
	})
	return pullRequests, nil
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
