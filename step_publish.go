package main

import (
	"fmt"
	"regexp"
	"strings"
)

type PublishStep struct{}

func (s *PublishStep) Execute(c *Context) bool {
	gitStatus := &GitCommand{[]string{"status"}, "Cant get status"}
	cmdOut := gitStatus.Execute()

	re := regexp.MustCompile(`On branch [\w\/\#\-]{0,}`)

	branchName := ""
	for _, match := range re.FindAllString(string(cmdOut), -1) {
		branchName = strings.ReplaceAll(match, "On branch ", "")
		fmt.Println("Current branch: ", branchName)
	}

	gitPush := &GitCommand{[]string{"push", "origin", branchName}, "cant push"}
	_ = gitPush.Execute()

	return false
}

func (s *PublishStep) Stepname() string {
	return "publish"
}