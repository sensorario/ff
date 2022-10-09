package ff

import "fmt"

type fetchAllStep struct{}

func (s fetchAllStep) Execute(c *Context) bool {
	args := []string{
		"fetch",
		"--all",
	}

	gitFetchAll := &GitCommand{
		c.Logger,
		args,
		"Cant fetch all",
		c.Conf,
	}

	out := gitFetchAll.Execute()

	fmt.Println(out)

	c.CurrentStep = &finalStep{}

	return true
}

func (s fetchAllStep) Stepname() string {
	return "fetch-all"
}
