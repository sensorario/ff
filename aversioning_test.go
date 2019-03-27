package main

import (
	"testing"
)

func TestAbs(t *testing.T) {
	meta := Meta{"v21.2.54-42-b67b5vr", "master"}

	if meta.Branch() != "master" {
		t.Errorf("bad branch detection")
	}

	if meta.MajorVersion() != "21" {
		t.Errorf("bad branch detection" + meta.MajorVersion())
	}

	if meta.MinorVersion() != "2" {
		t.Errorf("bad branch detection" + meta.MinorVersion())
	}

	if meta.PatchVersion() != "54" {
		t.Errorf("bad branch detection" + meta.PatchVersion())
	}

	if meta.CommitsFromLastTag() != "42" {
		t.Errorf("bad number detection" + meta.CommitsFromLastTag())
	}

	if meta.NextPatchTag() != "v21.2.55" {
		t.Errorf("Bad branch detection // " + meta.NextPatchTag())
	}

	if meta.NextMinorTag() != "v21.3.0" {
		t.Errorf("Bad branch detection // " + meta.NextMinorTag())
	}
}
