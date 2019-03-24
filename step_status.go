package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type StatusStep struct{}

func (s *StatusStep) Execute(c *Context) bool {
	c.EnterStep()
	var (
		cmdOut []byte
		err    error
	)

	cmdName := "git"
	cmdArgs := []string{"status"}

	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		if err.Error() == "exit status 128" {
			fmt.Println("\033[1;31mgit repository not found\033[0m")
		}
		os.Exit(1)
	}

	re := regexp.MustCompile(`On branch [\w\/\#\-]{0,}`)

	isDevelopmentBranch := false

	for _, match := range re.FindAllString(string(cmdOut), -1) {
		branchName := strings.ReplaceAll(match, "On branch ", "")
		fmt.Println("Current branch: ", branchName)
		isDevelopmentBranch = branchName == "master"
	}

	if isDevelopmentBranch {
		fmt.Println(color.GreenString("Development branch"))
		fmt.Println("you can create feature branch from here")
	}

	return false
}

func (s *StatusStep) Stepname() string {
	return "status"
}
