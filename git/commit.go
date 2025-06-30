package git

import (
	"fmt"
	"git-auto-commit/achelper/logger"
	"os"
	"os/exec"
)

var Commit = func(commitMsg string) error {
	logger.GitLogger(fmt.Sprintf("commit is: %s", commitMsg))

	cmd := exec.Command("git", "commit", "-m", commitMsg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
