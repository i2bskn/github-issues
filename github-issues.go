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
	}
	app.Action = func(c *cli.Context) {
		options := newOptions(c)
		issues, _, err := githubIssues(options)
		if err != nil {
			fail("Failed get issues")
		}

		for _, issue := range issues {
			fmt.Printf("%d\t%s\t%s\t%s\n", *issue.Number, *issue.HTMLURL, *issue.Title, *issue.User.Login)
		}
	}
	app.Run(os.Args)
}

func fail(message string) {
	fmt.Println(message)
	os.Exit(1)
}
