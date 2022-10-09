package ff

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fatih/color"
)

type configStep struct{}

func (s configStep) Execute(c *Context) bool {

	if len(c.getInput()) < 3 {
		confIndented, _ := json.MarshalIndent(c.Conf, "", "  ")
		fmt.Println(string(confIndented))
		c.CurrentStep = &finalStep{}
		return true
	}

	inputs := c.getInput()
	feature := inputs[2] // ff config <feature>
	fmt.Println(feature)

	knownConfigs := []string{
		"applyFirstTag",
		"disableUndoCommand",
		"enableGitCommandLog",
		"forceOnPublish",
		"pushTagsOnPublish",
		"removeRemotelyMerged",
		"stopAskingForTags",
		"tagAfterMerge",
		"lang",
	}

	found := false
	for _, f := range knownConfigs {
		if f == feature {
			found = true
		}
	}

	if found {
		if feature == "removeRemotelyMerged" {
			c.Conf.Features.RemoveRemotelyMerged = c.Conf.Features.RemoveRemotelyMerged == false
		}

		if feature == "tagAfterMerge" {
			c.Conf.Features.TagAfterMerge = c.Conf.Features.TagAfterMerge == false
		}

		if feature == "disableUndoCommand" {
			c.Conf.Features.DisableUndoCommand = c.Conf.Features.DisableUndoCommand == false
		}

		if feature == "stopAskingForTags" {
			c.Conf.Features.StopAskingForTags = c.Conf.Features.StopAskingForTags == false
		}

		if feature == "applyFirstTag" {
			c.Conf.Features.ApplyFirstTag = c.Conf.Features.ApplyFirstTag == false
		}

		if feature == "enableGitCommandLog" {
			c.Conf.Features.EnableGitCommandLog = c.Conf.Features.EnableGitCommandLog == false
		}

		if feature == "forceOnPublish" {
			c.Conf.Features.ForceOnPublish = c.Conf.Features.ForceOnPublish == false
		}

		if feature == "pushTagsOnPublish" {
			c.Conf.Features.PushTagsOnPublish = c.Conf.Features.PushTagsOnPublish == false
		}

		if feature == "lang" {
			if len(inputs) == 4 {
				if inputs[3] == "en" || inputs[3] == "it" {
					c.Conf.Features.Lang = inputs[3]
				}

				if inputs[3] != "en" && inputs[3] != "it" {
					fmt.Println(color.RedString(
						strings.Join([]string{"Language", inputs[3], "is not available"}, " "),
					))
				}
			} else {
				fmt.Println(color.RedString(
					strings.Join([]string{
						"Wrong parameter count",
					}, " "),
				))
			}
		}

		confIndented, _ := json.MarshalIndent(c.Conf, "", "  ")
		ioutil.WriteFile(".git/ff.conf.json", confIndented, 0644)
	} else {
		fmt.Println(color.RedString("Command not found"))
	}

	c.CurrentStep = &finalStep{}

	return true
}

func (s configStep) Stepname() string {
	return "config"
}
