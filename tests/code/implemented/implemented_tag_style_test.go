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

func TestImplementedTagStyle(t *testing.T) {
	gitOutput := `
		A	styles/components.css
	`

	code.ExecCommand = func(name string, args ...string) *exec.Cmd {
		return tests.FakeExecCommand(gitOutput)
	}
	defer func() {
		code.ExecCommand = exec.Command
	}()

	c := &code.Code{}

	files := []string{"styles/components.css"}
	msg, err := c.FormattedCode(files)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !strings.Contains(msg, constants.Type_CommitStyle) {
		t.Errorf("Expected commit message including:%q Got: %q", constants.Type_CommitStyle, msg)
	}

	fmt.Println("==> Formatted commit message:", msg)
}
