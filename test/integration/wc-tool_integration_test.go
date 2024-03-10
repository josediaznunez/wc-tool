package test

import (
	"os"
	"os/exec"
	"testing"
)

func TestWCToolByteCountFromFile(t *testing.T) {
	setupWCTool(t)

	file := "./../test.txt"

	cmd := exec.Command("./wc-tool", "-c", file)
	output, err := cmd.CombinedOutput()
	if err != nil {
		cleanupWCTool(t)
		t.Fatalf("Failed to execute wc-tool: %v", err)
	}

	expected := "\t342190" + " " + file + "\n"
	if string(output) != expected {
		t.Errorf("Expected output: %q, but got: %q", expected, output)
	}

	cleanupWCTool(t)
}

func TestWCToolLineCountFromFile(t *testing.T) {
	setupWCTool(t)

	file := "./../test.txt"

	cmd := exec.Command("./wc-tool", "-l", file)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to execute wc-tool: %v", err)
	}

	expected := "\t7145" + " " + file + "\n"
	if string(output) != expected {
		t.Errorf("Expected output: %q, but got: %q", expected, output)
	}

	cleanupWCTool(t)
}

func TestWCToolWordCountFromFile(t *testing.T) {
	setupWCTool(t)

	file := "./../test.txt"

	cmd := exec.Command("./wc-tool", "-w", file)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to execute wc-tool: %v", err)
	}

	expected := "\t58164" + " " + file + "\n"
	if string(output) != expected {
		t.Errorf("Expected output: %q, but got: %q", expected, output)
	}

	cleanupWCTool(t)
}

func TestWCToolCharacterCountFromFile(t *testing.T) {
	setupWCTool(t)

	file := "./../test.txt"

	cmd := exec.Command("./wc-tool", "-m", file)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to execute wc-tool: %v", err)
	}

	expected := "\t339292" + " " + file + "\n"
	if string(output) != expected {
		t.Errorf("Expected output: %q, but got: %q", expected, output)
	}

	cleanupWCTool(t)
}

func TestWCToolDefaultCountFromFile(t *testing.T) {
	setupWCTool(t)

	file := "./../test.txt"

	cmd := exec.Command("./wc-tool", file)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to execute wc-tool: %v", err)
	}

	expected := "\t7145 \t58164 \t342190" + " \t" + file + "\n"
	if string(output) != expected {
		t.Errorf("Expected output: %q, but got: %q", expected, output)
	}

	cleanupWCTool(t)
}

func TestWCToolLineCountFromStandardInput(t *testing.T) {
	setupWCTool(t)

	file := "./../test.txt"

	content := read(file, t)

	pipeReader, pipeWriter, err := os.Pipe()
	if err != nil {
		t.Fatalf("Error creating pipe: %v", err)
	}
	defer pipeReader.Close()
	defer pipeWriter.Close()

	write(pipeWriter, content, t)

	cmd := exec.Command("./wc-tool", "-l")
	cmd.Stdin = pipeReader
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to execute wc-tool: %v", err)
	}

	expected := "\t7145" + " " + "\n"
	if string(output) != expected {
		t.Errorf("Expected output: %q, but got: %q", expected, output)
	}

	cleanupWCTool(t)
}

func TestWCToolWordCountFromStandardInput(t *testing.T) {
	setupWCTool(t)

	file := "./../test.txt"

	content := read(file, t)

	pipeReader, pipeWriter, err := os.Pipe()
	if err != nil {
		t.Fatalf("Error creating pipe: %v", err)
	}
	defer pipeReader.Close()
	defer pipeWriter.Close()

	write(pipeWriter, content, t)

	cmd := exec.Command("./wc-tool", "-w")
	cmd.Stdin = pipeReader
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to execute wc-tool: %v", err)
	}

	expected := "\t58164" + " " + "\n"
	if string(output) != expected {
		t.Errorf("Expected output: %q, but got: %q", expected, output)
	}

	cleanupWCTool(t)
}

func TestWCToolByteCountFromStandardInput(t *testing.T) {
	setupWCTool(t)

	file := "./../test.txt"

	content := read(file, t)

	pipeReader, pipeWriter, err := os.Pipe()
	if err != nil {
		t.Fatalf("Error creating pipe: %v", err)
	}
	defer pipeReader.Close()
	defer pipeWriter.Close()

	write(pipeWriter, content, t)

	cmd := exec.Command("./wc-tool", "-c")
	cmd.Stdin = pipeReader
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to execute wc-tool: %v", err)
	}

	expected := "\t342190" + " " + "\n"
	if string(output) != expected {
		t.Errorf("Expected output: %q, but got: %q", expected, output)
	}

	cleanupWCTool(t)
}

func TestWCToolCharacterCountFromStandardInput(t *testing.T) {
	setupWCTool(t)

	file := "./../test.txt"

	content := read(file, t)

	pipeReader, pipeWriter, err := os.Pipe()
	if err != nil {
		t.Fatalf("Error creating pipe: %v", err)
	}
	defer pipeReader.Close()
	defer pipeWriter.Close()

	write(pipeWriter, content, t)

	cmd := exec.Command("./wc-tool", "-m")
	cmd.Stdin = pipeReader
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to execute wc-tool: %v", err)
	}

	expected := "\t339292" + " " + "\n"
	if string(output) != expected {
		t.Errorf("Expected output: %q, but got: %q", expected, output)
	}

	cleanupWCTool(t)
}

func TestWCToolDefaultCountFromStandardInput(t *testing.T) {
	setupWCTool(t)

	file := "./../test.txt"

	content := read(file, t)

	pipeReader, pipeWriter, err := os.Pipe()
	if err != nil {
		t.Fatalf("Error creating pipe: %v", err)
	}
	defer pipeReader.Close()
	defer pipeWriter.Close()

	write(pipeWriter, content, t)

	cmd := exec.Command("./wc-tool")
	cmd.Stdin = pipeReader
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to execute wc-tool: %v", err)
	}

	expected := "\t7145 \t58164 \t342190" + " \t\n"
	if string(output) != expected {
		t.Errorf("Expected output: %q, but got: %q", expected, output)
	}

	cleanupWCTool(t)
}

func setupWCTool(t *testing.T) {
	buildCmd := exec.Command("go", "build", "-o", "wc-tool", "./../../cmd/wc-tool/main.go")
	err := buildCmd.Run()
	if err != nil {
		t.Fatalf("Failed to build wc-tool: %v", err)
	}
}

func cleanupWCTool(t *testing.T) {
	err := os.Remove("wc-tool")
	if err != nil {
		t.Fatalf("Failed to remove wc-tool executable: %v", err)
	}
}

func read(file string, t *testing.T) []byte {
	content, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("Error reading test file: %v", err)
	}
	return content
}

func write(pipeWriter *os.File, content []byte, t *testing.T) {
	go func() {
		defer pipeWriter.Close()
		_, err := pipeWriter.Write(content)
		if err != nil {
			t.Errorf("Error writing to pipe: %v", err)
		}
	}()
}
