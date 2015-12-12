package main

const (
	// Personal Access Token in .gitconfig
	personalAccessTokenKey = "github.token"
)

// GitHub API request common options
type Options struct {
	Token string
}

func NewOptions() *Options {
	return &Options{}
}
