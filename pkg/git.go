package pkg

import (
	"bytes"
	"fmt"
	"git-auto-commit/constants"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func DetectLang(filename string) string {
	ext := filepath.Ext(filename)
	return constants.LANGUAGE_MAP[ext]
}

func GetDiff(file string) string {
	var out bytes.Buffer

	cmd := exec.Command("git", "diff", "--cached", file)

	cmd.Stdout = &out
	cmd.Run()

	return out.String()
}

func BuildCommitMessage(funcs []string) string {
	n := len(funcs)
	if n == 0 {
		return "auto commit"
	}
	if n == 1 {
		return "added " + funcs[0] + " function"
	}
	
	return "added " + strings.Join(funcs[:n-1], ", ") + " and " + funcs[n-1] + " functionality"
}

func Commit(commitMessage string) error {
	cmd := exec.Command("git", "commit", "-m", commitMessage)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to commit:", err)
	}

	return nil
}
