package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/sensorario/gol"
)

type gitCommand struct {
	Logger  gol.Logger
	args    []string
	message string
	conf    jsonConf
}

func (gc gitCommand) ErrorMessage() string {
	return gc.message
}

func (gc gitCommand) Arguments() []string {
	return gc.args
}

func (gc *gitCommand) Execute() string {
    fmt.Println("Execute")
	var cmdOut []byte
	var err error
	cmdName := "git"
	cmdArgs := gc.Arguments()
    fmt.Println(cmdArgs)

	cmdOut, err = exec.Command(cmdName, cmdArgs...).Output()
    fmt.Println(string(cmdOut))

	if gc.conf.Features.EnableGitCommandLog == true {
		gc.Logger.Info(color.YellowString(strings.Join(cmdArgs, " ")))
		gc.Logger.Info(color.GreenString("<<< Response\n") + string(cmdOut))
	}

    gc.Logger.Info("@@")
    gc.Logger.Info(cmdName)
    // fmt.Println("Cant tag")
    // fmt.Println(err)
    // fmt.Println(string(cmdOut))
    // fmt.Println(cmdName)
    // fmt.Println(gc.Arguments())
    // sembra dare errore quando viene eseguito git statu
    // e non ci sono files aggiunti nella staging area?
    // da capire
    // os.Exit(1)

	if err != nil {

		fmt.Println(color.RedString(err.Error()))
		fmt.Println(color.RedString(gc.ErrorMessage()))

		if err.Error() == "exit status 128" {
			gc.Logger.Error(color.RedString("Error"))
			gc.Logger.Error(color.RedString(err.Error()))
			gc.Logger.Error(color.RedString(gc.ErrorMessage()))
		}

        fmt.Println("222")
		os.Exit(22)
	} else {
        fmt.Println(cmdName)
    }

	return string(cmdOut)
}
