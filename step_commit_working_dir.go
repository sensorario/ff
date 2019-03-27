package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/fatih/color"
)

type WorkingDirStep struct{}

func (s *WorkingDirStep) Execute(c *Context) bool {
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

	re := regexp.MustCompile(`(?m)nothing to commit, working tree clean`)

	isWorkingTreeClean := false
	for _, _ = range re.FindAllString(string(cmdOut), -1) {
		isWorkingTreeClean = true
		fmt.Println(color.RedString("working tree clean"))
	}

	if isWorkingTreeClean {
		return false
	}

	c.CurrentStep = &CommitStep{}

	return true
}

func (s *WorkingDirStep) Stepname() string {
	return "check-working-directory"
}
