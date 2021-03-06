package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

type hotfixStep struct{}

func (s hotfixStep) Execute(c *context) bool {
	developmentBranch := c.currentBranch()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(color.RedString("Hotfix description: "))
	readedString, _ := reader.ReadString('\n')
	hotfixDescription := slugify(readedString)

	hotfixBranch := "hotfix/" + hotfixDescription + "/" + developmentBranch
	fmt.Println("Hotfix: ", color.YellowString(hotfixBranch))

	gitCheckoutNewBranch := &gitCommand{
		c.Logger,
		[]string{"checkout", "-b", hotfixBranch},
		"Cant create new branch",
		c.conf,
	}

	_ = gitCheckoutNewBranch.Execute()

	c.CurrentStep = &finalStep{}

	return true
}

func (s hotfixStep) Stepname() string {
	return "create-hotfix-branch"
}
