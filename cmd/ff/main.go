package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/sensorario/ff"
	"github.com/sensorario/gol/v2"
)

// genLog generate log
func genLog(repositoryRoot string) gol.Logger {
	defaultLogDirectory := "/.git"

	if envLogPath := os.Getenv("FF_LOG_PATH"); envLogPath != "" {
		return gol.NewCustomLogger(envLogPath)
	}

	return gol.NewCustomLogger(
		repositoryRoot + defaultLogDirectory,
	)
}

// Everything starts from here
func main() {

	// lang := os.Getenv("LANG")
	// fmt.Println("Lang:", lang)
	// known languages: en_US.UTF-8

	currentFolder, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	repositoryRoot, err := ff.FindRepository(currentFolder)

	if err == nil {
		if err := os.Chdir(repositoryRoot); err != nil {
			log.Fatalf("Could not change to repository root %q: %v", repositoryRoot, err)
		}
	}

	repositoryExists := !os.IsNotExist(err)
	if !repositoryExists {
		fmt.Println(color.RedString("No repository found"))
	}

	conf, _ := ff.ReadConfiguration(repositoryRoot)

	logger := genLog(repositoryRoot)

	if !repositoryExists {
		guidedRepositoryCreation(logger, conf)
	} else {
		confIndented, _ := json.MarshalIndent(conf, "", "  ")
		if _, err := os.Stat(".git/ff.conf.json"); os.IsNotExist(err) {
			_ = ioutil.WriteFile(".git/ff.conf.json", confIndented, 0644)
		}
	}

	DevBranchName := conf.Branches.Historical.Development

	cntxt := ff.Context{
		RepositoryRoot: repositoryRoot,
		CurrentStep:    &ff.CheckTagStep{},
		Logger:         logger,
		DevBranchName:  DevBranchName,
		Conf:           conf,
	}

	cntxt.EnterStep()

	for cntxt.CurrentStep.Execute(&cntxt) {
		cntxt.EnterStep()
	}
}

func guidedRepositoryCreation(logger gol.Logger, conf ff.JsonConf) {
	for {
		fmt.Println(color.YellowString("Want you create new git repository here? (yes/no)"))

		reader := bufio.NewReader(os.Stdin)
		response, _ := reader.ReadString('\n')

		if string(response) == "yes\n" {
			gitInit := &ff.GitCommand{
				Args:    []string{"init"},
				Message: "Cant create new branch",
			}
			_ = gitInit.Execute()
			fmt.Println(color.YellowString(
				"repository initialized",
			))

			confIndented, _ := json.MarshalIndent(conf, "", "  ")
			if _, err := os.Stat(".git/ff.conf.json"); os.IsNotExist(err) {
				_ = ioutil.WriteFile(".git/ff.conf.json", confIndented, 0644)
			}
			fmt.Println(color.YellowString("configuration file created"))

			if _, err := os.Stat("README.md"); os.IsNotExist(err) {
				os.Create("README.md")
				fmt.Println(color.YellowString("readme file added"))
			} else {
				fmt.Println(color.YellowString("readme file preserved"))
			}

			gitInit = &ff.GitCommand{
				Logger:  logger,
				Args:    []string{"add", "."},
				Message: "Cant stage everything",
			}
			_ = gitInit.Execute()
			fmt.Println(color.YellowString(
				"readme file staged",
			))

			gitInit = &ff.GitCommand{
				Logger:  logger,
				Args:    []string{"commit", "-m", "start"},
				Message: "Cant commit",
			}
			_ = gitInit.Execute()
			fmt.Println(color.YellowString(
				"first commit committed",
			))

			fmt.Println(color.YellowString("Want you start tagging repository now? (yes/no)"))
			reader := bufio.NewReader(os.Stdin)
			response, _ := reader.ReadString('\n')
			if string(response) == "yes\n" {
				gitInit = &ff.GitCommand{
					Logger:  logger,
					Args:    []string{"tag", "v0.0.0"},
					Message: "Cant apply first tag",
				}
				_ = gitInit.Execute()
				fmt.Println(color.YellowString(
					"first tag v0.0.0 added!",
				))
			}

			os.Exit(0)
		}

		if string(response) == "no\n" {
			os.Exit(0)
		}
	}
}
