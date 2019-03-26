package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
)

type CommitStep struct{}

func (s *CommitStep) Execute(c *Context) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is the commit message: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	var err error

	cmdName := "git"

	cmdArgs := []string{"add", "."}
	if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Println(color.RedString("cant add working directory"))
		os.Exit(1)
	}

	cmdArgs = []string{"commit", "-m", text}
	if _, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Println(color.RedString("something went wrong"))
		os.Exit(1)
	}

	c.CurrentStep = &FinalStep{}

	return false
}

func (s *CommitStep) Stepname() string {
	return "commit"
}
