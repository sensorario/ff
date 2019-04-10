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
	}

	cmdOut := gitDescribeTags.Execute()

	fmt.Print("current tag: ", color.GreenString(string(cmdOut)))

	tagName := ""

	branchName := ""
	re := regexp.MustCompile(`On branch [\w\/\#\-\.]{0,}`)
	for _, match := range re.FindAllString(string(cmdOut), -1) {
		branchName = strings.ReplaceAll(match, "On branch ", "")
	}

	mt := meta{string(cmdOut), branchName}

	tagName = mt.NextPatchTag()

	fmt.Println("next tag:   ", color.GreenString(tagName))

	gitTag := &gitCommand{
		c.Logger,
		[]string{"tag", tagName, "-f"},
		"cant tag",
	}
	_ = gitTag.Execute()

	c.CurrentStep = &finalStep{}

	return true
}

func (s tagStep) Stepname() string {
	return "force-tag"
}
