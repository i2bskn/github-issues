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

func githubIssues(token string) (issues []github.Issue, err error) {
	ts := &tokenSource{
		&oauth2.Token{AccessToken: token},
	}

	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)
	options := &github.IssueListOptions{}
	issues, _, err = client.Issues.List(true, options)
	return
}
