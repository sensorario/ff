package main

func main() {
	t := Context{CurrentStep: &InputReadingStep{}}
	for t.CurrentStep.Execute(&t) {
	}
}
