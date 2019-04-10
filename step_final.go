package main

type finalStep struct{}

func (s finalStep) Execute(c *Context) bool {
	return false
}

func (s finalStep) Stepname() string {
	return "final"
}
