package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/sensorario/gol"
)

func genLog() gol.Logger {
	if envLogPath := os.Getenv("FF_LOG_PATH"); envLogPath != "" {
		return gol.NewCustomLogger(envLogPath)
	}

	dir, _ := os.Getwd()
	return gol.NewCustomLogger(dir + "/.git")
}

func main() {
	logger := genLog()

	dir, _ := os.Getwd()
	if _, err := os.Stat(dir + "/.git"); os.IsNotExist(err) {
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println(color.RedString("No repository found"))
			fmt.Println(color.YellowString("Want you create new git repository here? (yes/no)"))
			response, _ := reader.ReadString('\n')

			if string(response) == "yes\n" {
				gitInit := &gitCommand{
					args:    []string{"init"},
					message: "Cant create new branch",
				}
				_ = gitInit.Execute()
				fmt.Println(color.YellowString(
					"repository initialized",
				))

				dir, _ := os.Getwd()
				os.Create(dir + "/README.md")
				fmt.Println(color.YellowString(
					"readme file added",
				))

				gitInit = &gitCommand{
					Logger:  logger,
					args:    []string{"add", "."},
					message: "Cant stage everything",
				}
				_ = gitInit.Execute()
				fmt.Println(color.YellowString(
					"readme file staged",
				))

				gitInit = &gitCommand{
					Logger:  logger,
					args:    []string{"commit", "-m", "start"},
					message: "Cant commit",
				}
				_ = gitInit.Execute()
				fmt.Println(color.YellowString(
					"first commit committed",
				))

				gitInit = &gitCommand{
					Logger:  logger,
					args:    []string{"tag", "v0.0.0"},
					message: "Cant apply first tag",
				}
				_ = gitInit.Execute()
				fmt.Println(color.YellowString(
					"first tag v0.0.0 added",
				))

				os.Exit(0)
			}

			if string(response) == "no\n" {
				os.Exit(0)
			}
		}
	}

	// fatal: your current branch 'master' does not have any commits yet
	canTag := true
	if _, err := exec.Command("git", "log").Output(); err != nil {
		canTag = false
	}

	if canTag == true {
		cmdName := "git"
		cmdArgs := []string{"describe", "--tags"}
		if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
			gitInit := &gitCommand{
				Logger:  logger,
				args:    []string{"tag", "v0.0.0"},
				message: "Cant apply first tag",
			}
			_ = gitInit.Execute()
		}
	}

	cntxt := context{
		CurrentStep: &inputReadingStep{},
		Logger:      logger,
	}

	cntxt.enterStep()

	for cntxt.CurrentStep.Execute(&cntxt) {
		cntxt.enterStep()
	}
}
