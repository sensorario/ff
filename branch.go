package main

import (
	"strings"
)

type branch struct {
	name string
}

func (b branch) destination() string {
	tokens := strings.Split(b.name, "/")
	return tokens[len(tokens)-1]
}

func (b branch) isRefactoring() bool {
	return strings.HasPrefix(b.name, "refactor/")
}

func (b branch) isFeature() bool {
	return strings.HasPrefix(b.name, "feature/")
}

func (b branch) isHotfix() bool {
	return strings.HasPrefix(b.name, "hotfix/")
}

func (b branch) isBugfix() bool {
	return strings.HasPrefix(b.name, "bugfix/")
}

func (b branch) isDevelopment(devBranchName string) bool {
	return strings.HasPrefix(b.name, devBranchName)
}

func (b branch) isRelease() bool {
	return strings.HasPrefix(b.name, "release/")
}
