package main

type pullStep struct{}

func (s pullStep) Execute(c *context) bool {
	branchName := c.currentBranch()

	gitPull := &gitCommand{
		c.Logger,
		[]string{"pull", "origin", branchName, "--tags", "-f"},
		"cant pull current branch and tags",
		c.conf,
	}

	_ = gitPull.Execute()

	return false
}

func (s pullStep) Stepname() string {
	return "pull"
}
