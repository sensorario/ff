package ff

import (
	"regexp"
	"strings"
	//	"fmt"

	"github.com/sensorario/branch"
	"github.com/sensorario/gol/v2"
	"github.com/sensorario/tongue"
)

type Context struct {
	RepositoryRoot string
	CurrentStep    fFStep
	Exit           bool
	Logger         gol.Logger
	DevBranchName  string
	Conf           JsonConf
	st             string
	arguments      []string
	version        string
	remote         string
}

func (c Context) getInput() []string {
	return c.arguments
}

func (c *Context) setCurrentVersion(version string) {
	c.version = version
}

func (c Context) getRemote() string {
	return c.remote
}

func (c *Context) setRemote(remote string) {
	c.remote = remote
}

func (c Context) getVersion() string {
	return c.version
}

func (c *Context) args(input []string) {
	c.arguments = input
}

func (c Context) status() string {
	if c.st == "" {
		gitStatus := &GitCommand{
			c.Logger,
			[]string{"status"},
			"Cant get status",
			c.Conf,
		}

		c.st = gitStatus.Execute()
	}

	return c.st
}

func (c Context) currentBranch() string {
	branchName := ""
	re := regexp.MustCompile(`On branch [\w\/\#\-\.]{0,}`)
	for _, match := range re.FindAllString(string(c.status()), -1) {
		branchName = strings.ReplaceAll(match, "On branch ", "")
	}

	if branchName != "" {
		return branchName
	}

	re = regexp.MustCompile(`Sul branch [\w\/\#\-\.]{0,}`)
	for _, match := range re.FindAllString(string(c.status()), -1) {
		branchName = strings.ReplaceAll(match, "Sul branch ", "")
	}

	return branchName
}

type stepType struct {
	Step        fFStep
	Description string
}

type translation struct {
	t string
}

func CreateDictionary() tongue.Dict {
	dict := tongue.LoadDictionary()

	dict.Add("it", "this.help", "questo help")
	dict.Add("it", "status", "status")
	dict.Add("it", "publish", "metti questo branch nel remote")
	dict.Add("it", "pull", "tira giu questo branch dal remote")
	dict.Add("it", "list.committers", "elenca tutti quelli che hanno committatto")

	dict.Add("en", "this.help", "this help")
	dict.Add("en", "status", "status")
	dict.Add("en", "publish", "publish current branch on its remote")
	dict.Add("en", "pull", "get updates from remote")
	dict.Add("en", "list.committers", "list all committers")

	return dict
}

func (c Context) container() map[string]map[string]stepType {
	dict := CreateDictionary()

	// fmt.Println("language")
	//fmt.Println(c.Conf.Features.Lang)

	ss := map[string]map[string]stepType{}

	ss["exec"] = make(map[string]stepType)
	ss["start"] = make(map[string]stepType)
	ss["working"] = make(map[string]stepType)

	ss["exec"]["help"] = stepType{helpStep{}, dict.Get(c.Conf.Features.Lang, "this.help")}
	ss["exec"]["status"] = stepType{&statusStep{}, dict.Get(c.Conf.Features.Lang, "status")}
	ss["exec"]["publish"] = stepType{publishStep{}, dict.Get(c.Conf.Features.Lang, "publish")}
	ss["exec"]["push"] = stepType{publishStep{}, dict.Get(c.Conf.Features.Lang, "publish")}
	ss["exec"]["pull"] = stepType{pullStep{}, dict.Get(c.Conf.Features.Lang, "pull")}
	ss["exec"]["authors"] = stepType{authorsStep{}, dict.Get(c.Conf.Features.Lang, "list.committers")}
	ss["exec"]["fetch_all"] = stepType{fetchAllStep{}, "fetch all branches"}
	ss["exec"]["conf"] = stepType{confStep{}, "show configuration"}
	ss["exec"]["config"] = stepType{configStep{}, "update configuration"}
	ss["exec"]["lang <language>"] = stepType{configStep{}, "change language"}

	if !c.isWorkingDirClean() {
		ss["working"]["commit"] = stepType{wokingDirStep{}, "commit everything"}
		ss["working"]["reset"] = stepType{resetStep{}, "reset working directory and stage"}
	} else {
		if c.Conf.Features.DisableUndoCommand == false {
			ss["working"]["undo"] = stepType{undoStep{}, "undo last commit"}
		}
	}

	name := c.currentBranch()
	sem := branch.Branch{name}

	if sem.IsDevelopment(c.DevBranchName) {
		ss["start"]["bugfix"] = stepType{bugfixStep{}, "create new bugfix branch"}
		ss["start"]["feature"] = stepType{featureStep{}, "create new feature branch"}
		ss["start"]["refactor"] = stepType{refactoringStep{}, "create new refactor branch"}
		ss["start"]["patch"] = stepType{patchStep{}, "create new patch branch"}
		ss["exec"]["tag"] = stepType{tagStep{}, "force creation of new tag"}
	} else {
		ss["start"]["hotfix"] = stepType{hotfixStep{}, "create new hotfix branch"}
	}

	if sem.IsRefactoring() || sem.IsPatch() || sem.IsFeature() || sem.IsHotfix() || sem.IsBugfix() {
		if c.isWorkingDirClean() {
			ss["start"]["complete"] = stepType{
				completeBranchStep{},
				"merge current branch into " + c.Conf.Branches.Historical.Development,
			}
		} else {
			c.Logger.Info(`Working directory is not clean`)
		}
	}

	return ss
}

func (c Context) isWorkingDirClean() bool {
	re := regexp.MustCompile(`(?m)nothing to commit, working tree clean`)
	for range re.FindAllString(string(c.status()), -1) {
		c.Logger.Info("working dir clean")
		return true
	}

	re = regexp.MustCompile(`non c'è nulla di cui eseguire il commit`)
	for range re.FindAllString(string(c.status()), -1) {
		c.Logger.Info("working dir clean")
		return true
	}

	c.Logger.Info("working dir dirty")
	return false
}

func (c Context) Groups() []string {
	return []string{
		"exec",
		"start",
		"working",
	}
}
