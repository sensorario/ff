package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
)

type CommitStep struct{ interactive bool }

func (s *CommitStep) Execute(c *Context) bool {
	text := "default commit message"
	if s.interactive {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("What is the commit message: ")
		text, _ = reader.ReadString('\n')
		fmt.Println(text)
	}

	var err error

	cmdName := "git"

	cmdArgs := []string{"add", "."}
	if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Println(color.RedString("cant add working directory"))
		os.Exit(1)
	}

	if s.interactive {
		cmdArgs = []string{"commit", "-m", text}
	} else {
		cmdArgs = []string{"commit", "-m", "default-commit-message"}
	}
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
