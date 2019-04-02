package main

import (
	"regexp"
	"strings"

	"github.com/sensorario/gol"
)

type Context struct {
	CurrentStep FussyStepInterface
	Exit        bool
	Logger      gol.Logger
}

func (c Context) CurrentBranch() string {
	gitStatus := &GitCommand{
		c.Logger,
		[]string{"status"},
		"Cant get status",
	}

	cmdOut := gitStatus.Execute()

	re := regexp.MustCompile(`On branch [\w\/\#\-]{0,}`)

	branchName := ""

	for _, match := range re.FindAllString(string(cmdOut), -1) {
		branchName = strings.ReplaceAll(match, "On branch ", "")
	}

	return branchName
}
