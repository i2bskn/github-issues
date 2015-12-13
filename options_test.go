package main

import (
	"testing"
)

func TestGetOwnerAndRepo__ownerAndRepo(t *testing.T) {
	options := NewOptions()
	options.repository = "i2bskn/github-issues"

	owner, repo, err := options.getOwnerAndRepo()
	if owner != "i2bskn" {
		t.Fatalf("Expected %v, but %v", "i2bskn", owner)
	}

	if repo != "github-issues" {
		t.Fatalf("Expected %v, but %v", "github-issues", repo)
	}

	if err != nil {
		t.Fatal(err)
	}
}
