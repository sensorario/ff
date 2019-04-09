package main

type ResetStep struct{}

func (s ResetStep) Execute(c *Context) bool {
	gitAddEverything := &gitCommand{
		c.Logger,
		[]string{"add", "."},
		"Cant add working directory to stage",
	}

	_ = gitAddEverything.Execute()

	gitResetHard := &gitCommand{
		c.Logger,
		[]string{"reset", "--hard"},
		"Cant reset",
	}

	_ = gitResetHard.Execute()

	c.CurrentStep = &FinalStep{}

	return true
}

func (s ResetStep) Stepname() string {
	return "reset-everything"
}
