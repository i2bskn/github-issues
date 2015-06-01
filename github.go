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

func githubIssues(options *Options) (issues []github.Issue, res *github.Response, err error) {
	ts := &tokenSource{
		&oauth2.Token{AccessToken: options.token},
	}

	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)
	issueListOptions := &github.IssueListOptions{
		Filter: "all",
		State:  "open",
		Sort:   "comments",
		ListOptions: github.ListOptions{
			Page:    options.page,
			PerPage: options.perPage,
		},
	}
	issues, res, err = client.Issues.List(true, issueListOptions)
	return
}
