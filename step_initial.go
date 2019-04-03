package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type InputReadingStep struct{}

func (s InputReadingStep) Execute(c *Context) bool {
	command := "help"

	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	ss := c.Container()

	item, ok := ss[command]

	if !ok {
		fmt.Println(color.RedString(
			command + " is not available here",
		))
		os.Exit(1)
	}

	c.CurrentStep = item.Step

	return true
}

func (s InputReadingStep) Stepname() string {
	return "command-detection"
}
