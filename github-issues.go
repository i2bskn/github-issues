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
			Usage: "Specify further pages.(default: 1)",
		},
		cli.IntFlag{
			Name:  "per-page, n",
			Usage: "Specify a custom page size.(default: 100)",
		},
		cli.BoolFlag{
			Name:  "assigned, a",
			Usage: "Issues assigned to you.(default: all)",
		},
		cli.BoolFlag{
			Name:  "created, c",
			Usage: "Issues created by you.(default: all)",
		},
		cli.BoolFlag{
			Name:  "mentioned, m",
			Usage: "Issues mentioning you.(default: all)",
		},
		cli.BoolFlag{
			Name:  "closed",
			Usage: "Closed issue only.(default: open)",
		},
		cli.BoolFlag{
			Name:  "all",
			Usage: "With closed issue.(default: open)",
		},
		cli.StringFlag{
			Name:  "format, f",
			Usage: "Specify display format.",
		},
	}
	app.Action = func(c *cli.Context) {
		options := newOptions(c)
		issues, _, err := githubIssues(options)
		if err != nil {
			fail("Failed get issues")
		}

		for _, issue := range issues {
			fmt.Println(options.format.Apply(issue))
		}
	}
	app.Run(os.Args)
}

func fail(message string) {
	fmt.Println(message)
	os.Exit(1)
}
