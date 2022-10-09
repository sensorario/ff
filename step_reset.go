package ff

type resetStep struct{}

func (s resetStep) Execute(c *Context) bool {
	gitAddEverything := &GitCommand{
		c.Logger,
		[]string{"add", "."},
		"Cant add working directory to stage",
		c.Conf,
	}

	_ = gitAddEverything.Execute()

	gitResetHard := &GitCommand{
		c.Logger,
		[]string{"reset", "--hard"},
		"Cant reset",
		c.Conf,
	}

	_ = gitResetHard.Execute()

	c.CurrentStep = &finalStep{}

	return true
}

func (s resetStep) Stepname() string {
	return "reset-everything"
}
