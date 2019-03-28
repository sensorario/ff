package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

type StatusStep struct{}

func (s *StatusStep) Execute(c *Context) bool {
	gitStatus := &GitCommand{[]string{"status"}, "Cant get status"}
	cmdOut := gitStatus.Execute()

	re := regexp.MustCompile(`On branch [\w\/\#\-]{0,}`)

	isDevelopmentBranch := false

	for _, match := range re.FindAllString(string(cmdOut), -1) {
		branchName := strings.ReplaceAll(match, "On branch ", "")
		fmt.Println("Current branch: ", branchName)
		isDevelopmentBranch = branchName == "master"
	}

	if isDevelopmentBranch {
		fmt.Println(color.GreenString("Development branch"))
		fmt.Println("you can create feature branch from here")
	}

	return false
}

func (s *StatusStep) Stepname() string {
	return "status"
}
