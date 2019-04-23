package main

import "fmt"

type treeStep struct{}

func (s treeStep) Execute(c *context) bool {
	args := []string{
		"log",
		"--graph",
		"--all",
		"--decorate",
		"--oneline",
		"-10",
	}

	gitTree := &gitCommand{
		c.Logger,
		args,
		"Cant lista authors",
	}

	out := gitTree.Execute()

	fmt.Println(out)

	c.CurrentStep = &finalStep{}

	return true
}

func (s treeStep) Stepname() string {
	return "tree"
}
