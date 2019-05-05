package main

import (
	"fmt"
	"io/ioutil"

	"github.com/fatih/color"
)

type confStep struct{}

func (s confStep) Execute(c *context) bool {

	file, _ := ioutil.ReadFile(
		c.RepositoryRoot + "/.git/ff.conf.json",
	)

	fmt.Println("")
	fmt.Println(color.YellowString("[confguration]"))
	fmt.Println("")
	fmt.Println(string(file))
	fmt.Println("")

	c.CurrentStep = &finalStep{}

	return true
}

func (s confStep) Stepname() string {
	return "show-configuration"
}
