package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Commit(commitMsg string) error {
	GitLogger(fmt.Sprintf("commit is: %s", commitMsg))

	cmd := exec.Command("git", "commit", "-m", commitMsg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
