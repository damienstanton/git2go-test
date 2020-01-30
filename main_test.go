package main

import (
	"testing"

	git "gopkg.in/libgit2/git2go.v27"
)

func TestMain(t *testing.T) {
	repo, err := git.OpenRepository(".")
	if err != nil {
		t.Fatalf("git2go error: %v", err)
	}

	if repo.IsBare() {
		t.Fatal("This repo itself cannot be bare")
	}
}
