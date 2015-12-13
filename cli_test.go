package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRun__versionFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("github-issues -version", " ")

	status := cli.Run(args)
	if status != ExitCodeOK {
		t.Fatalf("Expected %v, but %v", ExitCodeOK, status)
	}

	expected := fmt.Sprintf("%s version %s", AppName, Version)
	if !strings.Contains(errStream.String(), expected) {
		t.Fatalf("Expected %v, but %v", expected, errStream.String())
	}
}
