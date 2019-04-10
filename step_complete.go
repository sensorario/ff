package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

type completeBranchStep struct{}

func (s completeBranchStep) Execute(c *Context) bool {
	gitStatus := &gitCommand{
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

	br := branch{branchName}
	fmt.Println(color.RedString("destination: " + br.destination()))

	gitCheckoutMaster := &gitCommand{
		c.Logger,
		[]string{"checkout", br.destination()},
		"Cant checkout destination branch",
	}

	_ = gitCheckoutMaster.Execute()

	gitMergeNoFF := &gitCommand{
		c.Logger,
		[]string{"merge", "--no-ff", branchName},
		"cant merge",
	}

	_ = gitMergeNoFF.Execute()

	gitDescribeTags := &gitCommand{
		c.Logger,
		[]string{"describe", "--tags"},
		"cant get tag description",
	}

	cmdOut = gitDescribeTags.Execute()

	fmt.Print("current tag: ", color.GreenString(string(cmdOut)))

	tagName := ""

	meta := Meta{string(cmdOut), branchName}

	if br.isHotfix() || br.isRefactoring() || br.isBugfix() {
		c.Logger.Info("Is Patch branch")
		tagName = meta.NextPatchTag()
	}

	if br.isFeature() {
		c.Logger.Info("Is Feature branch")
		tagName = meta.NextMinorTag()
	}

	fmt.Println("next tag:   ", color.GreenString(tagName))

	gitTag := &gitCommand{
		c.Logger,
		[]string{"tag", tagName, "-f"},
		"cant tag",
	}
	_ = gitTag.Execute()

	gitDeleteOldBranch := &gitCommand{
		c.Logger,
		[]string{"branch", "-D", branchName},
		"cant merge",
	}
	_ = gitDeleteOldBranch.Execute()

	fmt.Println(color.GreenString("branch " + branchName + " deleted"))

	if br.destination() != "master" {
		gitCheckoutMaster := &gitCommand{
			c.Logger,
			[]string{"checkout", "master"},
			"Cant checkout destination branch",
		}
		_ = gitCheckoutMaster.Execute()

		gitMergeNoFastForward := &gitCommand{
			c.Logger,
			[]string{"merge", "--no-ff", br.destination()},
			"cant move to master updates",
		}
		_ = gitMergeNoFastForward.Execute()
	}

	c.CurrentStep = &FinalStep{}

	return true
}

func (s completeBranchStep) Stepname() string {
	return "checkout-master"
}
