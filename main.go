package main

import "github.com/sensorario/gol"

func main() {

	context := Context{
		CurrentStep: &InputReadingStep{},
		Logger:      gol.Logger{Application: "ff", LogFile: "info"},
	}

	context.EnterStep()

	for context.CurrentStep.Execute(&context) {
		context.EnterStep()
	}

}
