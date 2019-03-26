package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
)

type PruneStep struct{}

func (s *PruneStep) Execute(c *Context) bool {
	cmdName := "git"
	cmdArgs := []string{"fetch", "-a", "--prune"}
	if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Println(color.RedString("something went wrong"))
		os.Exit(1)
	}

	c.CurrentStep = &InputReadingStep{}

	return false
}

func (s *PruneStep) Stepname() string {
	return "prune"
}
