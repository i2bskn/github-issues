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
	CodeSuccess       = 0
	CodeFlagParseFail = 101
)

type CLI struct{}

func NewCLI() *CLI {
	return &CLI{}
}

func (cli *CLI) Run(args []string) (int, error) {
	options := NewOptions()
	flags := flag.NewFlagSet(AppName, flag.ContinueOnError)

	version := flags.Bool("v", false, "Show version")

	if err := flags.Parse(args[1:]); err != nil {
		return CodeFlagParseFail, err
	}

	if *version {
		cli.ShowVersion()
		return CodeSuccess, nil
	}

	fmt.Printf("%v\n", options) // debug
	return CodeSuccess, nil
}

func (cli *CLI) ShowVersion() {
	fmt.Println(strings.Join([]string{AppName, "versin", Version}, " "))
}
