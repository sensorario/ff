package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

type StatusStep struct{}

func (s StatusStep) Execute(c *Context) bool {
	gitStatus := &GitCommand{c.Logger, []string{"status"}, "Cant get status"}
	cmdOut := gitStatus.Execute()

	re := regexp.MustCompile(`On branch [\w\/\#\-]{0,}`)

	for _, match := range re.FindAllString(string(cmdOut), -1) {
		branchName := strings.ReplaceAll(match, "On branch ", "")
		fmt.Println("Current branch is ", color.GreenString(branchName))
	}

	return false
}

func (s StatusStep) Stepname() string {
	return "status"
}
