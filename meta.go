package main

import (
	"strconv"
	"strings"
)

type Meta struct {
	describe string
	branch   string
}

func (m Meta) Branch() string {
	return m.branch
}

func (m Meta) majorVersion() string {
	tokens := strings.Split(m.describe, "-")
	tokens = strings.Split(tokens[0], ".")
	tokens = strings.Split(tokens[0], "v")
	return tokens[1]
}

func (m Meta) minorVersion() string {
	tokens := strings.Split(m.describe, "-")
	tokens = strings.Split(tokens[0], ".")
	return tokens[1]
}

func (m Meta) IncpatchVersion() int {
	tokens := strings.Split(m.describe, "-")
	tokens = strings.Split(tokens[0], ".")
	number, _ := strconv.Atoi(tokens[2])
	nextNumber := number + 1
	return nextNumber
}

func (m Meta) patchVersion() string {
	tokens := strings.Split(m.describe, "-")
	tokens = strings.Split(tokens[0], ".")
	return tokens[2]
}

func (m Meta) patchVersionInt() string {
	tokens := strings.Split(m.describe, "-")
	tokens = strings.Split(tokens[0], ".")
	number, _ := strconv.Atoi(tokens[1])
	return string(number)
}

func (m Meta) CommitsFromLastTag() string {
	tokens := strings.Split(m.describe, "-")
	tokens = strings.Split(tokens[1], ".")
	return tokens[0]
}

func (m Meta) NextPatchTag() string {
	patch := m.patchVersion()
	minor := m.minorVersion()
	major := m.majorVersion()

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

func (m Meta) NextMinorTag() string {
	minor := m.minorVersion()
	major := m.majorVersion()

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
