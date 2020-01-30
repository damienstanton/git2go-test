package main

import (
	"fmt"

	git "gopkg.in/libgit2/git2go.v27"
)

func main() {
	repo, err := git.OpenRepository(".")
	if err != nil {
		fmt.Errorf("git2go error: %v", err)
	}

	if repo.IsBare() {
		fmt.Println("Repo is bare")
	} else {
		fmt.Println("Repo is not bare")
	}
}
