package main

import (
	"fmt"
	"os"
)

type InputReadingStep struct{}

func (s *InputReadingStep) Execute(c *Context) bool {
	command := "status"
	specification := "default"

	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	if len(os.Args) > 2 {
		specification = os.Args[2]
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
		if specification == "default" {
			c.CurrentStep = &FeatureStep{}
		}
		if specification == "complete" {
			c.CurrentStep = &CompleteFeatureStep{}
		}
		return true
	}

	if command == "hotfix" {
		c.CurrentStep = &HotfixStep{}
		return true
	}

	c.CurrentStep = &FinalStep{}

	return true
}

func (s *InputReadingStep) Stepname() string {
	return "command-detection"
}
