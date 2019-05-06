package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

type checkTagStep struct{}

func (s checkTagStep) Execute(c *context) bool {
	canTag := true
	if _, err := exec.Command("git", "log").Output(); err != nil {
		canTag = false
	}

	if canTag == true {

		cmdName := "git"
		cmdArgs := []string{"describe", "--tags"}
		if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {

			applyTag := false

			if !c.conf.Features.StopAskingForTags {
				fmt.Println(
					color.YellowString(
						"Want you start tagging repository now? (yes/no)",
					),
				)

				reader := bufio.NewReader(os.Stdin)
				response, _ := reader.ReadString('\n')
				if string(response) == "yes\n" {
					applyTag = true
				}
			}

			c.conf.Features.ApplyFirstTag = applyTag

			if c.conf.Features.ApplyFirstTag == true {
				fmt.Println(
					color.YellowString(
						"Tag!",
					),
				)
				gitInit := &gitCommand{
					Logger:  c.Logger,
					args:    []string{"tag", "v0.0.0"},
					message: "Cant apply first tag",
				}
				_ = gitInit.Execute()
			}

			if !c.conf.Features.StopAskingForTags {
				fmt.Println(
					color.YellowString(
						"Want you remember this response? (yes/no)",
					),
				)

				reader := bufio.NewReader(os.Stdin)
				response, _ := reader.ReadString('\n')

				if string(response) == "yes\n" {
					conf, _ := readConfiguration(c.RepositoryRoot)
					conf.Features.StopAskingForTags = true
					confIndented, _ := json.MarshalIndent(conf, "", "  ")
					ioutil.WriteFile(".git/ff.conf.json", confIndented, 0644)
				}
			}
		}

	}

	c.CurrentStep = &inputReadingStep{}

	return true
}

func (s checkTagStep) Stepname() string {
	return "check-tag"
}
