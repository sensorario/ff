package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/sensorario/gol"
)

func genLog() gol.Logger {
	if envLogPath := os.Getenv("FF_LOG_PATH"); envLogPath != "" {
		return gol.NewCustomLogger(envLogPath)
	}

	dir, _ := os.Getwd()
	return gol.NewCustomLogger(dir + "/.git/")
}

func main() {
	logger := genLog()

	dir, _ := os.Getwd()
	if _, err := os.Stat(dir + "/.git/"); os.IsNotExist(err) {
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println(color.RedString("No repository found"))
			fmt.Println(color.YellowString("Want you create new git repository here? (yes/no)"))
			response, _ := reader.ReadString('\n')
			fmt.Println(response)
			fmt.Println("no")

			if string(response) == "yes\n" {
				gitInit := &gitCommand{
					args:    []string{"init"},
					message: "Cant create new branch",
				}
				_ = gitInit.Execute()

				dir, _ := os.Getwd()
				os.Create(dir + "/README.md")

				gitInit = &gitCommand{
					Logger:  logger,
					args:    []string{"add", "."},
					message: "Cant stage everything",
				}
				_ = gitInit.Execute()

				gitInit = &gitCommand{
					Logger:  logger,
					args:    []string{"commit", "-m", "start"},
					message: "Cant commit",
				}
				_ = gitInit.Execute()

				gitInit = &gitCommand{
					Logger:  logger,
					args:    []string{"tag", "v0.0.0"},
					message: "Cant apply first tag",
				}
				_ = gitInit.Execute()

				os.Exit(0)
			}

			if string(response) == "no\n" {
				os.Exit(0)
			}
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
