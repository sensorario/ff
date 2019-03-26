package main

import "testing"

func TestHelloWorld(t *testing.T) {
	c := &Context{}
	cs := CommitStep{false}
	cs.Execute(c)
	if c.CurrentStep.Stepname() != "commit" {
		t.Fatal(c.CurrentStep.Stepname(), "not expected")
	}
}
