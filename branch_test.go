package main

import (
	"testing"
)

func TestAnalyze(t *testing.T) {
	br := branch{"feature/branch-semantico/master"}

	if br.isFeature() == false {
		t.Errorf("branch should be recognyzed as feature")
	}

	if br.destination() != "master" {
		t.Errorf("destination branch should be master")
	}
}

func TestExtractCurrentBranch(t *testing.T) {
	br := branch{"1.0"}
	if br.name != "1.0" {
		t.Errorf("Oops! Branch detection fails!")
	}
}
