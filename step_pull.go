package main

type pullStep struct{}

func (s pullStep) Execute(c *context) bool {
	branchName := c.currentBranch()

	gitPush := &gitCommand{
		c.Logger,
		[]string{"pull", "origin", branchName, "--tags", "-f"},
		"cant pull current branch and tags",
	}

	_ = gitPush.Execute()

	return false
}

func (s pullStep) Stepname() string {
	return "pull"
}
