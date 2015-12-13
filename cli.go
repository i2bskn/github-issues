package main

import (
	"flag"
	"fmt"
	"io"
)

// Application information
const (
	AppName = "github-issues"
	Version = "0.1.0"
)

// Return code of github-issues command
const (
	ExitCodeOK int = 0

	ExitCodeFlagParseError = 10 + iota
	ExitCodeInvalidOptions
	ExitCodeAPIRequestError
)

// CLI is application object.
type CLI struct {
	outStream, errStream io.Writer
}

// Run to execute application from arguments.
func (cli *CLI) Run(args []string) int {
	var version, help bool
	var baseFormat string
	options := NewOptions()

	flags := flag.NewFlagSet(AppName, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)
	flags.Usage = func() {
		fmt.Fprint(cli.outStream, usage)
	}

	// Repository
	flags.StringVar(&options.repository, "repo", "", "")
	flags.BoolVar(&options.CurrentRepo, "current", false, "")
	flags.BoolVar(&options.Self, "self", false, "")

	// Pagination
	flags.IntVar(&options.Page, "page", 1, "")
	flags.IntVar(&options.PerPage, "per-page", 100, "")

	// Refine
	flags.BoolVar(&options.assigned, "assigned", false, "")
	flags.BoolVar(&options.created, "created", false, "")
	flags.BoolVar(&options.mentioned, "mentioned", false, "")

	// State
	flags.StringVar(&options.State, "state", "open", "")

	// Sort
	flags.StringVar(&options.Sort, "sort", "updated", "")

	// Format
	flags.StringVar(&baseFormat, "format", defaultFormat, "")

	// GitHub personal access token
	flags.StringVar(&options.token, "token", "", "")

	flags.BoolVar(&version, "version", false, "")
	flags.BoolVar(&help, "help", false, "")

	if err := flags.Parse(args[1:]); err != nil {
		fmt.Fprintln(cli.errStream, err)
		return ExitCodeFlagParseError
	}

	if version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", AppName, Version)
		return ExitCodeOK
	}

	if help {
		flags.Usage()
		return ExitCodeOK
	}

	// Options validation
	if err := options.Validation(); err != nil {
		fmt.Fprintln(cli.errStream, err)
		return ExitCodeInvalidOptions
	}

	githubIssue := NewIssues(options.Token())
	issues, _, err := githubIssue.ListByOptions(options)
	if err != nil {
		fmt.Fprintln(cli.errStream, err)
		return ExitCodeAPIRequestError
	}

	format := NewFormat(baseFormat)
	for _, issue := range issues {
		fmt.Fprintln(cli.outStream, format.Apply(issue))
	}
	return ExitCodeOK
}

const usage = `
Usage: github-issues [OPTIONS]

List of GitHub issues.

Options:

  -current            Current repository only
  -repo=<owner/repo>  Specific repository only
  -self               Your own repositories only

  -page               Specify further pages (default: 1)
  -per-page           Specify a custom page size (default: 100)

  -assigned           Refine issues assigned to you
  -created            Refine issues created by you
  -mentioned          Refine issues mentioning you

  -state              Specify the state of the issues to display.
                      Can be either open, closed, all
  -sort               Specify the sort of the issues to display.
                      Can be either created, updated, comments
  -format             Specify the format of the issues to display
  -token              GitHub personal access token

  -version            Show version number and quit
  -help               This help text
`
