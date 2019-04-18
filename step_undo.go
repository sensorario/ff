package main

type undoStep struct{}

func (s undoStep) Execute(c *context) bool {
	gitUndo := &gitCommand{
		c.Logger,
		[]string{"revert", "HEAD"},
		"Cant undo last commit",
	}

	_ = gitUndo.Execute()

	c.CurrentStep = &finalStep{}

	return true
}

func (s undoStep) Stepname() string {
	return "reset-undo"
}
