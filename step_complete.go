package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

type CompleteBranchStep struct{}

func (s CompleteBranchStep) Execute(c *Context) bool {
	gitStatus := &GitCommand{
		args:    []string{"status"},
		message: "Cant get status",
		Logger:  c.Logger,
	}

	cmdOut := gitStatus.Execute()

	re := regexp.MustCompile(`On branch [\w\/\#\-\.]{0,}`)

	branchName := ""
	for _, match := range re.FindAllString(string(cmdOut), -1) {
		branchName = strings.ReplaceAll(match, "On branch ", "")
	}

	fmt.Println(color.RedString("leaving: " + branchName))

	branch := Branch{branchName}
	fmt.Println(color.RedString("destination: " + branch.Destination()))

	gitCheckoutMaster := &GitCommand{
		c.Logger,
		[]string{"checkout", branch.Destination()},
		"Cant checkout destination branch",
	}

	_ = gitCheckoutMaster.Execute()

	gitMergeNoFF := &GitCommand{
		c.Logger,
		[]string{"merge", "--no-ff", branchName},
		"cant merge",
	}

	_ = gitMergeNoFF.Execute()

	gitDescribeTags := &GitCommand{
		c.Logger,
		[]string{"describe", "--tags"},
		"cant get tag description",
	}

	cmdOut = gitDescribeTags.Execute()

	fmt.Print("current tag: ", color.GreenString(string(cmdOut)))

	tagName := ""

	meta := Meta{string(cmdOut), branchName}

	if branch.IsHotfix() || branch.IsRefactoring() || branch.IsBugfix() {
		c.Logger.Info("Is Patch branch")
		tagName = meta.NextPatchTag()
	}

	if branch.IsFeature() {
		c.Logger.Info("Is Feature branch")
		tagName = meta.NextMinorTag()
	}

	fmt.Println("next tag:   ", color.GreenString(tagName))

	gitTag := &GitCommand{
		c.Logger,
		[]string{"tag", tagName},
		"cant tag",
	}
	_ = gitTag.Execute()

	gitDeleteOldBranch := &GitCommand{
		c.Logger,
		[]string{"branch", "-D", branchName},
		"cant merge",
	}
	_ = gitDeleteOldBranch.Execute()

	fmt.Println(color.GreenString("branch " + branchName + " deleted"))

	c.CurrentStep = &FinalStep{}

	return true
}

func (s CompleteBranchStep) Stepname() string {
	return "checkout-master"
}