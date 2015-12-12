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

func githubIssues(options *Options) ([]github.Issue, *github.Response, error) {
	ts := &tokenSource{
		&oauth2.Token{AccessToken: options.Token()},
	}

	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)
	issueListOptions := &github.IssueListOptions{
		Filter: options.Filter(),
		State:  options.State,
		Sort:   options.Sort,
		ListOptions: github.ListOptions{
			Page:    options.Page,
			PerPage: options.PerPage,
		},
	}
	return client.Issues.List(true, issueListOptions)
}
