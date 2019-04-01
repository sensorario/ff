package main

import (
	"fmt"
	"regexp"

	"github.com/fatih/color"
)

type WorkingDirStep struct{}

func (s WorkingDirStep) Execute(c *Context) bool {
	gitStatus := &GitCommand{
		c.Logger,
		[]string{"status"},
		"Cant get status",
	}

	cmdOut := gitStatus.Execute()

	re := regexp.MustCompile(`(?m)nothing to commit, working tree clean`)

	isWorkingTreeClean := false
	for _, _ = range re.FindAllString(string(cmdOut), -1) {
		isWorkingTreeClean = true
		fmt.Println(color.RedString("working tree clean"))
	}

	if isWorkingTreeClean {
		return false
	}

	c.CurrentStep = &CommitStep{}

	return true
}

func (s WorkingDirStep) Stepname() string {
	return "check-working-directory"
}
