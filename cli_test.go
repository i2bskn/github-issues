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

	exitCode := cli.Run(args)
	if exitCode != ExitCodeOK {
		t.Fatalf("Expected %v, but %v", ExitCodeOK, exitCode)
	}

	expected := fmt.Sprintf("%s version %s", AppName, Version)
	if !strings.Contains(errStream.String(), expected) {
		t.Fatalf("Expected %v, but %v", expected, errStream.String())
	}
}

func TestRun__helpFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("github-issues -help", " ")

	exitCode := cli.Run(args)
	if exitCode != ExitCodeOK {
		t.Fatalf("Expected %v, but %v", ExitCodeOK, exitCode)
	}

	if !strings.Contains(outStream.String(), usage) {
		t.Fatalf("Expected %v, but %v", usage, outStream.String())
	}
}

func TestRun__parseError(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("github-issues -invalid", " ")

	exitCode := cli.Run(args)
	if exitCode != ExitCodeFlagParseError {
		t.Fatalf("Expected %v, but %v", ExitCodeFlagParseError, exitCode)
	}

	if !strings.Contains(outStream.String(), usage) {
		t.Fatalf("Expected %v, but %v", usage, outStream.String())
	}
}
