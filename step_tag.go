package ff

import (
	"fmt"
	"regexp"
	"strings"
    // "os"

	"github.com/fatih/color"
)

type tagStep struct{}

func (s tagStep) Execute(c *Context) bool {
	gitDescribeTags := &GitCommand{
		c.Logger,
		[]string{"describe", "--tags"},
		"cant force tag",
		c.Conf,
	}

	cmdOut := gitDescribeTags.Execute()

	initialTag := strings.TrimSuffix(
		string(cmdOut),
		"\n",
	)

	fmt.Println("current tag: ", color.GreenString(initialTag))

	branchName := ""
	re := regexp.MustCompile(`On branch [\w\/\#\-\.]{0,}`)
	for _, match := range re.FindAllString(string(cmdOut), -1) {
		branchName = strings.ReplaceAll(match, "On branch ", "")
	}

    // meta struct {describe, branch}
	mt := meta{string(cmdOut), branchName}

    // qui bisogna assicurarsi che meta sia nel formato corretto
    // perche' NextPatchTag da per scontato che il formato sia
    // di un certo tipo

	tagName := ""
	tagName = mt.NextPatchTag()
	tagName = strings.TrimSuffix(string(tagName), "\n",)

    // fmt.Println(mt)
    // fmt.Println(tagName)
    // os.Exit(1);

	if initialTag == tagName {
		fmt.Println("")
		fmt.Println(color.YellowString("\tsame tag found"))
		fmt.Println(color.YellowString("\tno new tag will be added"))
		fmt.Println("")
	} else {
		fmt.Println(color.YellowString("different tag"))
	}

	fmt.Println("next tag:   ", color.GreenString(tagName))

	gitTag := &GitCommand{
		c.Logger,
		[]string{"tag", tagName, "-f"},
		"cant tag",
		c.Conf,
	}
	_ = gitTag.Execute()

	c.CurrentStep = &finalStep{}

	return true
}

func (s tagStep) Stepname() string {
	return "force-tag"
}
