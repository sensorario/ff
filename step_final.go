package main

type finalStep struct{}

func (s finalStep) Execute(c *context) bool {
	return false
}

func (s finalStep) Stepname() string {
	return "final"
}
