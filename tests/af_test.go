package tests

import (
	"git-auto-commit/achelper/code"
	"git-auto-commit/diff"
	"git-auto-commit/parser"
	"testing"
)

func TestFormattedFunction_AddedGoFunction(t *testing.T) {
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

func TestFormattedFunction_DeletedGoFunction(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	diff.GetDiff = func(file string) (string, error) {
		return "-func TestParser() {}", nil
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

func TestFormattedFunction_ChangedParamNameGoFunction(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	diff.GetDiff = func(file string) (string, error) {
		return "-func ParamTest(a int)\n+func ParamTest(b int)", nil
	}

	code.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := parser.Parser([]string{"auto-commit-parser-test.go"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "renamed function TestParser -> TestParser"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}
