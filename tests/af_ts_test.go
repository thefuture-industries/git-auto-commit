package tests

import (
	"git-auto-commit/achelper/code"
	"git-auto-commit/diff"
	"git-auto-commit/parser"
	"testing"
)

func TestFormattedFunction_AddedTSFunction(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	diff.GetDiff = func(file string) (string, error) {
		return "+function AddedTSFunction() {}", nil
	}

	code.DetectLanguage = func(filename string) string {
		return "typescript"
	}

	msg, err := parser.Parser([]string{"auto-commit-parser-test.ts"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "added function AddedTSFunction"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedFunction_AddedTSFunctions(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	diff.GetDiff = func(file string) (string, error) {
		return "+function AddedTSFunction1() {}\n+function AddedTSFunction2() {}\n+function AddedTSFunction3() {}", nil
	}

	code.DetectLanguage = func(filename string) string {
		return "typescript"
	}

	msg, err := parser.Parser([]string{"auto-commit-parser-test.ts"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "added functions: AddedTSFunction1, AddedTSFunction2, AddedTSFunction3"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedFunction_DeletedTSFunction(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	diff.GetDiff = func(file string) (string, error) {
		return "-function DeletedTSFunction() {}", nil
	}

	code.DetectLanguage = func(filename string) string {
		return "typescript"
	}

	msg, err := parser.Parser([]string{"auto-commit-parser-test.ts"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "deleted function DeletedTSFunction"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedFunction_DeletedTSFunctions(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	diff.GetDiff = func(file string) (string, error) {
		return "-function DeletedTSFunction1() {}\n-function DeletedTSFunction2() {}\n-function DeletedTSFunction3() {}", nil
	}

	code.DetectLanguage = func(filename string) string {
		return "typescript"
	}

	msg, err := parser.Parser([]string{"auto-commit-parser-test.ts"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "deleted functions: DeletedTSFunction1, DeletedTSFunction2, DeletedTSFunction3"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedFunction_ChangedParamNameTSFunction(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	diff.GetDiff = func(file string) (string, error) {
		return "-function ParamTest(a: number) {}\n+function ParamTest(b: number) {}", nil
	}

	code.DetectLanguage = func(filename string) string {
		return "typescript"
	}

	msg, err := parser.Parser([]string{"auto-commit-parser-test.ts"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "changed parameter in ParamTest function"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedFunction_ChangedParamNameTSFunctions(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	diff.GetDiff = func(file string) (string, error) {
		return "-function ParamTest1(a: number) {}\n+function ParamTest1(b: number) {}\n-function ParamTest2(x: string) {}\n+function ParamTest2(y: string) {}", nil
	}

	code.DetectLanguage = func(filename string) string {
		return "typescript"
	}

	msg, err := parser.Parser([]string{"auto-commit-parser-test.ts"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "changed parameters in functions: ParamTest1 function, ParamTest2 function"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedFunction_ChangedParamTypeTSFunction(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	diff.GetDiff = func(file string) (string, error) {
		return "-function TypeTest(a: number) {}\n+function TypeTest(a: string) {}", nil
	}

	code.DetectLanguage = func(filename string) string {
		return "typescript"
	}

	msg, err := parser.Parser([]string{"auto-commit-parser-test.ts"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "changed parameter type a in TypeTest function"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}
