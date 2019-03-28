package main

import (
	"log"
	"os"
	"strconv"
	"strings"
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
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	logger := log.New(file, "prefix", log.LstdFlags)
	logger.Println("text to append")
	logger.Println("more text to append")

	patch := m.PatchVersion()
	minor := m.MinorVersion()
	major := m.MajorVersion()

	patchString, _ := strconv.Atoi(patch)
	minorString, _ := strconv.Atoi(minor)
	majorString, _ := strconv.Atoi(major)

	before := []string{
		strconv.Itoa(majorString),
		strconv.Itoa(minorString),
		strconv.Itoa(patchString),
	}

	log.Print(strings.Join([]string{"v", strings.Join(before, ".")}, ""))

	patchString++

	after := []string{
		strconv.Itoa(majorString),
		strconv.Itoa(minorString),
		strconv.Itoa(patchString),
	}

	log.Print(strings.Join([]string{"v", strings.Join(after, ".")}, ""))

	version := []string{"v", strings.Join(after, ".")}

	return strings.Join(version, "")
}

func (m *Meta) NextMinorTag() string {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	logger := log.New(file, "prefix", log.LstdFlags)
	logger.Println("text to append")
	logger.Println("more text to append")

	patch := m.PatchVersion()
	minor := m.MinorVersion()
	major := m.MajorVersion()

	patchString, _ := strconv.Atoi(patch)
	minorString, _ := strconv.Atoi(minor)
	majorString, _ := strconv.Atoi(major)

	before := []string{
		strconv.Itoa(majorString),
		strconv.Itoa(minorString),
		strconv.Itoa(patchString),
	}

	log.Print(strings.Join([]string{"v", strings.Join(before, ".")}, ""))

	minorString++

	after := []string{
		strconv.Itoa(majorString),
		strconv.Itoa(minorString),
		"0",
	}

	log.Print(strings.Join([]string{"v", strings.Join(after, ".")}, ""))

	version := []string{"v", strings.Join(after, ".")}

	return strings.Join(version, "")
}
