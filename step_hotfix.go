package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"strings"
)

type HotfixStep struct{}

func (s *HotfixStep) Execute(c *Context) bool {
	cmdName := "git"
	cmdArgs := []string{"checkout", "master"}
	if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		if err.Error() == "exit status 128" {
			fmt.Println(color.RedString("git repository not found"))
		}
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(color.RedString("Hotfix description: "))
	featureName, _ := reader.ReadString('\n')
	featureName = strings.ReplaceAll(featureName, " ", "-")
	featureName = strings.ReplaceAll(featureName, "\n", "")

	fmt.Print(
		"Hotfix: ",
	)

	featureBranchName := "hotfix/" + featureName
	fmt.Println(color.YellowString(featureBranchName))

	cmdStartBranch := "git"
	arguments := []string{"checkout", "-b", featureBranchName}
	fmt.Println(arguments)
	if _, err := exec.Command(cmdStartBranch, arguments...).Output(); err != nil {
		fmt.Println(color.RedString(err.Error()))
		os.Exit(1)
	}

	c.CurrentStep = &FinalStep{}

	return true
}

func (s *HotfixStep) Stepname() string {
	return "checkout master"
}
