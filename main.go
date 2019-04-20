package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	currentFolder, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	repositoryRoot, err := findRepository(currentFolder)

	if err == nil {
		_ = os.Chdir(repositoryRoot)
	}

	repositoryExists := !os.IsNotExist(err)
	if !repositoryExists {
		fmt.Println(color.RedString("No repository found"))
	}

	conf, err := ReadConfiguration(repositoryRoot)

	logger := genLog()

	if !repositoryExists {
		for {
			fmt.Println(color.YellowString("Want you create new git repository here? (yes/no)"))

			reader := bufio.NewReader(os.Stdin)
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

				// @todo ask what is the development branche

				confIndented, _ := json.MarshalIndent(conf, "", "  ")
				if _, err := os.Stat(repositoryRoot + "/.git/ff.conf.json"); os.IsNotExist(err) {
					_ = ioutil.WriteFile(repositoryRoot+"/.git/ff.conf.json", confIndented, 0644)
				}
				fmt.Println(color.YellowString("configuration file created"))

				if _, err := os.Stat(repositoryRoot + "/README.md"); os.IsNotExist(err) {
					os.Create(repositoryRoot + "/README.md")
					fmt.Println(color.YellowString("readme file added"))
				} else {
					fmt.Println(color.YellowString("readme file preserved"))
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
	} else {
		confIndented, _ := json.MarshalIndent(conf, "", "  ")
		if _, err := os.Stat(repositoryRoot + "/.git/ff.conf.json"); os.IsNotExist(err) {
			_ = ioutil.WriteFile(repositoryRoot+"/.git/ff.conf.json", confIndented, 0644)
		}
	}

	devBranchName := conf.Branches.Historical.Development

	cntxt := context{
		CurrentStep:   &checkTagStep{},
		Logger:        logger,
		devBranchName: devBranchName,
		conf:          conf,
	}

	cntxt.enterStep()

	for cntxt.CurrentStep.Execute(&cntxt) {
		cntxt.enterStep()
	}
}
