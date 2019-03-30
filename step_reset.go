package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

type ResetStep struct{}

func (s *ResetStep) Execute(c *Context) bool {
	gitAddEverything := &GitCommand{[]string{"add", "."}, "Cant add working directory to stage"}
	_ = gitAddEverything.Execute()

	gitResetHard := &GitCommand{[]string{"reset", "--hard"}, "Cant reset"}
	_ = gitResetHard.Execute()

	c.CurrentStep = &FinalStep{}

	return true
}

func (s *ResetStep) Stepname() string {
	return "reset-everything"
}
