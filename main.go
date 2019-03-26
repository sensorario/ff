package main

func main() {
	context := Context{CurrentStep: &InputReadingStep{}}
	context.EnterStep()
	for context.CurrentStep.Execute(&context) {
		context.EnterStep()
	}
}
