package ff

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/sensorario/branch"
)

type completeBranchStep struct{}

func (s completeBranchStep) Execute(c *Context) bool {
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

	br := branch.Branch{branchName}
	fmt.Println(color.RedString("destination: " + br.Destination()))

	gitCheckoutToDev := &GitCommand{
		c.Logger,
		[]string{"checkout", br.Destination()},
		"Cant checkout destination branch",
		c.Conf,
	}

	_ = gitCheckoutToDev.Execute()

	gitMergeNoFF := &GitCommand{
		c.Logger,
		[]string{"merge", "--no-ff", branchName},
		"cant merge",
		c.Conf,
	}

	_ = gitMergeNoFF.Execute()

	if c.Conf.Features.TagAfterMerge == true {
		gitDescribeTags := &GitCommand{
			c.Logger,
			[]string{"describe", "--tags"},
			"cant get tag description",
			c.Conf,
		}

		cmdOut := gitDescribeTags.Execute()

		fmt.Print("current tag: ", color.GreenString(string(cmdOut)))

		tagName := ""

		mt := meta{string(cmdOut), branchName}

		if br.IsHotfix() || br.IsPatch() || br.IsRefactoring() || br.IsBugfix() {
			c.Logger.Info("Is Patch branch")
			tagName = mt.NextPatchTag()
		}

		if br.IsFeature() {
			c.Logger.Info("Is Feature branch")
			tagName = mt.NextMinorTag()
		}

		fmt.Println("next tag:   ", color.GreenString(tagName))

		gitTag := &GitCommand{
			c.Logger,
			[]string{"tag", tagName, "-f"},
			"cant tag",
			c.Conf,
		}
		_ = gitTag.Execute()
	} else {
		fmt.Println(color.RedString("tag skipped"))
	}

	gitDeleteOldBranch := &GitCommand{
		c.Logger,
		[]string{"branch", "-D", branchName},
		"cant merge",
		c.Conf,
	}
	_ = gitDeleteOldBranch.Execute()
	fmt.Println(color.GreenString("branch " + branchName + " deleted"))

	if !br.IsDevelopment(branchName) {
		gitCheckoutToDev := &GitCommand{
			c.Logger,
			[]string{"checkout", c.Conf.Branches.Historical.Development},
			"Cant checkout destination branch",
			c.Conf,
		}
		_ = gitCheckoutToDev.Execute()

		gitMergeNoFastForward := &GitCommand{
			c.Logger,
			[]string{"merge", "--no-ff", br.Destination()},
			"cant move to " + br.Destination() + " updates",
			c.Conf,
		}
		_ = gitMergeNoFastForward.Execute()
	}

	c.CurrentStep = &finalStep{}

	return true
}

func (s completeBranchStep) Stepname() string {
	return "checkout-dev-branch"
}
