package ff

import (
	"github.com/fatih/color"
	"regexp"
)

type wokingDirStep struct{}

func (s wokingDirStep) Execute(c *Context) bool {
	re := regexp.MustCompile(`(?m)nothing to commit, working tree clean`)

	isWorkingTreeClean := false
	for range re.FindAllString(string(c.status()), -1) {
		isWorkingTreeClean = true
		c.Logger.Info(color.RedString("working tree clean"))
	}

	if isWorkingTreeClean {
		return false
	}

	c.CurrentStep = &commitStep{}

	return true
}

func (s wokingDirStep) Stepname() string {
	return "check-working-directory"
}
