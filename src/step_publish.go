package main

type publishStep struct{}

func (s publishStep) Execute(c *context) bool {
	branchName := c.currentBranch()

	gitPush := &gitCommand{
		c.Logger,
		[]string{"push", "origin", branchName, "--tags", "-f"},
		"cant push current branch and tags",
	}

	_ = gitPush.Execute()

	return false
}

func (s publishStep) Stepname() string {
	return "publish"
}
