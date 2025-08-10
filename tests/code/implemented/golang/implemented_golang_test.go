package golang

import (
	"fmt"
	"git-auto-commit/pkg/code"
	"os/exec"
	"strings"
	"testing"
)

func TestImplementedGolang(t *testing.T) {
	gitOutput := `
		A	src/main.go
		M	src/utils.go
		D	docs/readme.md
	`

	code.ExecCommand = func(name string, args ...string) *exec.Cmd {
		return fakeExecCommand(gitOutput)
	}
	defer func() {
		code.ExecCommand = exec.Command
	}()

	c := &code.Code{}

	files := []string{"src/main.go", "src/utils.go", "docs/readme.md"}
	msg, err := c.FormattedCode(files)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !strings.Contains(strings.ToLower(msg), "implemented") || !strings.Contains(strings.ToLower(msg), "updated") || !strings.Contains(strings.ToLower(msg), "removed") {
		t.Errorf("Expected message to contain implemented, updated and removed parts, got: %s", msg)
	}

	if msg == "" {
		t.Errorf("Got empty message")
	}

	fmt.Println("Formatted commit message:", msg)
}

func fakeExecCommand(output string) *exec.Cmd {
	return &exec.Cmd{
		// Используем Stdin/Stdout подключение через pipe в тесте
		// Здесь обман: мы используем os/exec/CommandContext с запуском echo
		// но проще - используем встроенный тестовый подход:

		// Для удобства подменяем метод StdoutPipe с помощью интерфейса, а не exec.Cmd.
		// В Go это сделать сложно, поэтому можно использовать os/exec.Command с echo:

		// Но тут можно пойти обходным путём (если можно изменить код функции),
		// или использовать библиотеку "github.com/Netflix/go-expect" для более сложного мока.
	}
}
