package implmoddel

import (
	"fmt"
	"git-auto-commit/pkg/code"
	"git-auto-commit/tests"
	"os/exec"
	"strings"
	"testing"
)

const implementeModifieDeleteExpectedTest string = "[fix] Implemented source code files, updated public reusable packages, removed error"

func TestImplementedModifiedDeleted(t *testing.T) {
	gitOutput := `
		M	pkg/modified.mod
		A	src/implemented.impl
		D	error/deleted.del
	`

	code.ExecCommand = func(name string, args ...string) *exec.Cmd {
		return tests.FakeExecCommand(gitOutput)
	}
	defer func() {
		code.ExecCommand = exec.Command
	}()

	c := &code.Code{}

	files := []string{"src/implemented.impl", "pkg/modified.mod", "error/deleted.del"}
	msg, err := c.FormattedCode(files)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !strings.EqualFold(strings.TrimSpace(msg), strings.TrimSpace(implementeModifieDeleteExpectedTest)) {
		t.Errorf("Expected commit message:\n%q\nGot:\n%q", implementeModifieDeleteExpectedTest, msg)
	}

	fmt.Println("==> Formatted commit message:", msg)
}
