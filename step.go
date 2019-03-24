package main

import (
	"fmt"
	"github.com/fatih/color"
)

type FussyStepInterface interface {
	Execute(c *Context) bool
	Stepname() string
}

type Context struct {
	CurrentStep FussyStepInterface
	Exit        bool
}

func (c *Context) EnterStep() {
	fmt.Println(color.RedString("[step/" + c.CurrentStep.Stepname() + "]"))
}
