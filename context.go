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

type Complesso struct {
	Step        FussyStepInterface
	Description string
}

func (c Context) Container() map[string]Complesso {
	container := map[string]Complesso{}

	// always
	container["help"] = Complesso{
		HelpStep{},
		"this help",
	}

	container["status"] = Complesso{
		&StatusStep{},
		"status",
	}

	// only if working dir is dirty
	container["commit"] = Complesso{
		WorkingDirStep{},
		"commit everything",
	}

	container["reset"] = Complesso{
		ResetStep{},
		"reset working directory and stage",
	}

	branch := c.CurrentBranch()
	sem := Branch{branch}

	if sem.IsMaster() == true {
		container["publish"] = Complesso{
			PublishStep{},
			"push current branch into remote",
		}

		container["hotfix"] = Complesso{
			HotfixStep{},
			"create new hotfix branch",
		}

		container["feature"] = Complesso{
			FeatureStep{},
			"create new feature branch",
		}

		container["refactor"] = Complesso{
			RefactoringStep{},
			"create new refactor branch",
		}
	}

	if sem.IsFeature() || sem.IsHotfix() {
		container["complete"] = Complesso{
			CompleteBranchStep{},
			"merge current branch into master",
		}
	}

	return container
}
