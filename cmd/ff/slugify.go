package main

import (
	"strings"
)

func slugify(sentence string) string {
	sentence = strings.ReplaceAll(sentence, " ", "-")
	sentence = strings.ReplaceAll(sentence, "'", "-")
	sentence = strings.ReplaceAll(sentence, "\n", "")
	return sentence
}
