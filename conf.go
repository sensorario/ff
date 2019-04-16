package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
)

type jsonConf struct {
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

func ReadConfiguration() jsonConf {
	dir, err := os.Getwd()

	if err != nil {
		fmt.Println(color.RedString(err.Error()))
		os.Exit(1)
	}

	file, err := ioutil.ReadFile(dir + "/.git/ff.conf.json")

	c := jsonConf{}

	if err != nil {
		c.Branches.Historical.Development = "master"
	}

	json.Unmarshal([]byte(file), &c)

	return c
}
