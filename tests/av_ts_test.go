package tests

import (
	"git-auto-commit/achelper/code"
	"git-auto-commit/diff"
	"git-auto-commit/parser"
	"testing"
)

func TestFormattedVariables_AddedTSVar(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	diff.GetDiff = func(file string) (string, error) {
		return "+let testVar: number = 5;", nil
	}

	code.DetectLanguage = func(filename string) string {
		return "typescript"
	}

	msg, err := parser.Parser([]string{"auto-commit-parser-test.ts"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "added variable testVar"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedVariables_AddedTSVars(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	diff.GetDiff = func(file string) (string, error) {
		return "+let testVar1: number = 5;\n+const testVar2: string = 'hi';", nil
	}

	code.DetectLanguage = func(filename string) string {
		return "typescript"
	}

	msg, err := parser.Parser([]string{"auto-commit-parser-test.ts"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "added variables: testVar1, testVar2"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedVariables_RenamedTSVar(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	diff.GetDiff = func(file string) (string, error) {
		return "-let testVar1: number = 5;\n+let testVar: number = 5;", nil
	}

	code.DetectLanguage = func(filename string) string {
		return "typescript"
	}

	msg, err := parser.Parser([]string{"auto-commit-parser-test.ts"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "renamed variable testVar1 -> testVar"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedVariables_ChangedTypeTSVar(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	diff.GetDiff = func(file string) (string, error) {
		return "-let testVar: number = 5;\n+let testVar: string = 5;", nil
	}

	code.DetectLanguage = func(filename string) string {
		return "typescript"
	}

	msg, err := parser.Parser([]string{"auto-commit-parser-test.ts"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "changed type of variable testVar (number -> string)"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedVariables_ChangedTypesTSVars(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	diff.GetDiff = func(file string) (string, error) {
		return "-let testVar: number = 5;\n+let testVar: string = 5;\n-var testVar2: boolean = true;\n+var testVar2: number = true;", nil
	}

	code.DetectLanguage = func(filename string) string {
		return "typescript"
	}

	msg, err := parser.Parser([]string{"auto-commit-parser-test.ts"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "changed types of variables: testVar (number -> string), testVar2 (boolean -> number)"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}
