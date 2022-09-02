package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
)

type jsonConf struct {
	Features struct {
		ApplyFirstTag        bool   `json:"applyFirstTag"`
		DisableUndoCommand   bool   `json:"disableUndoCommand"`
		EnableGitCommandLog  bool   `json:"enableGitCommandLog"`
		ForceOnPublish       bool   `json:"forceOnPublish"`
		PushTagsOnPublish    bool   `json:"pushTagsOnPublish"`
		RemoveRemotelyMerged bool   `json:"removeRemotelyMerged"`
		StopAskingForTags    bool   `json:"stopAskingForTags"`
		TagAfterMerge        bool   `json:"tagAfterMerge"`
		Lang                 string `json:"lang"`
	} `json:"features"`
	Branches struct {
		Historical struct {
			Development string `json:"development"`
			//Poduction   string `json:"production"`
		} `json:"historical"`
		//Support struct {
		//Feature string `json:"feature"`
		//Release string `json:"release"`
		//Hotfix  string `json:"hotfix"`
		//Bugfix  string `json:"bugfix"`
		//} `json:"support"`
	} `json:"branches"`
}

func readConfiguration(repositoryRoot string) (jj jsonConf, err error) {
	c := jsonConf{}

	c.Branches.Historical.Development = "master"
	c.Features.DisableUndoCommand = false
	c.Features.EnableGitCommandLog = false
	c.Features.RemoveRemotelyMerged = false
	c.Features.StopAskingForTags = true
	c.Features.TagAfterMerge = false
    c.Features.Lang = "it"

	if repositoryRoot == "" {
		return c, fmt.Errorf("invalid repository folder")
	}

	file, errReadingConf := ioutil.ReadFile(
		repositoryRoot + "/.git/ff.conf.json",
	)

	if os.IsNotExist(errReadingConf) {
		return c, nil
	}

	errUnmarshal := json.Unmarshal([]byte(file), &c)
	if errUnmarshal != nil {
		fmt.Println(color.RedString("config file is corrupted"))
		os.Exit(1)
	}

	confIndented, _ := json.MarshalIndent(c, "", "  ")

	ioutil.WriteFile(".git/ff.conf.json", confIndented, 0644)

	if errReadingConf != nil {
		fmt.Println(color.RedString(errReadingConf.Error()))
		os.Exit(1)
	}

	if c.Branches.Historical.Development == "" {
		fmt.Println(color.RedString("Oops! Bad configuration: development branch cannot be empty"))
		os.Exit(1)
	}

	return c, nil
}
