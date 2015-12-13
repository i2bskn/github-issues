package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
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
	Self          bool
	CurrentRepo   bool
	assigned      bool
	created       bool
	mentioned     bool
	repository    string
	token         string
	obtainedToken string
}

// Regular expression to match the GitHub repository
var reRepo = regexp.MustCompile(`^([^/]+)/([^/]+?)(?:\.git)?$`)
var reURL = regexp.MustCompile(`^(?:(?:ssh://)?git@github\.com(?::|/)|https://github\.com/)([^/]+)/([^/]+?)(?:\.git)?$`)

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
	if token, err := opt.getGitConfig(tokenConfig); err == nil && len(token) > 0 {
		opt.obtainedToken = token
		return token
	}

	// Returns empty string if personal access token not find.
	return ""
}

// Validation of input and settings.
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

func (opt *Options) getOwnerAndRepo() (string, string, error) {
	if len(opt.repository) > 0 {
		matches := reRepo.FindStringSubmatch(opt.repository)

		if len(matches) != 3 {
			return "", "", fmt.Errorf("Failed parse %s", opt.repository)
		}

		return matches[1], matches[2], nil
	}

	url, err := opt.getGitConfig("remote.origin.url")
	if err != nil {
		return "", "", err
	}

	matches := reURL.FindStringSubmatch(url)
	if len(matches) != 3 {
		return "", "", errors.New("Failed parse remote.origin.url")
	}

	return matches[1], matches[2], nil
}

func (opt *Options) getGitConfig(key string) (string, error) {
	out, err := exec.Command("git", "config", "--get", key).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
