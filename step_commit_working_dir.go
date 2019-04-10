package main

import (
	"github.com/fatih/color"
	"regexp"
)

type wokingDirStep struct{}

func (s wokingDirStep) Execute(c *context) bool {
	gitStatus := &gitCommand{
		c.Logger,
		[]string{"status"},
		"Cant get status",
	}

	cmdOut := gitStatus.Execute()

	re := regexp.MustCompile(`(?m)nothing to commit, working tree clean`)

	isWorkingTreeClean := false
	for _, _ = range re.FindAllString(string(cmdOut), -1) {
		isWorkingTreeClean = true
		c.Logger.Info(color.RedString("working tree clean"))
	}

	if isWorkingTreeClean {
		return false
	}

	c.CurrentStep = &commitStep{}

	return true
}

func (s wokingDirStep) Stepname() string {
	return "check-working-directory"
}
