package main

import (
	"fmt"
	"os"
)

func main() {
	cli := NewCLI()
	code, err := cli.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
	}
	os.Exit(code)
}
