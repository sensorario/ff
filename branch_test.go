package main

import (
	"testing"
)

func TestAnalyze(t *testing.T) {
	branch := SemBranch{"feature/branch-semantico/master"}

	if branch.IsFeature() == false {
		t.Errorf("branch should be recognyzed as feature")
	}

	if branch.Destination() != "master" {
		t.Errorf("destination branch shoudl be master")
	}
}
