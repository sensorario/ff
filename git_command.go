package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

type GitCommand struct {
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
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		if err.Error() == "exit status 128" {
			fmt.Println(color.RedString(gc.ErrorMessage()))
		}
		os.Exit(1)
	}
	return string(cmdOut)
}
