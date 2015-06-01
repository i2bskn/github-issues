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
	// FormatDefault is default value of display format
	FormatDefault = "%n\t%l\t%t\t%u"
)

// Options API Request
type Options struct {
	page    int
	perPage int
	token   string
	format  *Format
}

func newOptions(c *cli.Context) *Options {
	token, err := getGitConfig(PersonalAccessTokenKey)
	if err != nil {
		fail("Must be token settings to .gitconfig")
	}

	page := PageDefault
	perPage := PerPageDefault
	format := FormatDefault

	if c.Int("page") != 0 {
		page = c.Int("page")
	}

	if c.Int("per-page") != 0 {
		perPage = c.Int("per-page")
	}

	if c.String("format") != "" {
		format = c.String("format")
	}

	return &Options{
		page:    page,
		perPage: perPage,
		token:   token,
		format:  newFormat(format),
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
