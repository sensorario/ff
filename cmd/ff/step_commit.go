package main

import (
	"bufio"
	"fmt"
	"os"
)

type commitStep struct{}

func (s commitStep) Execute(c *context) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is the commit message: ")

	text, _ := reader.ReadString('\n')

	gitAddAll := &gitCommand{
		c.Logger,
		[]string{"add", "."},
		"Cant add files",
		c.conf,
	}

	gitAddAll.Execute()

	name := c.currentBranch()
	sem := branch{name}

    result := sem.commitPrefix() + text
	gitCommit := &gitCommand{
		c.Logger,
		[]string{"commit", "-m", result},
		"Cant add more files",
		c.conf,
	}

	gitCommit.Execute()

	c.CurrentStep = &finalStep{}

	return false
}

func (s commitStep) Stepname() string {
	return "commit"
}
