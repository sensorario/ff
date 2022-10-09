package ff

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

type CheckTagStep struct{}

func (s CheckTagStep) Execute(c *Context) bool {
	canTag := true
	if _, err := exec.Command("git", "log").Output(); err != nil {
		canTag = false
	}

	if canTag == true {

		// check remote origin
		cmdGetUrlArgs := []string{"remote", "get-url", "origin"}
		outputRemote, err := exec.Command("git", cmdGetUrlArgs...).Output()
		c.setRemote(string(outputRemote))

		// check current version
		cmdArgs := []string{"describe", "--tags"}
		outputWithVersion, err := exec.Command("git", cmdArgs...).Output()
		c.setCurrentVersion(string(outputWithVersion))

		if err != nil {

			applyTag := false

			if !c.Conf.Features.StopAskingForTags {
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

			c.Conf.Features.ApplyFirstTag = applyTag

			if c.Conf.Features.ApplyFirstTag == true {
				fmt.Println(
					color.YellowString(
						"Tag!",
					),
				)
				gitInit := &GitCommand{
					Logger:  c.Logger,
					Args:    []string{"tag", "v0.0.0"},
					Message: "Cant apply first tag",
				}
				_ = gitInit.Execute()
			}

			if !c.Conf.Features.StopAskingForTags {
				fmt.Println(
					color.YellowString(
						"Want you remember this response? (yes/no)",
					),
				)

				reader := bufio.NewReader(os.Stdin)
				response, _ := reader.ReadString('\n')

				if string(response) == "yes\n" {
					conf, _ := ReadConfiguration(c.RepositoryRoot)
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

func (s CheckTagStep) Stepname() string {
	return "check-tag"
}
