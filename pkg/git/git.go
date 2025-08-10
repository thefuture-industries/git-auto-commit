package git

import (
	"bytes"
	"git-auto-commit/pkg/pkgerror"
	"os/exec"
	"strings"
)

func GetGitRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "", pkgerror.Err_GitNotRepository
	}

	root := strings.TrimSpace(out.String())
	return root, nil
}

func GetCurrentBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return "", pkgerror.Err_GitNotInstalled
	}

	return strings.TrimSpace(string(out)), nil
}
