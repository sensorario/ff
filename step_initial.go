package main

import (
	"os"
)

type InputReadingStep struct{}

func (s InputReadingStep) Execute(c *Context) bool {
	command := "help"
	specification := "default"

	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	if len(os.Args) > 2 {
		specification = os.Args[2]
	}

	if command == "help" {
		c.CurrentStep = &HelpStep{}
		return true
	}

	if command == "publish" {
		c.CurrentStep = &PublishStep{}
		return true
	}

	if command == "reset" {
		c.CurrentStep = &ResetStep{}
		return true
	}

	if command == "status" {
		c.CurrentStep = &StatusStep{}
		return true
	}

	if command == "commit" {
		c.CurrentStep = &WorkingDirStep{}
		return true
	}

	if command == "feature" {
		if specification == "default" {
			c.CurrentStep = &FeatureStep{}
		}
		if specification == "complete" {
			c.CurrentStep = &CompleteBranchStep{}
		}
		return true
	}

	if command == "complete" {
		c.CurrentStep = &CompleteBranchStep{}
		return true
	}

	if command == "hotfix" {
		if specification == "default" {
			c.CurrentStep = &HotfixStep{}
		}
		if specification == "complete" {
			c.CurrentStep = &CompleteBranchStep{}
		}
		return true
	}

	c.CurrentStep = &FinalStep{}

	return true
}

func (s InputReadingStep) Stepname() string {
	return "command-detection"
}
