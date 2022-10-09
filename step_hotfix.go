package ff

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/sensorario/slugify"
)

type hotfixStep struct{}

func (s hotfixStep) Execute(c *Context) bool {
	developmentBranch := c.currentBranch()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(color.RedString("Hotfix description: "))
	readedString, _ := reader.ReadString('\n')
	hotfixDescription := slugify.Slugify(readedString)

	hotfixBranch := "hotfix/" + hotfixDescription + "/" + developmentBranch
	fmt.Println("Hotfix: ", color.YellowString(hotfixBranch))

	gitCheckoutNewBranch := &GitCommand{
		c.Logger,
		[]string{"checkout", "-b", hotfixBranch},
		"Cant create new branch",
		c.Conf,
	}

	_ = gitCheckoutNewBranch.Execute()

	c.CurrentStep = &finalStep{}

	return true
}

func (s hotfixStep) Stepname() string {
	return "create-hotfix-branch"
}
