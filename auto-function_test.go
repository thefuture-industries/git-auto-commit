package main

import "testing"

func TestParser_AddedGoFunction(t *testing.T) {
	GetDiff = func(file string) (string, error) {
		return "+func TestParser() {}", nil
	}

	DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := Parser([]string{"auto-commit-parser-test.go", "auto-commit-parser-test.go"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "added function TestParser"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}
