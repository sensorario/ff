package main

import (
	"strings"

	"github.com/fatih/color"
)

type finalStep struct{}

func (s finalStep) Execute(c *context) bool {

	if c.conf.Features.RemoveRemotelyMerged && c.currentBranch() == c.conf.Branches.Historical.Development {

		c.Logger.Info(color.RedString("will remove remotely merged branches"))

		gitCheckoutToDev := &gitCommand{
			c.Logger,
			[]string{"branch", "-a", "--merged"},
			"Cant list all local branches ",
			c.conf,
		}
		output := gitCheckoutToDev.Execute()

		lines := strings.Split(output, "\n")
		for _, line := range lines {
			if strings.Contains(line, "remotes/origin") {

				branchName := strings.Replace(strings.Trim(line, " "), "remotes/origin/", "", 1)
				if branchName != c.conf.Branches.Historical.Development && !strings.Contains(branchName, "HEAD") {
					deleteRemoteBranch := &gitCommand{
						c.Logger,
						[]string{"push", "origin", ":" + strings.Replace(strings.Trim(line, " "), "remotes/origin/", "", 1)},
						"Cant list all local branches ",
						c.conf,
					}

					deleteRemoteBranch.Execute()
				}
			}
		}

	} else {
		c.Logger.Info(color.GreenString("leave remotely merged branches"))
	}

	return false
}

func (s finalStep) Stepname() string {
	return "final"
}
