package main

import (
	"strings"
)

type SemBranch struct {
	branch string
}

func (sb *SemBranch) Destination() string {
	tokens := strings.Split(sb.branch, "/")
	return tokens[len(tokens)-1]
}

func (sb *SemBranch) IsFeature() bool {
	return strings.HasPrefix(sb.branch, "feature/")
}
