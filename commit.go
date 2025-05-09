package main

import (
	"fmt"
)

func Commit(commitMsg string) error {
	GitLogger(fmt.Sprintf("commit is: %s", commitMsg))

	// cmd := exec.Command("git", "commit", "-m", commitMsg)
	// return cmd.Run()
	return nil
}
