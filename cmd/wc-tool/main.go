package main

import (
	"flag"
	"fmt"
	"os"

	"wc-tool/pkg/commands"
)

func main() {
	countLines := flag.Bool("l", false, "Count lines")
	countWords := flag.Bool("w", false, "Count words")
	countBytes := flag.Bool("c", false, "Count bytes")
	countCharacters := flag.Bool("m", false, "Count characters")

	flag.Parse()

	var input string
	var filePath string
	if flag.NArg() > 0 {
		for _, file := range flag.Args() {
			filePath = file
			content, err := os.ReadFile(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", file, err)
				print(os.Getwd())
				os.Exit(1)
			}
			input += string(content)
		}
	}

	if *countLines {
		lines := commands.LineCount(input)
		fmt.Printf("%d %s\n", lines, filePath)
	} else if *countWords {
		words := commands.WordCount(input)
		fmt.Printf("%d %s\n", words, filePath)
	} else if *countBytes {
		bytes := commands.ByteCount(input)
		fmt.Printf("%d %s\n", bytes, filePath)
	} else if *countCharacters {
		characters := commands.CharacterCount(input)
		fmt.Printf("%d %s\n", characters, filePath)
	} else {
		lines := commands.LineCount(input)
		bytes := commands.ByteCount(input)
		words := commands.WordCount(input)
		fmt.Printf("%d %d %d %s\n", lines, words, bytes, filePath)
	}
}
