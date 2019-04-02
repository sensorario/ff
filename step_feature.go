package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

type FeatureStep struct{}

func (s FeatureStep) Execute(c *Context) bool {
	developmentBranch := "master"

	gitCheckoutMaster := &GitCommand{
		c.Logger,
		[]string{"checkout", developmentBranch},
		"Cant checkout master",
	}

	_ = gitCheckoutMaster.Execute()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(color.RedString("New feature description: "))
	featureName, _ := reader.ReadString('\n')
	featureName = strings.ReplaceAll(featureName, " ", "-")
	featureName = strings.ReplaceAll(featureName, "\n", "")

	fmt.Print(
		"Feature: ",
	)

	featureBranchName := "feature/" + featureName + "/" + developmentBranch
	fmt.Println(color.YellowString(featureBranchName))

	gitCheckoutNewBranch := &GitCommand{
		c.Logger,
		[]string{"checkout", "-b", featureBranchName},
		"Cant create new branch",
	}

	_ = gitCheckoutNewBranch.Execute()

	c.CurrentStep = &FinalStep{}

	return true
}

func (s FeatureStep) Stepname() string {
	return "create-feature-branch"
}
