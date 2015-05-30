package main

import (
	"bytes"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
	"strings"
)

const (
	// PersonalAccessTokenKey in .gitconfig
	PersonalAccessTokenKey = "github.token"
)

func main() {
	app := cli.NewApp()
	app.Name = "github-issues"
	app.Version = "0.0.1"
	app.Usage = "List of GitHub issues"
	app.Author = "i2bskn"
	app.Email = "i2bskn@gmail.com"
	app.Flags = []cli.Flag{}
	app.Action = func(c *cli.Context) {
		token, err := getGitConfig(PersonalAccessTokenKey)
		if err != nil {
			fail("Must be token settings to .gitconfig")
		}

		issues, err := githubIssues(token)
		if err != nil {
			fail("Failed get issues")
		}

		for _, issue := range issues {
			fmt.Println(issue)
		}
	}
	app.Run(os.Args)
}

func getGitConfig(key string) (out string, err error) {
	cmd := exec.Command("git", "config", key)
	var result bytes.Buffer
	cmd.Stdout = &result

	err = cmd.Run()
	out = strings.TrimSpace(result.String())
	return
}

func fail(message string) {
	fmt.Println(message)
	os.Exit(1)
}
