package main

import (
	"bufio"
	"fmt"
	"os"
)

type CommitStep struct{}

func (s *CommitStep) Execute(c *Context) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is the commit message: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	gitAddAll := &GitCommand{
		c.Logger,
		[]string{"add", "."},
		"Cant add files",
	}

	_ = gitAddAll.Execute()

	gitCommit := &GitCommand{
		c.Logger,
		[]string{"commit", "-m", text},
		"Cant add files",
	}

	_ = gitCommit.Execute()

	c.CurrentStep = &FinalStep{}

	return false
}

func (s *CommitStep) Stepname() string {
	return "commit"
}
