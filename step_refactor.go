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
	developmentBranch := "master"

	gitCheckoutMaster := &gitCommand{
		c.Logger,
		[]string{"checkout", developmentBranch},
		"Cant checkout master",
	}

	_ = gitCheckoutMaster.Execute()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(color.RedString("Refactoring descrption: "))
	featureName, _ := reader.ReadString('\n')
	featureName = strings.ReplaceAll(featureName, " ", "-")
	featureName = strings.ReplaceAll(featureName, "\n", "")

	fmt.Print(
		"Feature: ",
	)

	featureBranchName := "refactor/" + featureName + "/" + developmentBranch
	fmt.Println(color.YellowString(featureBranchName))

	gitCheckoutNewBranch := &gitCommand{
		c.Logger,
		[]string{"checkout", "-b", featureBranchName},
		"Cant create new branch",
	}

	_ = gitCheckoutNewBranch.Execute()

	c.CurrentStep = &finalStep{}

	return true
}

func (s refactoringStep) Stepname() string {
	return "create-feature-branch"
}
