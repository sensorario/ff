package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

type CompleteBranchStep struct{}

func (s *CompleteBranchStep) Execute(c *Context) bool {
	gitStatus := &GitCommand{[]string{"status"}, "Cant get status"}
	cmdOut := gitStatus.Execute()

	re := regexp.MustCompile(`On branch [\w\/\#\-]{0,}`)

	branchName := ""
	for _, match := range re.FindAllString(string(cmdOut), -1) {
		branchName = strings.ReplaceAll(match, "On branch ", "")
	}

	fmt.Println(color.RedString("leaving: " + branchName))

	gitCheckoutMaster := &GitCommand{[]string{"checkout", "master"}, "Cant checkout master"}
	_ = gitCheckoutMaster.Execute()

	gitDescribeTags := &GitCommand{[]string{"describe", "--tags"}, "cant get tag description"}
	cmdOut = gitDescribeTags.Execute()

	isHotfix := strings.HasPrefix(branchName, "hotfix/")
	isFeature := strings.HasPrefix(branchName, "feature/")

	fmt.Print("current tag: ", color.GreenString(string(cmdOut)))

	tagName := ""

	meta := Meta{string(cmdOut), branchName}

	if isHotfix {
		tagName = meta.NextPatchTag()
	}

	if isFeature {
		tagName = meta.NextMinorTag()
	}

	fmt.Println("next tag:   ", color.RedString(tagName))

	gitMergeNoFF := &GitCommand{[]string{"merge", "--no-ff", branchName}, "cant merge"}
	_ = gitMergeNoFF.Execute()

	gitTag := &GitCommand{[]string{"tag", tagName}, "cant tag"}
	_ = gitTag.Execute()

	gitDeleteOldBranch := &GitCommand{[]string{"branch", "-D", branchName}, "cant merge"}
	_ = gitDeleteOldBranch.Execute()

	fmt.Println(color.GreenString("branch " + branchName + " deleted"))

	c.CurrentStep = &FinalStep{}

	return true
}

func (s *CompleteBranchStep) Stepname() string {
	return "checkout master"
}