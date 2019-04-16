package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type JSONConf struct {
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

func ReadConfiguration() JSONConf {
	fmt.Println("Read json configuration")
	dir, _ := os.Getwd()
	file, _ := ioutil.ReadFile(dir + ".git/conf.json")

	data := JSONConf{}
	data.Branches.Historical.Development = "master"

	json.Unmarshal([]byte(file), &data)

	return data
}
