package main

type resetStep struct{}

func (s resetStep) Execute(c *context) bool {
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

func (s resetStep) Stepname() string {
	return "reset-everything"
}
