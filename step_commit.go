package ff

import (
	"bufio"
	"fmt"
	"os"

    "github.com/sensorario/branch"
)

type commitStep struct{}

func (s commitStep) Execute(c *Context) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is the commit message: ")

	text, _ := reader.ReadString('\n')

	gitAddAll := &GitCommand{
		c.Logger,
		[]string{"add", "."},
		"Cant add files",
		c.Conf,
	}

	gitAddAll.Execute()

	name := c.currentBranch()
	sem := branch.Branch{name}

	result := sem.CommitPrefix() + text
	gitCommit := &GitCommand{
		c.Logger,
		[]string{"commit", "-m", result},
		"Cant add more files",
		c.Conf,
	}

	gitCommit.Execute()

	c.CurrentStep = &finalStep{}

	return false
}

func (s commitStep) Stepname() string {
	return "commit"
}
