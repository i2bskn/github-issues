package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// GitHub personal access token
const (
	tokenConfig = "github.token"
	tokenEnv    = "GITHUB_TOKEN"
)

// Options is GitHub API request common options.
type Options struct {
	Page          int
	PerPage       int
	State         string
	Sort          string
	assigned      bool
	created       bool
	mentioned     bool
	format        string
	token         string
	obtainedToken string
}

// NewOptions to generate common options.
func NewOptions() *Options {
	return &Options{}
}

// Filter issues by you operation.
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

// Format to generate format object from specify string.
func (opt *Options) Format() *Format {
	return NewFormat(opt.format)
}

// Token to read token from arguments or environments or gitconfig.
func (opt *Options) Token() string {
	if len(opt.obtainedToken) > 0 {
		return opt.obtainedToken
	}

	// Token from command line options
	if len(opt.token) > 0 {
		opt.obtainedToken = opt.token
		return opt.token
	}

	// Token from environments
	if token := os.Getenv(tokenEnv); len(token) > 0 {
		opt.obtainedToken = token
		return token
	}

	// Token from .gitconfig
	out, err := exec.Command("git", "config", tokenConfig).Output()
	if err == nil && len(string(out)) > 0 {
		token := strings.TrimSpace(string(out))
		opt.obtainedToken = token
		return token
	}

	// Returns empty string if personal access token not find.
	return ""
}

func (opt *Options) Validation() error {
	if len(opt.Token()) < 1 {
		return errors.New("Personal access token can not obtain.")
	}

	if opt.invalidState() {
		return fmt.Errorf("Invalid state: %s", opt.State)
	}

	if opt.invalidSort() {
		return fmt.Errorf("Invalid sort: %s", opt.Sort)
	}

	return nil
}

func (opt *Options) invalidState() bool {
	for _, state := range [...]string{"open", "closed", "all"} {
		if opt.State == state {
			return false
		}
	}
	return true
}

func (opt *Options) invalidSort() bool {
	for _, sort := range [...]string{"created", "updated", "comments"} {
		if opt.Sort == sort {
			return false
		}
	}
	return true
}
