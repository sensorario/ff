package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	meta := meta{"v21.2.54-42-b67b5vr", "master"}

	if meta.Branch() != "master" {
		t.Errorf("bad branch detection")
	}

	if meta.majorVersion() != "21" {
		t.Errorf("bad branch detection" + meta.majorVersion())
	}

	if meta.minorVersion() != "2" {
		t.Errorf("bad branch detection" + meta.minorVersion())
	}

	if meta.patchVersion() != "54" {
		t.Errorf("bad branch detection" + meta.patchVersion())
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

	num := meta.incPatchVersion()
	assert.Equal(t, num, 55, "Bad branch detection // "+meta.NextMinorTag())
}
