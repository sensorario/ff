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

func (b branch) isMaster() bool {
	return strings.HasPrefix(b.name, "master")
}

func (b branch) isRelease() bool {
	return strings.HasPrefix(b.name, "release/")
}

func (b branch) phase() string {
	if b.isRelease() ||
		b.isMaster() ||
		b.isFeature() ||
		b.isRefactoring() ||
		b.isBugfix() ||
		b.isHotfix() {
		return "development"
	}

	return "production"
}
