package main

import "fmt"

type fetchAllStep struct{}

func (s fetchAllStep) Execute(c *context) bool {
	args := []string{
		"fetch",
		"--all",
	}

	gitFetchAll := &gitCommand{
		c.Logger,
		args,
		"Cant fetch all",
		c.conf,
	}

	out := gitFetchAll.Execute()

	fmt.Println(out)

	c.CurrentStep = &finalStep{}

	return true
}

func (s fetchAllStep) Stepname() string {
	return "fetch-all"
}
