package git

import (
	"fmt"
	"git-auto-commit/infra/logger"
	"os"
	"os/exec"
)

func (g *Git) Commit(commitMsg string) error {
	logger.GitLogger(fmt.Sprintf("commit is: %s", commitMsg))

	cmd := exec.Command("git", "commit", "-m", commitMsg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
