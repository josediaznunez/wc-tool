package test

import (
	"testing"

	"wc-tool/pkg/commands"
)

var input string

func setup() {
	input = "Hello\nWorld\n"
}

func TestByteCount(t *testing.T) {
	setup()
	expectedBytes := 12
	bytes := commands.ByteCount(input)
	if bytes != expectedBytes {
		t.Errorf("Expected %d bytes, got %d", expectedBytes, bytes)
	}
}

func TestLineCount(t *testing.T) {
	setup()
	expectedLines := 2
	lines := commands.LineCount(input)
	if lines != expectedLines {
		t.Errorf("Expected %d lines, got %d", expectedLines, lines)
	}
}

func TestWordCount(t *testing.T) {
	setup()
	expectedWords := 2
	words := commands.WordCount(input)
	if words != expectedWords {
		t.Errorf("Expected %d words, got %d", expectedWords, words)
	}
}

func TestCharacterCount(t *testing.T) {
	setup()
	expectedCharacters := 12
	characters := commands.CharacterCount(input)
	if characters != expectedCharacters {
		t.Errorf("Expected %d words, got %d", expectedCharacters, characters)
	}
}
