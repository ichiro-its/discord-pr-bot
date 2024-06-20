package entity

import "github.com/shurcooL/githubv4"

type PullRequest struct {
	Url        githubv4.URI
	Title      githubv4.String
	CreatedAt  githubv4.DateTime
	Repository Repository
	Author     Author
}
