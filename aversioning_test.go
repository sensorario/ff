package main

import (
	"strconv"
	"strings"
	"testing"
)

type Meta struct {
	describe string
	branch   string
}

func (m *Meta) Branch() string {
	return m.branch
}

func (m *Meta) MajorVersion() string {
	tokens := strings.Split(m.describe, "-")
	tokens = strings.Split(tokens[0], ".")
	tokens = strings.Split(tokens[0], "v")
	return tokens[1]
}

func (m *Meta) MinorVersion() string {
	tokens := strings.Split(m.describe, "-")
	tokens = strings.Split(tokens[0], ".")
	return tokens[1]
}

func (m *Meta) PatchVersion() string {
	tokens := strings.Split(m.describe, "-")
	tokens = strings.Split(tokens[0], ".")
	return tokens[2]
}

func (m *Meta) PatchVersionInt() string {
	tokens := strings.Split(m.describe, "-")
	tokens = strings.Split(tokens[0], ".")
	number, _ := strconv.Atoi(tokens[1])
	return string(number)
}

func (m *Meta) CommitsFromLastTag() string {
	tokens := strings.Split(m.describe, "-")
	tokens = strings.Split(tokens[1], ".")
	return tokens[0]
}

func (m *Meta) NextPatchTag() string {
	patch := m.PatchVersion()
	minor := m.MinorVersion()
	major := m.MajorVersion()

	patchString, _ := strconv.Atoi(patch)
	minorString, _ := strconv.Atoi(minor)
	majorString, _ := strconv.Atoi(major)

	patchString++

	foo := []string{
		strconv.Itoa(majorString),
		strconv.Itoa(minorString),
		strconv.Itoa(patchString),
	}

	version := []string{"v", strings.Join(foo, ".")}
	return strings.Join(version, "")
}

func (m *Meta) NextMinorTag() string {
	minor := m.MinorVersion()
	major := m.MajorVersion()

	minorString, _ := strconv.Atoi(minor)
	majorString, _ := strconv.Atoi(major)

	minorString++

	foo := []string{
		strconv.Itoa(majorString),
		strconv.Itoa(minorString),
		"0",
	}

	version := []string{"v", strings.Join(foo, ".")}
	return strings.Join(version, "")
}

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
