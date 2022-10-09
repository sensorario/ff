package ff

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/sensorario/slugify"
)

type patchStep struct{}

func (s patchStep) Execute(c *Context) bool {
	developmentBranch := c.Conf.Branches.Historical.Development

	gitCheckoutToDev := &GitCommand{
		c.Logger,
		[]string{"checkout", developmentBranch},
		"Cant checkout " + developmentBranch,
		c.Conf,
	}

	_ = gitCheckoutToDev.Execute()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(color.RedString("Patch descrption: "))
	readedString, _ := reader.ReadString('\n')
	patchDescription := slugify.Slugify(readedString)

	fmt.Print(
		"Patch: ",
	)

	patchBranchName := "patch/" + patchDescription + "/" + developmentBranch
	fmt.Println(color.YellowString(patchBranchName))

	gitCheckoutNewBranch := &GitCommand{
		c.Logger,
		[]string{"checkout", "-b", patchBranchName},
		"Cant create new branch",
		c.Conf,
	}

	_ = gitCheckoutNewBranch.Execute()

	c.CurrentStep = &finalStep{}

	return true
}

func (s patchStep) Stepname() string {
	return "patch-branch"
}
