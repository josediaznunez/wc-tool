package test

import (
	"testing"

	"wc-tool/pkg/commands"
)

var sampleText string

func setup() {
	sampleText = "Hello\nWorld\n"
}

func TestByteCount(t *testing.T) {
	setup()

	expectedBytes := 12

	bytes := commands.ByteCount(sampleText)

	if bytes != expectedBytes {
		t.Errorf("Expected %d bytes, got %d", expectedBytes, bytes)
	}
}

func TestLineCount(t *testing.T) {
	setup()

	expectedLines := 2

	bytes := commands.LineCount(sampleText)

	if bytes != expectedLines {
		t.Errorf("Expected %d lines, got %d", expectedLines, bytes)
	}
}

func TestWordCount(t *testing.T) {
	setup()

	expectedWords := 2

	bytes := commands.WordCount(sampleText)

	if bytes != expectedWords {
		t.Errorf("Expected %d words, got %d", expectedWords, bytes)
	}
}

func TestCharacterCount(t *testing.T) {
	setup()

	expectedCharacters := 12

	characters := commands.CharacterCount(sampleText)

	if characters != expectedCharacters {
		t.Errorf("Expected %d words, got %d", expectedCharacters, characters)
	}
}
