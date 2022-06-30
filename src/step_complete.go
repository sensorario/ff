package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

type completeBranchStep struct{}

func (s completeBranchStep) Execute(c *context) bool {
    // @todo refactor here is mandatory
	re := regexp.MustCompile(`On branch [\w\/\#\-\.]{0,}`)
	branchName := ""
	for _, match := range re.FindAllString(string(c.status()), -1) {
		branchName = strings.ReplaceAll(match, "On branch ", "")
	}
    if branchName == "" {
        re := regexp.MustCompile(`Sul branch [\w\/\#\-\.]{0,}`)
        for _, match := range re.FindAllString(string(c.status()), -1) {
            branchName = strings.ReplaceAll(match, "Sul branch ", "")
        }
    }


	fmt.Println(color.RedString("leaving: " + branchName))

	br := branch{branchName}
	fmt.Println(color.RedString("destination: " + br.destination()))

	gitCheckoutToDev := &gitCommand{
		c.Logger,
		[]string{"checkout", br.destination()},
		"Cant checkout destination branch",
		c.conf,
	}

	_ = gitCheckoutToDev.Execute()

	gitMergeNoFF := &gitCommand{
		c.Logger,
		[]string{"merge", "--no-ff", branchName},
		"cant merge",
		c.conf,
	}

	_ = gitMergeNoFF.Execute()

	// @todo set CurrentStep as tagMergedBranchStep{}
	if c.conf.Features.TagAfterMerge == true {
		gitDescribeTags := &gitCommand{
			c.Logger,
			[]string{"describe", "--tags"},
			"cant get tag description",
			c.conf,
		}

		cmdOut := gitDescribeTags.Execute()

		fmt.Print("current tag: ", color.GreenString(string(cmdOut)))

		tagName := ""

		mt := meta{string(cmdOut), branchName}

		// @todo check from configuration if tag must be applied or not
		if br.isHotfix() || br.isPatch() ||  br.isRefactoring() || br.isBugfix() {
			c.Logger.Info("Is Patch branch")
			tagName = mt.NextPatchTag()
		}

		// @todo check from configuration if tag must be applied or not
		if br.isFeature() {
			c.Logger.Info("Is Feature branch")
			tagName = mt.NextMinorTag()
		}

		fmt.Println("next tag:   ", color.GreenString(tagName))

		gitTag := &gitCommand{
			c.Logger,
			[]string{"tag", tagName, "-f"},
			"cant tag",
			c.conf,
		}
		_ = gitTag.Execute()
	} else {
		fmt.Println(color.RedString("tag skipped"))
	}

	// @todo set CurrentStep as delete old branch{}
	gitDeleteOldBranch := &gitCommand{
		c.Logger,
		[]string{"branch", "-D", branchName},
		"cant merge",
		c.conf,
	}
	_ = gitDeleteOldBranch.Execute()
	fmt.Println(color.GreenString("branch " + branchName + " deleted"))

	// @todo set CurrentStep as mergeIntoDevBranchStep{}
	if !br.isDevelopment(branchName) {
		gitCheckoutToDev := &gitCommand{
			c.Logger,
			[]string{"checkout", c.conf.Branches.Historical.Development},
			"Cant checkout destination branch",
			c.conf,
		}
		_ = gitCheckoutToDev.Execute()

		gitMergeNoFastForward := &gitCommand{
			c.Logger,
			[]string{"merge", "--no-ff", br.destination()},
			"cant move to " + br.destination() + " updates",
			c.conf,
		}
		_ = gitMergeNoFastForward.Execute()
	}

	c.CurrentStep = &finalStep{}

	return true
}

func (s completeBranchStep) Stepname() string {
	return "checkout-dev-branch"
}
