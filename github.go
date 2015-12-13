package main

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type tokenSource struct {
	token *oauth2.Token
}

func (t *tokenSource) Token() (*oauth2.Token, error) {
	return t.token, nil
}

// Issues and pull requests of GitHub.
type Issues struct {
	client *github.Client
}

// NewIssues to generate issues object.
func NewIssues(token string) *Issues {
	source := &tokenSource{
		&oauth2.Token{AccessToken: token},
	}
	client := oauth2.NewClient(oauth2.NoContext, source)

	return &Issues{
		client: github.NewClient(client),
	}
}

// ListByOptions to get issues and pull requests from GitHub.
func (issues Issues) ListByOptions(options *Options) ([]github.Issue, *github.Response, error) {
	if len(options.repository) > 0 || options.CurrentRepo {
		return issues.listByRepo(options)
	}

	listOptions := &github.IssueListOptions{
		Filter: options.Filter(),
		State:  options.State,
		Sort:   options.Sort,
		ListOptions: github.ListOptions{
			Page:    options.Page,
			PerPage: options.PerPage,
		},
	}
	return issues.client.Issues.List(!options.Self, listOptions)
}

func (issues Issues) listByRepo(options *Options) ([]github.Issue, *github.Response, error) {
	owner, repo, err := options.getOwnerAndRepo()
	if err != nil {
		return nil, nil, err
	}

	listOptions := &github.IssueListByRepoOptions{
		State: options.State,
		Sort:  options.Sort,
		ListOptions: github.ListOptions{
			Page:    options.Page,
			PerPage: options.PerPage,
		},
	}
	return issues.client.Issues.ListByRepo(owner, repo, listOptions)
}
