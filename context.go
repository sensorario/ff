package main

import (
	"regexp"
	"strings"

	"github.com/sensorario/gol"
)

type Context struct {
	CurrentStep FussyStepInterface
	Exit        bool
	Logger      gol.Logger
}

func (c Context) CurrentBranch() string {
	gitStatus := &GitCommand{
		c.Logger,
		[]string{"status"},
		"Cant get status",
	}

	cmdOut := gitStatus.Execute()

	re := regexp.MustCompile(`On branch [\w\/\#\-]{0,}`)

	branchName := ""

	for _, match := range re.FindAllString(string(cmdOut), -1) {
		branchName = strings.ReplaceAll(match, "On branch ", "")
	}

	return branchName
}

type Step struct {
	Step        FussyStepInterface
	Description string
}

func (c Context) Container() map[string]Step {
	ss := map[string]Step{}

	ss["help"] = Step{HelpStep{}, "this help"}
	ss["status"] = Step{&StatusStep{}, "status"}

	// only if working dir is dirty
	ss["commit"] = Step{WorkingDirStep{}, "commit everything"}
	ss["reset"] = Step{ResetStep{}, "reset working directory and stage"}

	branch := c.CurrentBranch()
	sem := Branch{branch}

	if sem.IsMaster() {
		ss["publish"] = Step{PublishStep{}, "push current branch into remote"}
		ss["hotfix"] = Step{HotfixStep{}, "create new hotfix branch"}
		ss["feature"] = Step{FeatureStep{}, "create new feature branch"}
		ss["refactor"] = Step{RefactoringStep{}, "create new refactor branch"}
	}

	if sem.IsFeature() || sem.IsHotfix() {
		ss["complete"] = Step{CompleteBranchStep{}, "merge current branch into master"}
	}

	return ss
}
