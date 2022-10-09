package ff

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/sensorario/slugify"
)

type refactoringStep struct{}

func (s refactoringStep) Execute(c *Context) bool {
	developmentBranch := c.Conf.Branches.Historical.Development

	gitCheckoutToDev := &GitCommand{
		c.Logger,
		[]string{"checkout", developmentBranch},
		"Cant checkout " + developmentBranch,
		c.Conf,
	}

	_ = gitCheckoutToDev.Execute()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(color.RedString("Refactoring descrption: "))
	readedString, _ := reader.ReadString('\n')
	featureName := slugify.Slugify(readedString)

	fmt.Print(
		"Feature: ",
	)

	featureBranchName := "refactor/" + featureName + "/" + developmentBranch
	fmt.Println(color.YellowString(featureBranchName))

	gitCheckoutNewBranch := &GitCommand{
		c.Logger,
		[]string{"checkout", "-b", featureBranchName},
		"Cant create new branch",
		c.Conf,
	}

	_ = gitCheckoutNewBranch.Execute()

	c.CurrentStep = &finalStep{}

	return true
}

func (s refactoringStep) Stepname() string {
	return "create-feature-branch"
}
