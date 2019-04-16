package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
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
	dir, _ := os.Getwd()

	file, err := ioutil.ReadFile(dir + "/.git/conf.json")

	c := jsonConf{}

	// se non c'e' il file, ... allora imposto
	// dei valori di default
	if err != nil {
		c.Branches.Historical.Development = "master"
	}

	// defaults

	json.Unmarshal([]byte(file), &c)
	return c
}
