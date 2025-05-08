package main

import (
	"fmt"
	"os/exec"
)

func Commit(commitMsg string) error {
	fmt.Printf("\033[0;34m[git auto-commit] commit is: %s\033[0m\n", commitMsg)

	cmd := exec.Command("git", "commit", "-m", commitMsg)
	return cmd.Run()
}
