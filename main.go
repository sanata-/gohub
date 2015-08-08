package main

import (
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	x := ""
	fmt.Scan(&x)
	repo_name := x
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "__access__token___"},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	repo := &github.Repository{
		Name: github.String(repo_name),
		//Private: github.Bool(true),
	}
	client.Repositories.Create("", repo)
}
