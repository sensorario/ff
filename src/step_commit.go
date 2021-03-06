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
	fmt.Println(text)

	gitAddAll := &gitCommand{
		c.Logger,
		[]string{"add", "."},
		"Cant add files",
		c.conf,
	}

	_ = gitAddAll.Execute()

	gitCommit := &gitCommand{
		c.Logger,
		[]string{"commit", "-m", text},
		"Cant add files",
		c.conf,
	}

	_ = gitCommit.Execute()

	c.CurrentStep = &finalStep{}

	return false
}

func (s commitStep) Stepname() string {
	return "commit"
}
