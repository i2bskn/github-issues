package main

import (
	"bytes"
	"github.com/codegangsta/cli"
	"os/exec"
	"strings"
)

const (
	// PersonalAccessTokenKey in .gitconfig
	personalAccessTokenKey = "github.token"
)

// Options API Request
type Options struct {
	page    int
	perPage int
	filter  string
	state   string
	sort    string
	token   string
	format  *Format
}

func newOptions(c *cli.Context) *Options {
	token, err := getGitConfig(personalAccessTokenKey)
	if err != nil {
		fail(err.Error())
	}

	state, err := validState(c.String("state"))
	if err != nil {
		fail(err.Error())
	}

	sort, err := validSort(c.String("sort"))
	if err != nil {
		fail(err.Error())
	}

	return &Options{
		page:    c.Int("page"),
		perPage: c.Int("per-page"),
		filter:  parseFilter(c),
		state:   state,
		sort:    sort,
		token:   token,
		format:  newFormat(c.String("format")),
	}
}

func parseFilter(c *cli.Context) string {
	switch {
	case c.Bool("assigned"):
		return "assigned"
	case c.Bool("created"):
		return "created"
	case c.Bool("mentioned"):
		return "mentioned"
	default:
		return "all"
	}
}

func validState(state string) (valid_state string, err error) {
	invalid := true
	for _, s := range [...]string{"open", "closed", "all"} {
		if state == s {
			valid_state = state
			invalid = false
		}
	}
	if invalid {
		err = newError("Invalid state: " + state)
	}
	return
}

func validSort(sort string) (valid_sort string, err error) {
	invalid := true
	for _, s := range [...]string{"created", "updated", "comments"} {
		if sort == s {
			valid_sort = sort
			invalid = false
		}
	}
	if invalid {
		err = newError("Invalid sort: " + sort)
	}
	return
}

func getGitConfig(key string) (out string, err error) {
	cmd := exec.Command("git", "config", key)
	var result bytes.Buffer
	cmd.Stdout = &result

	err = cmd.Run()
	out = strings.TrimSpace(result.String())
	return
}
