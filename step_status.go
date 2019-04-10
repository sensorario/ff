package main

import (
	"fmt"

	"github.com/fatih/color"
)

type StatusStep struct{}

func (s StatusStep) Execute(c *Context) bool {
	branchName := c.currentBranch()

	fmt.Println(
		"Current branch is ",
		color.GreenString(branchName),
	)

	return false
}

func (s StatusStep) Stepname() string {
	return "status"
}
