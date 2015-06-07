package main

import (
	"fmt"
	"os"
)

type GithubIssuesError struct {
	message string
}

func newError(message string) *GithubIssuesError {
	return &GithubIssuesError{
		message: message,
	}
}

func (err GithubIssuesError) Error() string {
	return err.message
}

func fail(message string) {
	fmt.Println(message)
	os.Exit(1)
}
