package golang

import (
	"fmt"
	"git-auto-commit/pkg/code"
	"git-auto-commit/tests"
	"os/exec"
	"strings"
	"testing"
)

const deleteExpectedTest string = "[refactor] Removed middleware components"

func TestDeletedGolang(t *testing.T) {
	gitOutput := `
		D	middleware/rate.go
	`

	code.ExecCommand = func(name string, args ...string) *exec.Cmd {
		return tests.FakeExecCommand(gitOutput)
	}
	defer func() {
		code.ExecCommand = exec.Command
	}()

	c := &code.Code{}

	files := []string{"middleware/rate.go"}
	msg, err := c.FormattedCode(files)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !strings.EqualFold(strings.TrimSpace(msg), strings.TrimSpace(deleteExpectedTest)) {
		t.Errorf("Expected commit message:\n%q\nGot:%q", deleteExpectedTest, msg)
	}

	fmt.Println("Formatted commit message:", msg)
}
