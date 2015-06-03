package main

import (
	"bytes"
	"github.com/codegangsta/cli"
	"os/exec"
	"strings"
)

const (
	// PersonalAccessTokenKey in .gitconfig
	PersonalAccessTokenKey = "github.token"

	// PageDefault is default value of the `page` parameter
	PageDefault = 1
	// PerPageDefault is default value of the `per_page` parameter
	// Maximum value on the specifications of GitHub API
	PerPageDefault = 100
	// FilterDefault is default value of the `filter` parameter
	FilterDefault = "all"
	// StateDefault is default value of the `state` parameter
	StateDefault = "open"
	// SortDefault is default value of the `sort` parameter
	SortDefault = "updated"
	// FormatDefault is default value of display format
	FormatDefault = "%n\t%l\t%t\t%u"
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
	token, err := getGitConfig(PersonalAccessTokenKey)
	if err != nil {
		fail("Need to set a personal access token to " + PersonalAccessTokenKey + " in gitconfig.")
	}

	page := PageDefault
	perPage := PerPageDefault
	filter := parseFilter(c)
	state := parseState(c)
	sort := SortDefault
	format := FormatDefault

	if c.Int("page") > 0 {
		page = c.Int("page")
	}

	if c.Int("per-page") > 0 {
		perPage = c.Int("per-page")
	}

	if c.String("format") != "" {
		format = c.String("format")
	}

	return &Options{
		page:    page,
		perPage: perPage,
		filter:  filter,
		state:   state,
		sort:    sort,
		token:   token,
		format:  newFormat(format),
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
		return FilterDefault
	}
}

func parseState(c *cli.Context) string {
	switch {
	case c.Bool("closed"):
		return "closed"
	case c.Bool("all"):
		return "all"
	default:
		return StateDefault
	}
}

func getGitConfig(key string) (out string, err error) {
	cmd := exec.Command("git", "config", key)
	var result bytes.Buffer
	cmd.Stdout = &result

	err = cmd.Run()
	out = strings.TrimSpace(result.String())
	return
}
