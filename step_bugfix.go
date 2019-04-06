package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

type BugfixStep struct{}

func (s BugfixStep) Execute(c *Context) bool {
	developmentBranch := "master"

	gitCheckoutMaster := &GitCommand{
		c.Logger,
		[]string{"checkout", "master"},
		"Cant checkout master",
	}
	_ = gitCheckoutMaster.Execute()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(color.RedString("Bugfix description: "))
	bugfixDescription, _ := reader.ReadString('\n')
	bugfixDescription = strings.ReplaceAll(
		bugfixDescription,
		" ",
		"-",
	)
	bugfixDescription = strings.ReplaceAll(
		bugfixDescription,
		"\n",
		"",
	)

	bugfixBranch := "bugfix/" + bugfixDescription + "/" + developmentBranch
	fmt.Println("Bugfix: ", color.YellowString(bugfixBranch))

	gitCheckoutNewBranch := &GitCommand{
		c.Logger,
		[]string{"checkout", "-b", bugfixBranch},
		"Cant create new branch",
	}

	_ = gitCheckoutNewBranch.Execute()

	c.CurrentStep = &FinalStep{}

	return true
}

func (s BugfixStep) Stepname() string {
	return "create-bugfix-branch"
}
