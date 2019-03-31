package main

import (
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/sensorario/gol"
)

type GitCommand struct {
	Logger  gol.Logger
	args    []string
	message string
}

func (gc *GitCommand) ErrorMessage() string {
	return gc.message
}

func (gc *GitCommand) Arguments() []string {
	return gc.args
}

func (gc *GitCommand) Execute() string {
	var cmdOut []byte
	var err error
	cmdName := "git"
	cmdArgs := gc.Arguments()

	if cmdOut, _ = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		if err.Error() == "exit status 128" {
			gc.Logger.Error(color.RedString(err.Error()))
			gc.Logger.Error(color.RedString(gc.ErrorMessage()))
		}
		os.Exit(1)
	}

	return string(cmdOut)
}
