package main

import (
	"regexp"
	"strings"

	"github.com/sensorario/gol"
)

type context struct {
	CurrentStep fFStep
	Exit        bool
	Logger      gol.Logger
}

func (c context) currentBranch() string {
	gitStatus := &gitCommand{
		c.Logger,
		[]string{"status"},
		"Cant get status",
	}

	cmdOut := gitStatus.Execute()

	re := regexp.MustCompile(`On branch [\w\/\#\-\.]{0,}`)

	branchName := ""

	for _, match := range re.FindAllString(string(cmdOut), -1) {
		branchName = strings.ReplaceAll(match, "On branch ", "")
	}

	return branchName
}

type Step struct {
	Step        fFStep
	Description string
}

func (c context) container() map[string]map[string]Step {
	ss := map[string]map[string]Step{}

	ss["command"] = make(map[string]Step)
	ss["features"] = make(map[string]Step)
	ss["working"] = make(map[string]Step)

	ss["command"]["help"] = Step{HelpStep{}, "this help"}
	ss["command"]["status"] = Step{&StatusStep{}, "status"}
	ss["command"]["publish"] = Step{PublishStep{}, "push current branch into remote"}

	if !c.isWorkingDirClean() {
		ss["working"]["commit"] = Step{WorkingDirStep{}, "commit everything"}
		ss["working"]["reset"] = Step{ResetStep{}, "reset working directory and stage"}
	}

	name := c.currentBranch()
	sem := branch{name}

	if sem.isMaster() {
		ss["features"]["bugfix"] = Step{bugfixStep{}, "create new bugfix branch"}
		ss["features"]["feature"] = Step{featureStep{}, "create new feature branch"}
		ss["features"]["refactor"] = Step{RefactoringStep{}, "create new refactor branch"}
	}

	if sem.isRefactoring() || sem.isFeature() || sem.isHotfix() || sem.isBugfix() {
		if c.isWorkingDirClean() {
			ss["features"]["complete"] = Step{completeBranchStep{}, "merge current branch into master"}
		}
	}

	if sem.phase() == "production" {
		ss["features"]["hotfix"] = Step{hotfixStep{}, "create new hotfix branch"}
	}

	return ss
}

func (c context) isWorkingDirClean() bool {
	gitStatus := &gitCommand{
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

func (c context) Groups() []string {
	return []string{
		"command",
		"features",
		"working",
	}
}
