package ff

import (
	"fmt"

	"github.com/fatih/color"
)

type statusStep struct{}

func (s statusStep) Execute(c *Context) bool {
	branchName := c.currentBranch()

	fmt.Println(
		"Current branch is ",
		color.GreenString(branchName),
	)

	return false
}

func (s statusStep) Stepname() string {
	return "status"
}
