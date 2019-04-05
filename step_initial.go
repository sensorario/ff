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

	for _, group := range c.Groups() {
		item, ok := ss[group][command]
		if ok {
			c.CurrentStep = item.Step
			return true
		}
	}

	fmt.Println(color.RedString(
		command + " is not available here",
	))

	return false
}

func (s InputReadingStep) Stepname() string {
	return "command-detection"
}
