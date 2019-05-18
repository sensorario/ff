package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

type refactoringStep struct{}

func (s refactoringStep) Execute(c *context) bool {
	developmentBranch := c.conf.Branches.Historical.Development

	gitCheckoutToDev := &gitCommand{
		c.Logger,
		[]string{"checkout", developmentBranch},
		"Cant checkout " + developmentBranch,
		c.conf,
	}

	_ = gitCheckoutToDev.Execute()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(color.RedString("Refactoring descrption: "))
	featureName, _ := reader.ReadString('\n')
	featureName = strings.ReplaceAll(featureName, " ", "-")
	featureName = strings.ReplaceAll(featureName, "'", "-")
	featureName = strings.ReplaceAll(featureName, "\n", "")
	featureName = strings.ToLower(featureName)

	fmt.Print(
		"Feature: ",
	)

	featureBranchName := "refactor/" + featureName + "/" + developmentBranch
	fmt.Println(color.YellowString(featureBranchName))

	gitCheckoutNewBranch := &gitCommand{
		c.Logger,
		[]string{"checkout", "-b", featureBranchName},
		"Cant create new branch",
		c.conf,
	}

	_ = gitCheckoutNewBranch.Execute()

	c.CurrentStep = &finalStep{}

	return true
}

func (s refactoringStep) Stepname() string {
	return "create-feature-branch"
}
