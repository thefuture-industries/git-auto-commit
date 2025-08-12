package golang

import (
	"fmt"
	"git-auto-commit/pkg/code"
	"git-auto-commit/tests"
	"os/exec"
	"strings"
	"testing"
)

const implementeExpectedTest string = "Implemented source code files"

func TestImplementedGolang(t *testing.T) {
	gitOutput := `
		A	src/main.c
	`

	code.ExecCommand = func(name string, args ...string) *exec.Cmd {
		return tests.FakeExecCommand(gitOutput)
	}
	defer func() {
		code.ExecCommand = exec.Command
	}()

	c := &code.Code{}

	files := []string{"src/main.c"}
	msg, err := c.FormattedCode(files)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !strings.EqualFold(strings.TrimSpace(msg), strings.TrimSpace(implementeExpectedTest)) {
		t.Errorf("Expected commit message:\n%q\nGot:\n%q", implementeExpectedTest, msg)
	}

	fmt.Println("Formatted commit message:", msg)
}
