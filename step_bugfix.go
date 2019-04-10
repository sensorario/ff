package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

type bugfixStep struct{}

func (s bugfixStep) Execute(c *Context) bool {
	developmentBranch := "master"

	gitCheckoutMaster := &gitCommand{
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

	gitCheckoutNewBranch := &gitCommand{
		c.Logger,
		[]string{"checkout", "-b", bugfixBranch},
		"Cant create new branch",
	}

	_ = gitCheckoutNewBranch.Execute()

	c.CurrentStep = &finalStep{}

	return true
}

func (s bugfixStep) Stepname() string {
	return "create-bugfix-branch"
}
