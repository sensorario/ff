package main

import (
	"github.com/fatih/color"
	"github.com/sensorario/gol"
)

type FussyStepInterface interface {
	Execute(c *Context) bool
	Stepname() string
}

type Context struct {
	CurrentStep FussyStepInterface
	Exit        bool
	Logger      gol.Logger
}

func (c *Context) EnterStep() {
	c.Logger.Info(color.RedString("[step/" + c.CurrentStep.Stepname() + "]"))
}
