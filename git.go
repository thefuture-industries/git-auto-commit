package main

import (
	"bytes"
	"os/exec"
	"strings"
)

func GetGitRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "", err
	}

	root := strings.TrimSpace(out.String())
	return root, nil
}
