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
	fmt.Println(color.YellowString(h.Command) + ": " + color.GreenString(h.Description))
}

type HelpStep struct{}

func (s HelpStep) Execute(c *Context) bool {

	fmt.Println("")

	container := c.Container()

	for command, _ := range container {
		printHelp(Help{
			command,
			container[command].Description,
		})
	}

	fmt.Println("")

	c.CurrentStep = &FinalStep{}

	return true
}

func (s HelpStep) Stepname() string {
	return "help"
}
