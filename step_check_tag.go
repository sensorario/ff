package main

import "os/exec"

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
			gitInit := &gitCommand{
				Logger:  c.Logger,
				args:    []string{"tag", "v0.0.0"},
				message: "Cant apply first tag",
			}
			_ = gitInit.Execute()
		}
	}

	c.CurrentStep = &inputReadingStep{}

	return true
}

func (s checkTagStep) Stepname() string {
	return "check-tag"
}
