package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/fatih/color"
)

type configStep struct{}

func (s configStep) Execute(c *context) bool {

	inputs := c.getInput()
	feature := inputs[2] // ff config *****
	fmt.Println(feature)

	knownConfigs := []string{
		"tagAfterMerge",
		"disableUndoCommand",
		"stopAskingForTags",
		"applyFirstTag",
		"enableGitCommandLog",
		"forceOnPublish",
		"pushTagsOnPublish",
	}

	found := false
	for _, f := range knownConfigs {
		if f == feature {
			found = true
		}
	}

	if found {
		if feature == "tagAfterMerge" {
			c.conf.Features.TagAfterMerge = c.conf.Features.TagAfterMerge == false
		}

		if feature == "disableUndoCommand" {
			c.conf.Features.DisableUndoCommand = c.conf.Features.DisableUndoCommand == false
		}

		confIndented, _ := json.MarshalIndent(c.conf, "", "  ")
		ioutil.WriteFile(".git/ff.conf.json", confIndented, 0644)
	} else {
		fmt.Println(color.RedString("Comando NON trovato"))
	}

	c.CurrentStep = &finalStep{}

	return true
}

func (s configStep) Stepname() string {
	return "config"
}
