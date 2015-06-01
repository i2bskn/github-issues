package main

import (
	"github.com/google/go-github/github"
	"strconv"
	"strings"
)

type Format struct {
	base string
}

func newFormat(base string) *Format {
	return &Format{
		base: base,
	}
}

func (f Format) Apply(issue github.Issue) string {
	line := f.base
	items := map[string]string{
		"%n": strconv.Itoa(*issue.Number),
		"%l": *issue.HTMLURL,
		"%t": *issue.Title,
		"%u": *issue.User.Login,
	}

	for symbol, value := range items {
		line = strings.Replace(line, symbol, value, -1)
	}
	return line
}