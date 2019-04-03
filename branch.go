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

func (b *Branch) IsMaster() bool {
	return strings.HasPrefix(b.branch, "master")
}
