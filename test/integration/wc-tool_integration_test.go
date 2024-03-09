package test

import (
	"os"
	"os/exec"
	"testing"
)

func setup(t *testing.T) {
	buildCmd := exec.Command("go", "build", "-o", "wc-tool", "./../../cmd/wc-tool/main.go")
	err := buildCmd.Run()
	if err != nil {
		t.Fatalf("Failed to build wc-tool: %v", err)
	}
}

func cleanup(t *testing.T) {
	err := os.Remove("wc-tool")
	if err != nil {
		t.Fatalf("Failed to remove wc-tool executable: %v", err)
	}
}

func TestWCToolByteCount(t *testing.T) {
	setup(t)

	testFilePath := "./../test.txt"

	cmd := exec.Command("./wc-tool", "-c", testFilePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		cleanup(t)
		t.Fatalf("Failed to execute wc-tool: %v", err)
	}

	expected := "342190" + " " + testFilePath + "\n"
	if string(output) != expected {
		t.Errorf("Expected output: %q, but got: %q", expected, output)
	}

	cleanup(t)
}

func TestWCToolLineCount(t *testing.T) {
	setup(t)

	testFilePath := "./../test.txt"

	cmd := exec.Command("./wc-tool", "-l", testFilePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to execute wc-tool: %v", err)
	}

	expected := "7145" + " " + testFilePath + "\n"
	if string(output) != expected {
		t.Errorf("Expected output: %q, but got: %q", expected, output)
	}

	cleanup(t)
}

func TestWCToolWordCount(t *testing.T) {
	setup(t)

	testFilePath := "./../test.txt"

	cmd := exec.Command("./wc-tool", "-w", testFilePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to execute wc-tool: %v", err)
	}

	expected := "58164" + " " + testFilePath + "\n"
	if string(output) != expected {
		t.Errorf("Expected output: %q, but got: %q", expected, output)
	}

	cleanup(t)
}

func TestWCToolCharacterCount(t *testing.T) {
	setup(t)

	testFilePath := "./../test.txt"

	cmd := exec.Command("./wc-tool", "-m", testFilePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to execute wc-tool: %v", err)
	}

	expected := "339292" + " " + testFilePath + "\n"
	if string(output) != expected {
		t.Errorf("Expected output: %q, but got: %q", expected, output)
	}

	cleanup(t)
}
