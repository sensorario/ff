package main

import "fmt"

type authorsStep struct{}

func (s authorsStep) Execute(c *context) bool {
	args := []string{
		"shortlog",
		c.conf.Branches.Historical.Development,
		"--summary",
		"--numbered",
	}

	gitShortLog := &gitCommand{
		c.Logger,
		args,
		"Cant lista authors",
		c.conf,
	}

	out := gitShortLog.Execute()

	fmt.Println(out)

	c.CurrentStep = &finalStep{}

	return true
}

func (s authorsStep) Stepname() string {
	return "repository-authors"
}
