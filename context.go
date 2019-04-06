package main

import (
	"regexp"
	"strings"

	"github.com/sensorario/gol"
)

type Context struct {
	CurrentStep FFStep
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
	Step        FFStep
	Description string
}

func (c Context) Container() map[string]map[string]Step {
	ss := map[string]map[string]Step{}

	ss["command"] = make(map[string]Step)
	ss["features"] = make(map[string]Step)
	ss["working"] = make(map[string]Step)

	ss["command"]["help"] = Step{HelpStep{}, "this help"}
	ss["command"]["status"] = Step{&StatusStep{}, "status"}

	if !c.IsWorkingDirClean() {
		ss["working"]["commit"] = Step{WorkingDirStep{}, "commit everything"}
		ss["working"]["reset"] = Step{ResetStep{}, "reset working directory and stage"}
	}

	branch := c.CurrentBranch()
	sem := Branch{branch}

	if sem.IsMaster() {
		ss["command"]["publish"] = Step{PublishStep{}, "push current branch into remote"}
		ss["features"]["bugfix"] = Step{BugfixStep{}, "create new bugfix branch"}
		ss["features"]["feature"] = Step{FeatureStep{}, "create new feature branch"}
		ss["features"]["refactor"] = Step{RefactoringStep{}, "create new refactor branch"}
	}

	if sem.IsRefactoring() || sem.IsFeature() || sem.IsHotfix() || sem.IsBugfix() {
		if c.IsWorkingDirClean() {
			ss["features"]["complete"] = Step{CompleteBranchStep{}, "merge current branch into master"}
		}
	}

	if sem.Phase() == "production" {
		ss["features"]["hotfix"] = Step{
			HotfixStep{},
			"create new hotfix branch",
		}
	}

	return ss
}

func (c Context) IsWorkingDirClean() bool {
	gitStatus := &GitCommand{
		c.Logger,
		[]string{"status"},
		"Cant get status",
	}

	cmdOut := gitStatus.Execute()

	re := regexp.MustCompile(`(?m)nothing to commit, working tree clean`)

	for _, _ = range re.FindAllString(string(cmdOut), -1) {
		c.Logger.Info("working dir clean")
		return true
	}

	c.Logger.Info("working dir dirty")
	return false
}

func (c Context) Groups() []string {
	return []string{
		"command",
		"features",
		"working",
	}
}
