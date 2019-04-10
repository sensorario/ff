package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

type hotfixStep struct{}

func (s hotfixStep) Execute(c *context) bool {
	developmentBranch := c.currentBranch()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(color.RedString("Hotfix description: "))
	hotfixDescription, _ := reader.ReadString('\n')
	hotfixDescription = strings.ReplaceAll(
		hotfixDescription,
		" ",
		"-",
	)
	hotfixDescription = strings.ReplaceAll(
		hotfixDescription,
		"\n",
		"",
	)

	hotfixBranch := "hotfix/" + hotfixDescription + "/" + developmentBranch
	fmt.Println("Hotfix: ", color.YellowString(hotfixBranch))

	gitCheckoutNewBranch := &gitCommand{
		c.Logger,
		[]string{"checkout", "-b", hotfixBranch},
		"Cant create new branch",
	}

	_ = gitCheckoutNewBranch.Execute()

	c.CurrentStep = &finalStep{}

	return true
}

func (s hotfixStep) Stepname() string {
	return "create-hotfix-branch"
}
