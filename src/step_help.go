package main

import (
	"fmt"

	"github.com/fatih/color"
)

type help struct {
	Command     string
	Description string
}

func printHelp(h help) {
	fmt.Println("\t" + color.YellowString(h.Command) + ": " + color.WhiteString(h.Description))
}

type helpStep struct{}

func (s helpStep) Execute(c *context) bool {

	fmt.Println("")

	container := c.container()

	show := make(map[string]bool)

	for _, group := range c.Groups() {
		conta := container[group]
		show[group] = false
		for range conta {
			show[group] = true
		}
	}

	fmt.Print("      " + color.GreenString(c.getRemote()))
	fmt.Println("        " + color.WhiteString(c.getVersion()))

	fmt.Println("      " + color.GreenString("usage"))
	fmt.Println("        " + color.WhiteString("ff [command]"))
	fmt.Println("        " + color.WhiteString("ff config [feature]"))
	fmt.Println("")

	for _, group := range c.Groups() {
		conta := container[group]
		if show[group] {
			fmt.Println("      " + color.GreenString(group))
			for command := range conta {
				printHelp(help{
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

func (s helpStep) Stepname() string {
	return "help"
}
