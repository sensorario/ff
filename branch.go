package main

import (
	"strings"
)

type Branch struct {
	name string
}

func (b Branch) Destination() string {
	tokens := strings.Split(b.name, "/")
	return tokens[len(tokens)-1]
}

func (b *Branch) isRefactoring() bool {
	return strings.HasPrefix(b.name, "refactor/")
}

func (b *Branch) isFeature() bool {
	return strings.HasPrefix(b.name, "feature/")
}

func (b *Branch) isHotfix() bool {
	return strings.HasPrefix(b.name, "hotfix/")
}

func (b *Branch) isBugfix() bool {
	return strings.HasPrefix(b.name, "bugfix/")
}

func (b *Branch) isMaster() bool {
	return strings.HasPrefix(b.name, "master")
}

func (b *Branch) isRelease() bool {
	return strings.HasPrefix(b.name, "release/")
}

func (b *Branch) Phase() string {
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
