package main

import (
	"fmt"
	"regexp"
	"strings"
)

type authorsStep struct{}

func (s authorsStep) Execute(c *context) bool {
	gitStatus := &gitCommand{
		c.Logger,
		[]string{"status"},
		"cant get status",
		c.conf,
	}
	cmdOut := gitStatus.Execute()
	branchName := ""
	re := regexp.MustCompile(`On branch [\w\/\#\-\.]{0,}`)
	for _, match := range re.FindAllString(string(cmdOut), -1) {
		branchName = strings.ReplaceAll(match, "On branch ", "")
	}

	args := []string{
		"shortlog",
		branchName,
		"--summary",
		"--numbered",
	}

	gitShortLog := &gitCommand{
		c.Logger,
		args,
		"Cant lista authors",
		c.conf,
	}

	out := gitShortLog.Execute()

	fmt.Println(out)

	c.CurrentStep = &finalStep{}

	return true
}

func (s authorsStep) Stepname() string {
	return "repository-authors"
}
