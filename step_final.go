package main

type FinalStep struct{}

func (s *FinalStep) Execute(c *Context) bool {
	c.EnterStep()
	return false
}

func (s *FinalStep) Stepname() string {
	return "final"
}
