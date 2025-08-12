package golang

import (
	"fmt"
	"git-auto-commit/infra/constants"
	"git-auto-commit/pkg/code"
	"git-auto-commit/tests"
	"os/exec"
	"strings"
	"testing"
)

func TestImplementedTagTest(t *testing.T) {
	gitOutput := `
		A	tests/golang_test.impl
	`

	code.ExecCommand = func(name string, args ...string) *exec.Cmd {
		return tests.FakeExecCommand(gitOutput)
	}
	defer func() {
		code.ExecCommand = exec.Command
	}()

	c := &code.Code{}

	files := []string{"tests/golang_test.impl"}
	msg, err := c.FormattedCode(files)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !strings.Contains(msg, constants.Type_CommitTest) {
		t.Errorf("Expected commit message including:%q Got: %q", constants.Type_CommitTest, msg)
	}

	fmt.Println("==> Formatted commit message:", msg)
}
