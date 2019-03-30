package main

import (
	"fmt"

	"github.com/fatih/color"
)

type Help struct {
	Command     string
	Description string
}

func printHelp(h Help) {
	fmt.Println(color.YellowString(h.Command) + " " + color.GreenString(h.Description))
}

type HelpStep struct{}

func (s *HelpStep) Execute(c *Context) bool {

	printHelp(Help{"help:", "this help"})

	printHelp(Help{"commit:", "commit everything"})
	printHelp(Help{"feature:", "create new feature branch"})
	printHelp(Help{"hotfix:", "ceate new hotfix branch"})
	printHelp(Help{"complete:", "merge hotfix or feature branch"})

	printHelp(Help{"publish:", "push current branch into remote"})
	printHelp(Help{"reset:", "reset working directory and stage"})
	printHelp(Help{"status:", "check status of current branch"})

	c.CurrentStep = &FinalStep{}

	return true
}

func (s *HelpStep) Stepname() string {
	return "help"
}
