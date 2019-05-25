package main

import "github.com/fatih/color"

type finalStep struct{}

func (s finalStep) Execute(c *context) bool {

	if c.conf.Features.RemoveRemotelyMerged {
		c.Logger.Info(color.RedString("will remove remotely merged branches"))
	} else {
		c.Logger.Info(color.GreenString("leave remotely merged branches"))
	}

	return false
}

func (s finalStep) Stepname() string {
	return "final"
}
