package main

import (
	"flag"
	"fmt"
	"strings"
)

const (
	AppName = "github-issues"
	Version = "0.1.0"
)

const (
	CodeSuccess = 0

	CodeFlagParseFail = 101 + iota
	CodeAPIRequestFail
)

const helpText = `Usage: github-issues [OPTIONS]

List of GitHub issues.

Options:
`

type CLI struct{}

func NewCLI() *CLI {
	return &CLI{}
}

func (cli *CLI) Run(args []string) (int, error) {
	options := NewOptions()
	flags := flag.NewFlagSet(AppName, flag.ContinueOnError)
	flags.Usage = func() {
		fmt.Println(helpText)
		flags.PrintDefaults()
	}

	// Pagination
	flags.IntVar(&options.Page, "p", 1, "Specify further pages")
	flags.IntVar(&options.PerPage, "n", 100, "Specify a custom page size")

	// Refine
	flags.BoolVar(&options.assigned, "a", false, "Refine issues assigned to you")
	flags.BoolVar(&options.created, "c", false, "Refine issues created by you")
	flags.BoolVar(&options.mentioned, "m", false, "Refine issues mentioning you")

	// State
	flags.StringVar(&options.State, "state", "open",
		"Specify the state of the issues to display. Can be either open, closed, all")

	// Sort
	flags.StringVar(&options.Sort, "sort", "updated",
		"Specify the sort of the issues to display. Can be either created, updated, comments")

	// Format
	flags.StringVar(&options.format, "format", "%n\\t%l\\t%t\\t%u",
		"Specify the format of the issues to display")

	// GitHub personal access token
	flags.StringVar(&options.token, "token", "", "GitHub personal access token")

	version := flags.Bool("v", false, "Show version number and quit")
	help := flags.Bool("h", false, "This help text")

	if err := flags.Parse(args[1:]); err != nil {
		return CodeFlagParseFail, err
	}

	if *version {
		cli.ShowVersion()
		return CodeSuccess, nil
	}

	if *help {
		flags.Usage()
		return CodeSuccess, nil
	}

	issues, _, err := githubIssues(options)
	if err != nil {
		return CodeAPIRequestFail, err
	}

	format := options.Format()
	for _, issue := range issues {
		fmt.Println(format.Apply(issue))
	}
	return CodeSuccess, nil
}

func (cli *CLI) ShowVersion() {
	fmt.Println(strings.Join([]string{AppName, "versin", Version}, " "))
}
