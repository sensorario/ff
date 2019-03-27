package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

type CompleteFeatureStep struct{}

func (s *CompleteFeatureStep) Execute(c *Context) bool {
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
	branchName := ""
	for _, match := range re.FindAllString(string(cmdOut), -1) {
		branchName = strings.ReplaceAll(match, "On branch ", "")
	}

	fmt.Println(color.RedString("leaving: " + branchName))

	cmdArgs = []string{"checkout", "master"}
	if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		if err.Error() == "exit status 128" {
			fmt.Println(color.RedString("git repository not found"))
		}
		os.Exit(1)
	}

	cmdArgs = []string{"describe", "--tags"}
	cmdOut, _ = exec.Command(cmdName, cmdArgs...).Output()
	isHotfix := strings.HasPrefix(branchName, "hotfix/")
	isFeature := strings.HasPrefix(branchName, "feature/")

	if isHotfix {
		meta := Meta{string(cmdOut), branchName}
		fmt.Println("next tag: ", color.RedString(meta.NextPatchTag()))
	}

	if isFeature {
		meta := Meta{string(cmdOut), branchName}
		fmt.Println("next tag: ", color.RedString(meta.NextMinorTag()))
	}

	cmdArgs = []string{"merge", "--no-ff", branchName}
	if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		if err.Error() == "exit status 128" {
			fmt.Println(color.RedString("git repository not found"))
		}
		os.Exit(1)
	}

	cmdArgs = []string{"branch", "-D", branchName}
	if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		if err.Error() == "exit status 128" {
			fmt.Println(color.RedString("something went wrong during deletion"))
		}
		os.Exit(1)
	}
	fmt.Println(color.GreenString("branch " + branchName + " deleted"))

	c.CurrentStep = &FinalStep{}

	return true
}

func (s *CompleteFeatureStep) Stepname() string {
	return "checkout master"
}
