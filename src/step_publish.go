package main

type publishStep struct{}

func (s publishStep) Execute(c *context) bool {
	branchName := c.currentBranch()

	args = []string{"push", "origin", branchName}
	if gc.conf.Features.ForceOnPublish == true {
		args = append(args, "-f")
	}

	if gc.conf.Features.PushTagsOnPublish == true {
		args = append(args, "--tags")
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
