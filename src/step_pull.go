package main

import "fmt"

type pullStep struct{}

func (s pullStep) Execute(c *context) bool {
	branchName := c.currentBranch()

	gitPull := &gitCommand{
		c.Logger,
		[]string{"pull", "origin", branchName, "--tags", "-f"},
		"cant pull current branch and tags",
		c.conf,
	}

	output := gitPull.Execute()

	fmt.Println(output)

	return false
}

func (s pullStep) Stepname() string {
	return "pull"
}
