package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

type patchStep struct{}

func (s patchStep) Execute(c *context) bool {
	developmentBranch := c.conf.Branches.Historical.Development

	gitCheckoutToDev := &gitCommand{
		c.Logger,
		[]string{"checkout", developmentBranch},
		"Cant checkout " + developmentBranch,
		c.conf,
	}

	_ = gitCheckoutToDev.Execute()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(color.RedString("Patch descrption: "))
	patchDescription, _ := reader.ReadString('\n')
	patchDescription = strings.ReplaceAll(patchDescription, " ", "-")
	patchDescription = strings.ReplaceAll(patchDescription, "'", "-")
	patchDescription = strings.ReplaceAll(patchDescription, "\n", "")
	patchDescription = strings.ToLower(patchDescription)

	fmt.Print(
		"Patch: ",
	)

	patchBranchName := "refactor/" + patchDescription + "/" + developmentBranch
	fmt.Println(color.YellowString(patchBranchName))

	gitCheckoutNewBranch := &gitCommand{
		c.Logger,
		[]string{"checkout", "-b", patchBranchName},
		"Cant create new branch",
		c.conf,
	}

	_ = gitCheckoutNewBranch.Execute()

	c.CurrentStep = &finalStep{}

	return true
}

func (s patchStep) Stepname() string {
	return "patch-branch"
}
