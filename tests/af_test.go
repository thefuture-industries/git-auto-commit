package tests

import (
	"git-auto-commit/achelper/code"
	"git-auto-commit/diff"
	"git-auto-commit/parser"
	"testing"
)

func TestParser_AddedGoFunction(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	diff.GetDiff = func(file string) (string, error) {
		return "+func TestParser()", nil
	}

	code.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := parser.Parser([]string{"auto-commit-parser-test.go"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "added function TestParser"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestParser_DeletedGoFunction(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	diff.GetDiff = func(file string) (string, error) {
		return "-func TestParser()", nil
	}

	code.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := parser.Parser([]string{"auto-commit-parser-test.go"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "deleted function TestParser"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestParser_RenamedGoFunction(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	diff.GetDiff = func(file string) (string, error) {
		return "-func TestParser()", nil
	}

	code.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := parser.Parser([]string{"auto-commit-parser-test.go"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "deleted function TestParser"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}
