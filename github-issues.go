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
			Usage: "Specific pages",
		},
		cli.IntFlag{
			Name:  "per-page, n",
			Usage: "Specific pages",
		},
		cli.StringFlag{
			Name:  "format, f",
			Usage: "Display format",
		},
	}
	app.Action = func(c *cli.Context) {
		options := newOptions(c)
		format := newFormat(options.format)

		issues, _, err := githubIssues(options)
		if err != nil {
			fail("Failed get issues")
		}

		for _, issue := range issues {
			fmt.Println(format.Apply(issue))
		}
	}
	app.Run(os.Args)
}

func fail(message string) {
	fmt.Println(message)
	os.Exit(1)
}
