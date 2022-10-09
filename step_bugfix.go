package ff

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/sensorario/slugify"
)

type bugfixStep struct{}

func (s bugfixStep) Execute(c *Context) bool {
	developmentBranch := c.Conf.Branches.Historical.Development

	gitCheckoutBackToDev := &GitCommand{
		c.Logger,
		[]string{"checkout", developmentBranch},
		"Cant checkout " + developmentBranch,
		c.Conf,
	}
	_ = gitCheckoutBackToDev.Execute()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(color.RedString("Bugfix description: "))
	readedString, _ := reader.ReadString('\n')
	bugfixDescription := slugify.Slugify(readedString)

	bugfixBranch := "bugfix/" + bugfixDescription + "/" + developmentBranch
	fmt.Println("Bugfix: ", color.YellowString(bugfixBranch))

	gitCheckoutNewBranch := &GitCommand{
		c.Logger,
		[]string{"checkout", "-b", bugfixBranch},
		"Cant create new branch",
		c.Conf,
	}

	_ = gitCheckoutNewBranch.Execute()

	c.CurrentStep = &finalStep{}

	return true
}

func (s bugfixStep) Stepname() string {
	return "create-bugfix-branch"
}
