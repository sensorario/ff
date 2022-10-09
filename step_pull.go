package ff

import "fmt"

type pullStep struct{}

func (s pullStep) Execute(c *Context) bool {
	branchName := c.currentBranch()

	gitPull := &GitCommand{
		c.Logger,
		[]string{"pull", "origin", branchName, "--tags", "-f"},
		"cant pull current branch and tags",
		c.Conf,
	}

	output := gitPull.Execute()

	fmt.Println(output)

	return false
}

func (s pullStep) Stepname() string {
	return "pull"
}
