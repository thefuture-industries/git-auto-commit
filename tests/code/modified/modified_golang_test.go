package golang

import (
	"fmt"
	"git-auto-commit/pkg/code"
	"git-auto-commit/tests"
	"os/exec"
	"strings"
	"testing"
)

const modifieExpectedTest string = "[fix] Updated command-line tools"

func TestModified(t *testing.T) {
	gitOutput := `
		M	cmd/main.mdf
	`

	code.ExecCommand = func(name string, args ...string) *exec.Cmd {
		return tests.FakeExecCommand(gitOutput)
	}
	defer func() {
		code.ExecCommand = exec.Command
	}()

	c := &code.Code{}

	files := []string{"cmd/main.mdf"}
	msg, err := c.FormattedCode(files)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !strings.EqualFold(strings.TrimSpace(msg), strings.TrimSpace(modifieExpectedTest)) {
		t.Errorf("Expected commit message:\n%q\nGot:%q", modifieExpectedTest, msg)
	}

	fmt.Println("==> Formatted commit message:", msg)
}
