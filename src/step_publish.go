package main

import (
    "fmt"
)

type publishStep struct{}

func (s publishStep) Execute(c *context) bool {
	branchName := c.currentBranch()

	args := []string{"push", "origin", branchName}
	if c.conf.Features.ForceOnPublish == true {
		args = append(args, "-f")
	}

	if c.conf.Features.PushTagsOnPublish == true {
		args = append(args, "--tags")
	} else {
        // @todo keep sentence from a dictionary
        fmt.Println("Any tags will be pushed on listed origin")
    }

	gitPush := &gitCommand{
		c.Logger,
		args,
		"cant push current branch and tags",
		c.conf,
	}

	_ = gitPush.Execute()

	return false
}

func (s publishStep) Stepname() string {
	return "publish"
}
