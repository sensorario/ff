package main

import "regexp"
import "strings"
import "github.com/sensorario/gol"

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

type stepType struct {
	Step        fFStep
	Description string
}

func (c context) container() map[string]map[string]stepType {
	ss := map[string]map[string]stepType{}

	ss["command"] = make(map[string]stepType)
	ss["features"] = make(map[string]stepType)
	ss["working"] = make(map[string]stepType)

	ss["command"]["help"] = stepType{helpStep{}, "this help"}
	ss["command"]["status"] = stepType{&statusStep{}, "status"}
	ss["command"]["publish"] = stepType{publishStep{}, "push current branch into remote"}

	if !c.isWorkingDirClean() {
		ss["working"]["commit"] = stepType{wokingDirStep{}, "commit everything"}
		ss["working"]["reset"] = stepType{resetStep{}, "reset working directory and stage"}
	}

	name := c.currentBranch()
	sem := branch{name}

	if sem.isMaster() {
		ss["features"]["bugfix"] = stepType{bugfixStep{}, "create new bugfix branch"}
		ss["features"]["feature"] = stepType{featureStep{}, "create new feature branch"}
		ss["features"]["refactor"] = stepType{refactoringStep{}, "create new refactor branch"}
		ss["command"]["tag"] = stepType{tagStep{}, "force creation of new tag"}
	}

	if sem.isRefactoring() || sem.isFeature() || sem.isHotfix() || sem.isBugfix() {
		if c.isWorkingDirClean() {
			ss["features"]["complete"] = stepType{completeBranchStep{}, "merge current branch into master"}
		}
	}

	if sem.phase() == "production" {
		ss["features"]["hotfix"] = stepType{hotfixStep{}, "create new hotfix branch"}
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

	for _ = range re.FindAllString(string(cmdOut), -1) {
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
