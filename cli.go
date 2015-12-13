package main

import (
	"flag"
	"fmt"
	"strings"
)

// Application information
const (
	AppName = "github-issues"
	Version = "0.1.0"
)

// Return code of github-issues command
const (
	CodeSuccess = 0

	CodeFlagParseFail = 101 + iota
	CodeInvalidOptions
	CodeAPIRequestFail
)

const helpHeader = `Usage: github-issues [OPTIONS]

List of GitHub issues.

Options:
`

// CLI is application object.
type CLI struct{}

// NewCLI to generate application object.
func NewCLI() *CLI {
	return &CLI{}
}

// Run to execute application from arguments.
func (cli *CLI) Run(args []string) (int, error) {
	options := NewOptions()
	flags := flag.NewFlagSet(AppName, flag.ContinueOnError)
	flags.Usage = func() {
		fmt.Println(helpHeader)
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

	// Options validation
	if err := options.Validation(); err != nil {
		return CodeInvalidOptions, err
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

// ShowVersion to display the version number.
func (cli *CLI) ShowVersion() {
	fmt.Println(strings.Join([]string{AppName, "versin", Version}, " "))
}
