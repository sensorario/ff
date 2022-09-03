package main

import (
	"strconv"
	"strings"
	"fmt"
    // "os"
)

type meta struct {
	describe string
	branch   string
}

func (m meta) Branch() string {
	return m.branch
}

func (m meta) majorVersion() string {
	tokens := strings.Split(m.describe, "-")
	tokens = strings.Split(tokens[0], ".")
	tokens = strings.Split(tokens[0], "v")
    if len(tokens) == 1 {
        return tokens[0]
    }

	return tokens[1]
}

func (m meta) minorVersion() string {
	tokens := strings.Split(m.describe, "-")
	tokens = strings.Split(tokens[0], ".")
	return tokens[1]
}

func (m meta) incPatchVersion() int {
	tokens := strings.Split(m.describe, "-")
	tokens = strings.Split(tokens[0], ".")
	number, _ := strconv.Atoi(tokens[2])
	nextNumber := number + 1
	return nextNumber
}

func (m meta) patchVersion() string {
	tokens := strings.Split(m.describe, "-")
	tokens = strings.Split(tokens[0], ".")

    // fmt.Println("Tokens: ")
    // fmt.Println(tokens)
    // os.Exit(1)

	return tokens[2]
}

func (m meta) patchVersionInt() string {
	tokens := strings.Split(m.describe, "-")
	tokens = strings.Split(tokens[0], ".")
	number, _ := strconv.Atoi(tokens[1])
	return fmt.Sprint(number)
}

func (m meta) CommitsFromLastTag() string {
	tokens := strings.Split(m.describe, "-")
	tokens = strings.Split(tokens[1], ".")
	return tokens[0]
}

func (m meta) NextPatchTag() string {
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

func (m meta) NextMinorTag() string {
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
