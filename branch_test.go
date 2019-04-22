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

func TestRefactoringBranchStartsWithRefactor(t *testing.T) {
	br := branch{"refactor/description/develop"}
	if br.isRefactoring() == false {
		t.Errorf("branch should be rafactorng but is not")
	}
}

func TestHotfixBranchStartsWithHotfix(t *testing.T) {
	br := branch{"hotfix/description/develop"}
	if br.isHotfix() == false {
		t.Errorf("branch should be hotfix but is not")
	}
}

func TestBugfixBranchStartsWithBugfix(t *testing.T) {
	br := branch{"bugfix/description/develop"}
	if br.isBugfix() == false {
		t.Errorf("branch should be bugfix but is not")
	}
}

func TestDevelopmentBranchContains(t *testing.T) {
	br := branch{"master"}
	if br.isDevelopment("master") == false {
		t.Errorf("branch should be development but is not")
	}
}

func TestReleaseBranchStartsWithRelease(t *testing.T) {
	br := branch{"release/foo/bar"}
	if br.isRelease() == false {
		t.Errorf("branch should be release but is not")
	}
}
