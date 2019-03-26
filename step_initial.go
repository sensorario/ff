package main

import (
	"fmt"
	"os"
)

type InputReadingStep struct{}

func (s *InputReadingStep) Execute(c *Context) bool {
	c.EnterStep()

	command := "status"

	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	fmt.Println("command: " + command)

	if command == "status" {
		c.CurrentStep = &StatusStep{}
		return true
	}

	if command == "commit" {
		c.CurrentStep = &CommitStep{}
		return true
	}

	if command == "feature" {
		c.CurrentStep = &FeatureStep{}
		return true
	}

	c.CurrentStep = &FinalStep{}

	return true
}

func (s *InputReadingStep) Stepname() string {
	return "command-detection"
}
