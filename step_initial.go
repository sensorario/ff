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

	container := map[string]FussyStepInterface{}

	container["help"] = &HelpStep{}
	container["publish"] = &PublishStep{}
	container["reset"] = &ResetStep{}
	container["status"] = &StatusStep{}
	container["commit"] = &WorkingDirStep{}
	container["complete"] = &CompleteBranchStep{}
	container["feature"] = &FeatureStep{}
	container["hotfix"] = &HotfixStep{}
	container["refactor"] = &RefactoringStep{}

	step, ok := container[command]

	if !ok {
		fmt.Println(color.RedString(command + " is not in the map"))
		os.Exit(1)
	}

	c.CurrentStep = step

	return true
}

func (s InputReadingStep) Stepname() string {
	return "command-detection"
}
