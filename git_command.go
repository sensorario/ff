package ff

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/sensorario/gol"
)

type GitCommand struct {
	Logger  gol.Logger
	Args    []string
	Message string
	Conf    JsonConf
}

func (gc GitCommand) ErrorMessage() string {
	return gc.Message
}

func (gc GitCommand) Arguments() []string {
	return gc.Args
}

func (gc *GitCommand) Execute() string {
	var cmdOut []byte
	var err error
	cmdName := "git"
	cmdArgs := gc.Arguments()

	cmdOut, err = exec.Command(cmdName, cmdArgs...).Output()

	if gc.Conf.Features.EnableGitCommandLog == true {
		gc.Logger.Info(color.YellowString(strings.Join(cmdArgs, " ")))
		gc.Logger.Info(color.GreenString("<<< Response\n") + string(cmdOut))
	}

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
			gc.Logger.Error(color.RedString(err.Error()))
			gc.Logger.Error(color.RedString(gc.ErrorMessage()))
		}

		os.Exit(1)
	}

	return string(cmdOut)
}
