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

func TestImplementedDocGolang(t *testing.T) {
	gitOutput := `
		A	README.md
	`

	code.ExecCommand = func(name string, args ...string) *exec.Cmd {
		return tests.FakeExecCommand(gitOutput)
	}
	defer func() {
		code.ExecCommand = exec.Command
	}()

	c := &code.Code{}

	files := []string{"README.md"}
	msg, err := c.FormattedCode(files)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !strings.Contains(msg, constants.Type_CommitDocs) {
		t.Errorf("Expected commit message including:%q Got: %q", constants.Type_CommitDocs, msg)
	}

	fmt.Println("Formatted commit message:", msg)
}

func TestImplementedDocsGolang(t *testing.T) {
	gitNameStatusOutput := `
		A       docs/code/readme.md
		M       docs/code/download.txt
		M       main/docs/security.md
		D       pkg/create-commit-msg.go
		A       pkg/detect-tag.go
		A       pkg/parser/types.go
	`

	code.ExecCommand = func(name string, args ...string) *exec.Cmd {
		return tests.FakeExecCommand(gitNameStatusOutput)
	}
	defer func() {
		code.ExecCommand = exec.Command
	}()

	c := &code.Code{}

	files := []string{"docs/code/readme.md", "docs/code/download.txt"}
	msg, err := c.FormattedCode(files)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !strings.Contains(msg, constants.Type_CommitDocs) {
		t.Errorf("Expected commit message including:%q Got: %q", constants.Type_CommitDocs, msg)
	}

	fmt.Println("Formatted commit message:", msg)
}
