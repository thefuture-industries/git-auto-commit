package tests

import (
	"git-auto-commit/pkg"
	"git-auto-commit/pkg/git"
	"git-auto-commit/pkg/language"
	"testing"
)

func TestFormattedVariables_AddedGoVar(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	git.GetDiff = func(file string) (string, error) {
		return "+var testVar int = 5", nil
	}

	language.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := pkg.Parser([]string{"auto-commit-parser-test.go"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "added variable testVar"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedVariables_AddedGoVarEQ(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	git.GetDiff = func(file string) (string, error) {
		return "+testVar := 5", nil
	}

	language.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := pkg.Parser([]string{"auto-commit-parser-test.go"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "added variable testVar"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedVariables_AddedGoVars(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	git.GetDiff = func(file string) (string, error) {
		return "+var testVar1 int = 5\n+var testVar2 string\n+testVar3 := 0", nil
	}

	language.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := pkg.Parser([]string{"auto-commit-parser-test.go"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "added variables: testVar1, testVar2, testVar3"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedVariables_RenamedGoVar(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	git.GetDiff = func(file string) (string, error) {
		return "-var testVar1 int = 5\n+var testVar int = 5", nil
	}

	language.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := pkg.Parser([]string{"auto-commit-parser-test.go"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "renamed variable testVar1 -> testVar"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedVariables_ChangedTypeGoVar(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	git.GetDiff = func(file string) (string, error) {
		return "-var testVar int = 5\n+var testVar uint = 5", nil
	}

	language.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := pkg.Parser([]string{"auto-commit-parser-test.go"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "changed type of variable testVar (int -> uint)"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedVariables_ChangedTypeGoVars(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	git.GetDiff = func(file string) (string, error) {
		return "-var testVar uint8 = 5\n+var testVar uint16 = 5\n-var testVar2 uint16 = 5\n+var testVar2 int32 = 5", nil
	}

	language.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := pkg.Parser([]string{"auto-commit-parser-test.go"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "changed types of variables: testVar (uint8 -> uint16), testVar2 (uint16 -> int32)"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}
