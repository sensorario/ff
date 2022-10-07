package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type inputReadingStep struct{}

func (s inputReadingStep) Execute(c *context) bool {
	command := "help"

	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	ss := c.container()

	c.args(os.Args)

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

func (s inputReadingStep) Stepname() string {
	return "command-detection"
}
