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
	fmt.Println("\t" + color.YellowString(h.Command) + ": " + color.WhiteString(h.Description))
}

type HelpStep struct{}

func (s HelpStep) Execute(c *Context) bool {

	fmt.Println("")

	container := c.container()

	show := make(map[string]bool)

	for _, group := range c.Groups() {
		conta := container[group]
		show[group] = false
		for _, _ = range conta {
			show[group] = true
		}
	}

	fmt.Println("      " + color.GreenString("usage"))
	fmt.Println("        " + color.WhiteString("ff [command]"))
	fmt.Println("")

	for _, group := range c.Groups() {
		conta := container[group]
		if show[group] {
			fmt.Println("      " + color.GreenString(group))
			for command, _ := range conta {
				printHelp(Help{
					command,
					container[group][command].Description,
				})
			}
		}
	}

	fmt.Println("")

	c.CurrentStep = &finalStep{}

	return true
}

func (s HelpStep) Stepname() string {
	return "help"
}
