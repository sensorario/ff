package ff

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/sensorario/slugify"
)

type featureStep struct{}

func (s featureStep) Execute(c *Context) bool {
	developmentBranch := c.Conf.Branches.Historical.Development

	gitCheckoutToDev := &GitCommand{
		c.Logger,
		[]string{"checkout", developmentBranch},
		"Cant checkout " + developmentBranch,
		c.Conf,
	}
	output := gitCheckoutToDev.Execute()
	fmt.Println(output)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(color.RedString("New feature description: "))
	readedString, _ := reader.ReadString('\n')
	featureName := slugify.Slugify(readedString)

	fmt.Print(
		"Feature: ",
	)

	featurePrefix := c.Conf.Branches.Support.Feature
	featureBranchName := featurePrefix + "/" + featureName + "/" + developmentBranch
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

func (s featureStep) Stepname() string {
	return "create-feature-branch"
}
