package main

import (
	"fmt"

	"github.com/fatih/color"
)

type finalStep struct{}

func (s finalStep) Execute(c *context) bool {

	if c.conf.Features.RemoveRemotelyMerged {

		c.Logger.Info(color.RedString("will remove remotely merged branches"))

		gitCheckoutToDev := &gitCommand{
			c.Logger,
			[]string{"branch", "-a", "--merged"},
			"Cant list all local branches ",
			c.conf,
		}
		output := gitCheckoutToDev.Execute()

		fmt.Println(output)

	} else {
		c.Logger.Info(color.GreenString("leave remotely merged branches"))
	}

	return false
}

func (s finalStep) Stepname() string {
	return "final"
}
