package main

import (
	"strings"
)

type Branch struct {
	branch string
}

func (b Branch) Destination() string {
	tokens := strings.Split(b.branch, "/")
	return tokens[len(tokens)-1]
}

func (b *Branch) IsRefactoring() bool {
	return strings.HasPrefix(b.branch, "refactor/")
}

func (b *Branch) IsFeature() bool {
	return strings.HasPrefix(b.branch, "feature/")
}

func (b *Branch) IsHotfix() bool {
	return strings.HasPrefix(b.branch, "hotfix/")
}

func (b *Branch) IsBugfix() bool {
	return strings.HasPrefix(b.branch, "bugfix/")
}

func (b *Branch) IsMaster() bool {
	return strings.HasPrefix(b.branch, "master")
}

func (b *Branch) IsRelease() bool {
	return strings.HasPrefix(b.branch, "release/")
}

func (b *Branch) Branch() string {
	return b.branch
}

func (b *Branch) Phase() string {
	if b.IsRelease() ||
		b.IsMaster() ||
		b.IsFeature() ||
		b.IsRefactoring() ||
		b.IsBugfix() ||
		b.IsHotfix() {
		return "development"
	}

	return "production"
}
