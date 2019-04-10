package main

type PublishStep struct{}

func (s PublishStep) Execute(c *context) bool {
	branchName := c.currentBranch()

	gitPush := &gitCommand{
		c.Logger,
		[]string{"push", "origin", branchName, "--tags", "-f"},
		"cant push current branch and tags",
	}

	_ = gitPush.Execute()

	return false
}

func (s PublishStep) Stepname() string {
	return "publish"
}
