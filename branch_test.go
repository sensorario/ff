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

func TestMasterOrFeaturesBranchesAreDevelopment(t *testing.T) {
	tests := []struct {
		branch branch
		phase  string
	}{
		{branch{"feature/branch-semantico/master"}, "development"},
		{branch{"hotfix/branch-semantico/master"}, "development"},
		{branch{"release/branch-semantico/master"}, "development"},
		{branch{"master"}, "development"},
		{branch{"1.4"}, "production"},
	}

	for _, test := range tests {
		if got := test.branch.phase(); got != test.phase {
			t.Errorf("Oops! Expected " + test.phase + " but got " + test.branch.phase() + " using branch " + test.branch.name)
		}
	}
}

func TestExtractCurrentBranch(t *testing.T) {
	br := branch{"1.0"}
	if br.name != "1.0" {
		t.Errorf("Oops! Branch detection fails!")
	}
}
