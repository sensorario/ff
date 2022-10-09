package ff

import (
    "fmt"
)

type publishStep struct{}

func (s publishStep) Execute(c *Context) bool {
	branchName := c.currentBranch()

	args := []string{"push", "origin", branchName}
	if c.Conf.Features.ForceOnPublish == true {
		args = append(args, "-f")
	}

	if c.Conf.Features.PushTagsOnPublish == true {
		args = append(args, "--tags")
	} else {
        fmt.Println("Any tags will be pushed on listed origin")
    }

	gitPush := &GitCommand{
		c.Logger,
		args,
		"cant push current branch and tags",
		c.Conf,
	}

	_ = gitPush.Execute()

	return false
}

func (s publishStep) Stepname() string {
	return "publish"
}
