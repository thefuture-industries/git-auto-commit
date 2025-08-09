package tests

import (
	"git-auto-commit/pkg"
	"git-auto-commit/pkg/git"
	"git-auto-commit/pkg/language"
	"testing"
)

func TestFormattedFunction_AddedGoFunction(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	git.GetDiff = func(file string) (string, error) {
		return "+func AddedGoFunction() {}", nil
	}

	language.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := pkg.Parser([]string{"auto-commit-parser-test.go"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "added function AddedGoFunction"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedFunction_AddedGoFunctions(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	git.GetDiff = func(file string) (string, error) {
		return "+func AddedGoFunction1()\n+func AddedGoFunction2()\n+func AddedGoFunction3()", nil
	}

	language.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := pkg.Parser([]string{"auto-commit-parser-test.go"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "added functions: AddedGoFunction1, AddedGoFunction2, AddedGoFunction3"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedFunction_DeletedGoFunction(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	git.GetDiff = func(file string) (string, error) {
		return "-func DeletedGoFunction() {}", nil
	}

	language.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := pkg.Parser([]string{"auto-commit-parser-test.go"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "deleted function DeletedGoFunction"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedFunction_DeletedGoFunctions(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	git.GetDiff = func(file string) (string, error) {
		return "-func DeletedGoFunction1()\n-func DeletedGoFunction2()\n-func DeletedGoFunction3()", nil
	}

	language.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := pkg.Parser([]string{"auto-commit-parser-test.go"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "deleted functions: DeletedGoFunction1, DeletedGoFunction2, DeletedGoFunction3"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedFunction_ChangedParamNameGoFunction(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	git.GetDiff = func(file string) (string, error) {
		return "-func ParamTest(a int)\n+func ParamTest(b int)", nil
	}

	language.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := pkg.Parser([]string{"auto-commit-parser-test.go"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "changed parameter in ParamTest function"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedFunction_ChangedParamNameGoFunctions(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	git.GetDiff = func(file string) (string, error) {
		return "-func ParamTest1(a int)\n+func ParamTest1(b int)\n-func ParamTest2(x string)\n+func ParamTest2(y string)", nil
	}

	language.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := pkg.Parser([]string{"auto-commit-parser-test.go"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "changed parameters in functions: ParamTest1 function, ParamTest2 function"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}

func TestFormattedFunction_ChangedParamTypeGoFunction(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	git.GetDiff = func(file string) (string, error) {
		return "-func TypeTest(a int)\n+func TypeTest(a string)", nil
	}

	language.DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := pkg.Parser([]string{"auto-commit-parser-test.go"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "changed parameter type a in TypeTest function"
	if msg != expected {
		t.Errorf("expected '%s', got '%s'", expected, msg)
	}
}
