package commands

import (
	"strings"
)

func ByteCount(input string) int {
	return len([]byte(input))
}

func LineCount(input string) int {
	return strings.Count(input, "\n")
}

func WordCount(input string) int {
	return len(strings.Fields(input))
}

func CharacterCount(input string) int {
	return len([]rune(input))
}
