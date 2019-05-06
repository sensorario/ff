package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

type tagStep struct{}

func (s tagStep) Execute(c *context) bool {
	gitDescribeTags := &gitCommand{
		c.Logger,
		[]string{"describe", "--tags"},
		"cant force tag",
		c.conf,
	}

	cmdOut := gitDescribeTags.Execute()

	initialTag := strings.TrimSuffix(
		string(cmdOut),
		"\n",
	)

	fmt.Println("current tag: ", color.GreenString(initialTag))

	tagName := ""

	branchName := ""
	re := regexp.MustCompile(`On branch [\w\/\#\-\.]{0,}`)
	for _, match := range re.FindAllString(string(cmdOut), -1) {
		branchName = strings.ReplaceAll(match, "On branch ", "")
	}

	mt := meta{string(cmdOut), branchName}

	tagName = mt.NextPatchTag()

	tagName = strings.TrimSuffix(
		string(tagName),
		"\n",
	)

	if initialTag == tagName {
		fmt.Println("")
		fmt.Println(color.YellowString("\tsame tag found"))
		fmt.Println(color.YellowString("\tno new tag will be added"))
		fmt.Println("")
	} else {
		fmt.Println(color.YellowString("different tag"))
	}

	fmt.Println("next tag:   ", color.GreenString(tagName))

	gitTag := &gitCommand{
		c.Logger,
		[]string{"tag", tagName, "-f"},
		"cant tag",
		c.conf,
	}
	_ = gitTag.Execute()

	c.CurrentStep = &finalStep{}

	return true
}

func (s tagStep) Stepname() string {
	return "force-tag"
}
