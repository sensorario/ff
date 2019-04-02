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

	// always
	container["help"] = &HelpStep{}
	container["status"] = &StatusStep{}

	// only if working dir is dirty
	container["commit"] = &WorkingDirStep{}
	container["reset"] = &ResetStep{}

	branch := c.CurrentBranch()
	sem := Branch{branch}

	if sem.IsMaster() == true {
		container["publish"] = &PublishStep{}
		container["hotfix"] = &HotfixStep{}
		container["feature"] = &FeatureStep{}
		container["refactor"] = &RefactoringStep{}
	}

	if sem.IsFeature() || sem.IsHotfix() {
		container["complete"] = &CompleteBranchStep{}
	}

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
