package main

import (
	"fmt"
)

type configStep struct{}

func (s configStep) Execute(c *context) bool {

	for i := range c.getInput() {
		fmt.Println(i)
	}

	c.CurrentStep = &finalStep{}

	return true
}

func (s configStep) Stepname() string {
	return "config"
}
