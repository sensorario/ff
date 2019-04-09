package main

type PublishStep struct{}

func (s PublishStep) Execute(c *Context) bool {
	branchName := c.CurrentBranch()

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
