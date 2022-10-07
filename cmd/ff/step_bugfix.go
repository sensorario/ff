package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

type bugfixStep struct{}

func (s bugfixStep) Execute(c *context) bool {
	developmentBranch := c.conf.Branches.Historical.Development

	gitCheckoutBackToDev := &gitCommand{
		c.Logger,
		[]string{"checkout", developmentBranch},
		"Cant checkout " + developmentBranch,
		c.conf,
	}
	_ = gitCheckoutBackToDev.Execute()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(color.RedString("Bugfix description: "))
	readedString, _ := reader.ReadString('\n')
	bugfixDescription := slugify(readedString)

	bugfixBranch := "bugfix/" + bugfixDescription + "/" + developmentBranch
	fmt.Println("Bugfix: ", color.YellowString(bugfixBranch))

	gitCheckoutNewBranch := &gitCommand{
		c.Logger,
		[]string{"checkout", "-b", bugfixBranch},
		"Cant create new branch",
		c.conf,
	}

	_ = gitCheckoutNewBranch.Execute()

	c.CurrentStep = &finalStep{}

	return true
}

func (s bugfixStep) Stepname() string {
	return "create-bugfix-branch"
}
