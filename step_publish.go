package main

type PublishStep struct{}

func (s PublishStep) Execute(c *Context) bool {
	branchName := c.CurrentBranch()

	gitPush := &gitCommand{
		c.Logger,
		[]string{"push", "origin", branchName, "--tags"},
		"cant push",
	}

	_ = gitPush.Execute()

	return false
}

func (s PublishStep) Stepname() string {
	return "publish"
}
