package main

import (
	"os"
	"os/exec"
	"strings"
)

const (
	// Personal access token in .gitconfig
	tokenConfig = "github.token"
	// Personal access token in environments
	tokenEnv = "GITHUB_TOKEN"
)

// GitHub API request common options
type Options struct {
	Page      int
	PerPage   int
	State     string
	Sort      string
	assigned  bool
	created   bool
	mentioned bool
	format    string
	token     string
}

func NewOptions() *Options {
	return &Options{}
}

func (opt *Options) Filter() string {
	switch {
	case opt.assigned:
		return "assigned"
	case opt.created:
		return "created"
	case opt.mentioned:
		return "mentioned"
	default:
		return "all"
	}
}

func (opt *Options) Format() *Format {
	return NewFormat(opt.format)
}

func (opt *Options) Token() string {
	// Token from command line options
	if len(opt.token) > 0 {
		return opt.token
	}

	// Token from environments
	if token := os.Getenv(tokenEnv); len(token) > 0 {
		return token
	}

	// Token from .gitconfig
	out, err := exec.Command("git", "config", tokenConfig).Output()
	if err == nil && len(string(out)) > 0 {
		return strings.TrimSpace(string(out))
	}

	// Returns empty string
	return ""
}
