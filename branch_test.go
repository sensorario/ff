package main

import (
	"testing"
)

func TestAnalyze(t *testing.T) {
	branch := Branch{"feature/branch-semantico/master"}

	if branch.IsFeature() == false {
		t.Errorf("branch should be recognyzed as feature")
	}

	if branch.Destination() != "master" {
		t.Errorf("destination branch shoudl be master")
	}
}

func TestMasterOrFeaturesBranchesAreDevelopment(t *testing.T) {
	tests := []struct {
		branch Branch
		phase  string
	}{
		{Branch{"feature/branch-semantico/master"}, "development"},
		{Branch{"hotfix/branch-semantico/master"}, "development"},
		{Branch{"release/branch-semantico/master"}, "development"},
		{Branch{"master"}, "development"},
		{Branch{"1.4"}, "production"},
	}

	for _, test := range tests {
		if got := test.branch.Phase(); got != test.phase {
			t.Errorf("Oops! Expected " + test.phase + " but got " + test.branch.Phase() + " using branch " + test.branch.branch)
		}
	}
}
