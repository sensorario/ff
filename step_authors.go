package ff

import (
	"fmt"
	"regexp"
	"strings"
)

type authorsStep struct{}

func (s authorsStep) Execute(c *Context) bool {
	gitStatus := &GitCommand{
		c.Logger,
		[]string{"status"},
		"cant get status",
		c.Conf,
	}
	cmdOut := gitStatus.Execute()
	branchName := ""
	re := regexp.MustCompile(`On branch [\w\/\#\-\.]{0,}`)
	for _, match := range re.FindAllString(string(cmdOut), -1) {
		branchName = strings.ReplaceAll(match, "On branch ", "")
	}

	args := []string{
		"shortlog",
		branchName,
		"--summary",
		"--numbered",
	}

	gitShortLog := &GitCommand{
		c.Logger,
		args,
		"Cant lista authors",
		c.Conf,
	}

	out := gitShortLog.Execute()

	fmt.Println(out)

	c.CurrentStep = &finalStep{}

	return true
}

func (s authorsStep) Stepname() string {
	return "repository-authors"
}
