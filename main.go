package main

import (
	"os"

	"github.com/sensorario/gol"
)

func genLog() gol.Logger {
	if envLogPath := os.Getenv("FF_LOG_PATH"); envLogPath != "" {
		return gol.NewCustomLogger(envLogPath)
	}

	return gol.NewLogger("ff")
}

func main() {
	context := Context{
		CurrentStep: &InputReadingStep{},
		Logger:      genLog(),
	}

	context.EnterStep()

	for context.CurrentStep.Execute(&context) {
		context.EnterStep()
	}
}
