package main

import (
	"github.com/fatih/color"
)

type fFStep interface {
	Execute(c *Context) bool
	Stepname() string
}

func (c *Context) enterStep() {
	c.Logger.Info(color.RedString("[step/" + c.CurrentStep.Stepname() + "]"))
}
