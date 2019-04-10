package main

import (
	"os"

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
	cntxt := context{
		CurrentStep: &inputReadingStep{},
		Logger:      genLog(),
	}

	cntxt.enterStep()

	for cntxt.CurrentStep.Execute(&cntxt) {
		cntxt.enterStep()
	}
}
