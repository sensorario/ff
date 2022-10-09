package ff

type undoStep struct{}

func (s undoStep) Execute(c *Context) bool {
	gitUndo := &GitCommand{
		c.Logger,
		[]string{"revert", "HEAD"},
		"Cant undo last commit",
		c.Conf,
	}

	_ = gitUndo.Execute()

	c.CurrentStep = &finalStep{}

	return true
}

func (s undoStep) Stepname() string {
	return "reset-undo"
}
