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

	re := regexp.MustCompile(`On branch [\w\/\#\-]{0,}`)

	branchName := ""
	for _, match := range re.FindAllString(string(cmdOut), -1) {
		branchName = strings.ReplaceAll(match, "On branch ", "")
	}

	fmt.Println(color.RedString("leaving: " + branchName))

	branch := SemBranch{branchName}
	fmt.Println(color.RedString("destination: " + branch.Destination()))

	gitCheckoutMaster := &GitCommand{
		c.Logger,
		[]string{
			"checkout",
			branch.Destination(),
		},
		"Cant checkout destination branch",
	}
	_ = gitCheckoutMaster.Execute()

	gitDescribeTags := &GitCommand{
		c.Logger,
		[]string{"describe", "--tags"},
		"cant get tag description",
	}
	cmdOut = gitDescribeTags.Execute()

	isHotfix := strings.HasPrefix(branchName, "hotfix/")
	isFeature := strings.HasPrefix(branchName, "feature/")

	fmt.Print("current tag: ", color.GreenString(string(cmdOut)))

	tagName := ""

	meta := Meta{string(cmdOut), branchName}

	if isHotfix {
		c.Logger.Info("Is Hotfix")
		c.Logger.Info("patch" + meta.PatchVersion())
		c.Logger.Info("minor" + meta.MinorVersion())
		c.Logger.Info("major" + meta.MajorVersion())
		tagName = meta.NextPatchTag()
		c.Logger.Info("NextPatchTag")
		c.Logger.Info("tagName: " + tagName)
		c.Logger.Info(meta.IncPatchVersion())
	}

	if isFeature {
		c.Logger.Info("Is Feature")
		tagName = meta.NextMinorTag()
	}

	fmt.Println("next tag:   ", color.RedString(tagName))

	gitMergeNoFF := &GitCommand{
		c.Logger,
		[]string{"merge", "--no-ff", branchName},
		"cant merge",
	}
	_ = gitMergeNoFF.Execute()

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
