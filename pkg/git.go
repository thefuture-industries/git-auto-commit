package pkg

import (
	"bytes"
	"git-auto-commit/constants"
	"os/exec"
	"path/filepath"
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
