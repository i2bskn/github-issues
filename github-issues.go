package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "github-issues"
	app.Version = "0.0.1"
	app.Usage = "List of GitHub issues"
	app.Author = "i2bskn"
	app.Email = "i2bskn@gmail.com"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "page, p",
			Value: 1,
			Usage: "Specify further pages.",
		},
		cli.IntFlag{
			Name: "per-page, n",
			// Maximum value on the specifications of GitHub API
			Value: 100,
			Usage: "Specify a custom page size.",
		},
		cli.BoolFlag{
			Name:  "assigned, a",
			Usage: "Issues assigned to you.",
		},
		cli.BoolFlag{
			Name:  "created, c",
			Usage: "Issues created by you.",
		},
		cli.BoolFlag{
			Name:  "mentioned, m",
			Usage: "Issues mentioning you.",
		},
		cli.StringFlag{
			Name:  "state, s",
			Value: "open",
			Usage: "Specify the state of the issues to display. Can be either open, closed, all.",
		},
		cli.StringFlag{
			Name:  "sort",
			Value: "updated",
			Usage: "What to sort issues by. Can be either created, updated, comments.",
		},
		cli.StringFlag{
			Name:  "format, f",
			Value: "%n\\t%l\\t%t\\t%u",
			Usage: "Specify the format of the issues to display.",
		},
		cli.StringFlag{
			Name:  "token",
			Usage: "Specify the personal access token of GitHub.",
		},
	}
	app.Action = func(c *cli.Context) {
		if c.Bool("version") {
			printVersion(c)
		}

		if c.Bool("help") {
			printHelp(c)
		}

		options := newOptions(c)
		issues, _, err := githubIssues(options)
		if err != nil {
			fail(err.Error())
		}

		for _, issue := range issues {
			fmt.Println(options.format.Apply(issue))
		}
	}
	app.Run(os.Args)
}

func printVersion(c *cli.Context) {
	fmt.Println(c.App.Name + " " + c.App.Version)
	os.Exit(0)
}

func printHelp(c *cli.Context) {
	cli.ShowAppHelp(c)
	os.Exit(0)
}
