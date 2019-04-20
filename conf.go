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
		TagAfterMerge bool `json:"tagAfterMerge"`
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

func ReadConfiguration(repositoryRoot string) (jj jsonConf, err error) {
	if repositoryRoot == "" {
		return jsonConf{}, fmt.Errorf("invalid repository folder")
	}

	file, errReadingConf := ioutil.ReadFile(
		repositoryRoot + "/.git/ff.conf.json",
	)

	c := jsonConf{}

	if errReadingConf != nil {
		fmt.Println(color.RedString(errReadingConf.Error()))
		os.Exit(1)
	}

	if os.IsNotExist(errReadingConf) {
		c.Branches.Historical.Development = "master"
		c.Features.TagAfterMerge = true
	}

	errUnmarshal := json.Unmarshal([]byte(file), &c)

	if errUnmarshal != nil {
		fmt.Println(color.RedString("config file is corrupted"))
		os.Exit(1)
	}

	return c, nil
}
