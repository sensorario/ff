package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

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
	conf := ReadConfiguration()
	dir, _ := os.Getwd()

	confIndented, _ := json.MarshalIndent(conf, "", "  ")

	// salvo configurazione se non esiste
	if _, err := os.Stat(dir + "/.git/conf.json"); os.IsNotExist(err) {
		_ = ioutil.WriteFile(
			dir+"/.git/conf.json",
			confIndented,
			0644,
		)
	}

	logger := genLog()

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

				// check if file exists
				dir, _ := os.Getwd()
				if _, err := os.Stat(dir + "/README.md"); os.IsNotExist(err) {
					os.Create(dir + "/README.md")
					fmt.Println(color.YellowString(
						"readme file added",
					))
				} else {
					fmt.Println(color.YellowString(
						"readme file preserved",
					))
				}

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

	devBranchName := conf.Branches.Historical.Development

	cntxt := context{
		CurrentStep:   &checkTagStep{},
		Logger:        logger,
		devBranchName: devBranchName,
	}

	cntxt.enterStep()

	for cntxt.CurrentStep.Execute(&cntxt) {
		cntxt.enterStep()
	}
}
