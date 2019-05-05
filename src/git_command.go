package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/sensorario/gol"
)

type gitCommand struct {
	Logger  gol.Logger
	args    []string
	message string
}

func (gc gitCommand) ErrorMessage() string {
	return gc.message
}

func (gc gitCommand) Arguments() []string {
	return gc.args
}

func (gc *gitCommand) Execute() string {
	var cmdOut []byte
	var err error
	cmdName := "git"
	cmdArgs := gc.Arguments()

	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {

		fmt.Println(color.RedString(err.Error()))
		fmt.Println(color.RedString(gc.ErrorMessage()))

		if err.Error() == "exit status 128" {
			gc.Logger.Error(color.RedString(err.Error()))
			gc.Logger.Error(color.RedString(gc.ErrorMessage()))
		}

		os.Exit(1)
	}

	return string(cmdOut)
}
