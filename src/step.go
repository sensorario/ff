package main

import (
	"github.com/fatih/color"
)

type fFStep interface {
	Execute(c *context) bool
	Stepname() string
}

func (c *context) enterStep() {
	c.Logger.Info(color.RedString("[step/" + c.CurrentStep.Stepname() + "]"))
}
