package tests

import (
	"git-auto-commit/achelper/code"
	"git-auto-commit/diff"
	"git-auto-commit/parser"
	"testing"
)

func TestFormattedImport_AddedGoImport(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	diff.GetDiff = func(file string) (string, error) {
		return "+import \"fmt\"", nil
	}

	code.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := parser.Parser([]string{"main.go"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "included 'fmt' in main.go"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedImport_AddedGoImports(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	diff.GetDiff = func(file string) (string, error) {
		return "+import \"fmt\"\n+import \"os\"", nil
	}

	code.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := parser.Parser([]string{"main.go"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "included 'fmt', 'os' in main.go"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedImport_AddedImportsBlockGo(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	diff.GetDiff = func(file string) (string, error) {
		return "+import (\n+\"fmt\"\n+\"os\"\n+)", nil
	}

	code.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := parser.Parser([]string{"main.go"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "included 'fmt', 'os' in main.go"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedImport_AddedImportBlockGo(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	diff.GetDiff = func(file string) (string, error) {
		return "+import (\n+\"fmt\"\n+)", nil
	}

	code.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := parser.Parser([]string{"main.go"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "included 'fmt' in main.go"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}
