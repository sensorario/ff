package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

type HotfixStep struct{}

func (s HotfixStep) Execute(c *Context) bool {
	developmentBranch := "master"

	gitCheckoutMaster := &GitCommand{
		c.Logger,
		[]string{"checkout", "master"},
		"Cant checkout master",
	}
	_ = gitCheckoutMaster.Execute()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(color.RedString("Hotfix description: "))
	hotfixDescription, _ := reader.ReadString('\n')
	hotfixDescription = strings.ReplaceAll(hotfixDescription, " ", "-")
	hotfixDescription = strings.ReplaceAll(hotfixDescription, "\n", "")

	fmt.Print(
		"Hotfix: ",
	)

	hotfixBranch := "hotfix/" + hotfixDescription + "/" + developmentBranch
	fmt.Println(color.YellowString(hotfixBranch))

	gitCheckoutNewBranch := &GitCommand{c.Logger, []string{"checkout", "-b", hotfixBranch}, "Cant create new branch"}
	_ = gitCheckoutNewBranch.Execute()

	c.CurrentStep = &FinalStep{}

	return true
}

func (s HotfixStep) Stepname() string {
	return "create-hotfix-branch"
}
