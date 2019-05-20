package main

import "regexp"
import "strings"
import "github.com/sensorario/gol"

type context struct {
	RepositoryRoot string
	CurrentStep    fFStep
	Exit           bool
	Logger         gol.Logger
	devBranchName  string
	conf           jsonConf
	st             string
}

func (c context) status() string {
	if c.st == "" {
		gitStatus := &gitCommand{
			c.Logger,
			[]string{"status"},
			"Cant get status",
			c.conf,
		}

		c.st = gitStatus.Execute()
	}

	return c.st
}

func (c context) currentBranch() string {
	re := regexp.MustCompile(`On branch [\w\/\#\-\.]{0,}`)

	branchName := ""

	for _, match := range re.FindAllString(string(c.status()), -1) {
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

	ss["exec"] = make(map[string]stepType)
	ss["start"] = make(map[string]stepType)
	ss["working"] = make(map[string]stepType)

	ss["exec"]["help"] = stepType{helpStep{}, "this help"}
	ss["exec"]["status"] = stepType{&statusStep{}, "status"}
	ss["exec"]["publish"] = stepType{publishStep{}, "push current branch into remote"}
	ss["exec"]["pull"] = stepType{pullStep{}, "pull current branch from remote"}
	ss["exec"]["authors"] = stepType{authorsStep{}, "list all committers"}
	ss["exec"]["fetch_all"] = stepType{fetchAllStep{}, "fetch all branches"}
	ss["exec"]["conf"] = stepType{confStep{}, "show configuration"}
	ss["exec"]["config"] = stepType{configStep{}, "update configuration"}

	if !c.isWorkingDirClean() {
		ss["working"]["commit"] = stepType{wokingDirStep{}, "commit everything"}
		ss["working"]["reset"] = stepType{resetStep{}, "reset working directory and stage"}
	} else {
		if c.conf.Features.DisableUndoCommand == false {
			ss["working"]["undo"] = stepType{undoStep{}, "undo last commit"}
		}
	}

	name := c.currentBranch()
	sem := branch{name}

	if sem.isDevelopment(c.devBranchName) {
		ss["start"]["bugfix"] = stepType{bugfixStep{}, "create new bugfix branch"}
		ss["start"]["feature"] = stepType{featureStep{}, "create new feature branch"}
		ss["start"]["refactor"] = stepType{refactoringStep{}, "create new refactor branch"}
		ss["start"]["patch"] = stepType{patchStep{}, "create new patch branch"}
		ss["exec"]["tag"] = stepType{tagStep{}, "force creation of new tag"}
	} else {
		ss["start"]["hotfix"] = stepType{hotfixStep{}, "create new hotfix branch"}
	}

	if sem.isRefactoring() || sem.isFeature() || sem.isHotfix() || sem.isBugfix() {
		if c.isWorkingDirClean() {
			ss["start"]["complete"] = stepType{
				completeBranchStep{},
				"merge current branch into " + c.conf.Branches.Historical.Development,
			}
		}
	}

	return ss
}

func (c context) isWorkingDirClean() bool {
	re := regexp.MustCompile(`(?m)nothing to commit, working tree clean`)

	for range re.FindAllString(string(c.status()), -1) {
		c.Logger.Info("working dir clean")
		return true
	}

	c.Logger.Info("working dir dirty")
	return false
}

func (c context) Groups() []string {
	return []string{
		"exec",
		"start",
		"working",
	}
}
