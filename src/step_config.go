package main

type configStep struct{}

func (s configStep) Execute(c *context) bool {

	c.Logger.Info("ciaone")

	c.CurrentStep = &finalStep{}

	return true
}

func (s configStep) Stepname() string {
	return "config"
}
