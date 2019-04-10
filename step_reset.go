package main

type ResetStep struct{}

func (s ResetStep) Execute(c *context) bool {
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

	c.CurrentStep = &finalStep{}

	return true
}

func (s ResetStep) Stepname() string {
	return "reset-everything"
}
