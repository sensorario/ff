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
		TagAfterMerge       bool `json:"tagAfterMerge"`
		DisableUndoCommand  bool `json:"disableUndoCommand"`
		StopAskingForTags   bool `json:"stopAskingForTags"`
		ApplyFirstTag       bool `json:"applyFirstTag"`
		EnableGitCommandLog bool `json:"enableGitCommandLog"`
		ForceOnPublish      bool `json:"forceOnPublish"`
		PushTagsOnPublish   bool `json:"pushTagsOnPublish"`
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

func defaultConfig() jsonConf {
	c := jsonConf{}

	c.Features.TagAfterMerge = true
	c.Features.DisableUndoCommand = false
	c.Features.StopAskingForTags = false
	c.Features.ApplyFirstTag = true
	c.Features.EnableGitCommandLog = false
	c.Features.ForceOnPublish = false
	c.Features.PushTagsOnPublish = false

	c.Branches.Historical.Development = "master"

	return c
}

func readConfiguration(repositoryRoot string) (jj jsonConf, err error) {
	c := defaultConfig()

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
