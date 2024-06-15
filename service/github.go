package service

type GithubService interface {
	GetOpenPullRequestUrls() (pullRequestUrls []string, err error)
}
