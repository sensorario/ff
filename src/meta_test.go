package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	m := meta{"v21.2.54-42-b67b5vr", "master"}

	if m.Branch() != "master" {
		t.Errorf("bad branch detection")
	}

	if m.majorVersion() != "21" {
		t.Errorf("bad branch detection" + m.majorVersion())
	}

	if m.minorVersion() != "2" {
		t.Errorf("bad branch detection" + m.minorVersion())
	}

	if m.patchVersion() != "54" {
		t.Errorf("bad branch detection" + m.patchVersion())
	}

	if m.CommitsFromLastTag() != "42" {
		t.Errorf("bad number detection" + m.CommitsFromLastTag())
	}

	if m.NextPatchTag() != "v21.2.55" {
		t.Errorf("Bad branch detection // " + m.NextPatchTag())
	}

	num := m.incPatchVersion()
	assert.Equal(t, num, 55, "Bad branch detection // "+m.NextMinorTag())
}

func TestMajorVersionDetection(t *testing.T) {
    type test struct {
        m meta
        major string
        minor string
        patch string
    }

    tests := []test{
	    {meta{"1.2.3", "master"}, "1", "2", "3"},
	    {meta{"v1.2.3", "master"}, "1", "2", "3"},
    }

    for _, tc := range tests {
        if tc.m.majorVersion() != tc.major {
            t.Errorf( tc.m.majorVersion(), " should be equal to ", tc.major)
        }
        if tc.m.minorVersion() != tc.minor {
            t.Errorf( tc.m.minorVersion(), " should be equal to ", tc.minor)
        }
        if tc.m.patchVersion() != tc.patch {
            t.Errorf( tc.m.patchVersion(), " should be equal to ", tc.patch)
        }
    }

}

func TestMinorTagDetection(t *testing.T) {
	m := meta{"v1.2.3", "master"}
    expected := "v1.3.0"

	if m.NextMinorTag() != expected {
        t.Errorf(
            m.NextMinorTag(),
            " instead of ",
            expected,
        )
    }
}
