package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"wc-tool/pkg/commands"
)

func main() {
	var input string
	var file string
	countLines := flag.Bool("l", false, "Count lines")
	countWords := flag.Bool("w", false, "Count words")
	countBytes := flag.Bool("c", false, "Count bytes")
	countCharacters := flag.Bool("m", false, "Count characters")

	flag.Parse()

	if flag.NArg() > 0 {
		file = flag.Arg(0)
		input = readFile(file)
	} else {
		input = readStandardInput()
	}

	if *countLines {
		lines := commands.LineCount(input)
		fmt.Printf("\t%d %s\n", lines, file)
	} else if *countWords {
		words := commands.WordCount(input)
		fmt.Printf("\t%d %s\n", words, file)
	} else if *countBytes {
		bytes := commands.ByteCount(input)
		fmt.Printf("\t%d %s\n", bytes, file)
	} else if *countCharacters {
		characters := commands.CharacterCount(input)
		fmt.Printf("\t%d %s\n", characters, file)
	} else {
		lines := commands.LineCount(input)
		words := commands.WordCount(input)
		bytes := commands.ByteCount(input)
		fmt.Printf("\t%d \t%d \t%d \t%s\n", lines, words, bytes, file)
	}
}

func readFile(file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from file %s: %v\n", file, err)
		os.Exit(1)
	}
	return string(content)
}

func readStandardInput() string {
	content, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from standard input: %v\n", err)
		os.Exit(1)
	}
	return string(content)
}
