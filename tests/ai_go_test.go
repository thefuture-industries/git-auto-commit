package tests

import (
	"git-auto-commit/pkg"
	"git-auto-commit/pkg/git"
	"git-auto-commit/pkg/language"
	"testing"
)

func TestFormattedImport_AddedGoImport(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	git.GetDiff = func(file string) (string, error) {
		return "+import \"fmt\"", nil
	}

	language.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := pkg.Parser([]string{"main.go"})
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

	git.GetDiff = func(file string) (string, error) {
		return "+import \"fmt\"\n+import \"os\"", nil
	}

	language.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := pkg.Parser([]string{"main.go"})
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

	git.GetDiff = func(file string) (string, error) {
		return "+import (\n+\"fmt\"\n+\"os\"\n+)", nil
	}

	language.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := pkg.Parser([]string{"main.go"})
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

	git.GetDiff = func(file string) (string, error) {
		return "+import (\n+\"fmt\"\n+)", nil
	}

	language.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := pkg.Parser([]string{"main.go"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "included 'fmt' in main.go"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}
