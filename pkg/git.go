package pkg

import (
	"bytes"
	"fmt"
	"git-auto-commit/constants"
	"git-auto-commit/utils"
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

func BuildCommitMessage(funcs []string) (string, error) {
	added, err := utils.AdFileFolder()
	if err != nil {
		return "", err
	}
	deleted, err := utils.DelFileFolder()
	if err != nil {
		return "", err
	}
	renamed, err := utils.RnFileFolder()
	if err != nil {
		return "", err
	}
	changed, err := utils.ChFileFolder()
	if err != nil {
		return "", err
	}

	var parts []string

	if len(funcs) > 0 {
		if len(funcs) == 1 {
			parts = append(parts, "added "+funcs[0]+" functionality")
		} else {
			last := funcs[len(funcs)-1]
			parts = append(parts, "added "+strings.Join(funcs[:len(funcs)-1], ", ")+" and "+last+" functionality")
		}
	}

	if len(added) > 0 {
		parts = append(parts, "including "+strings.Join(added, ", "))
	}
	if len(deleted) > 0 {
		parts = append(parts, "deleted "+strings.Join(deleted, ", "))
	}
	if len(renamed) > 0 {
		parts = append(parts, "renamed "+strings.Join(renamed, ", "))
	}
	if len(changed) > 0 {
		parts = append(parts, "changed "+strings.Join(changed, ", "))
	}

	if len(parts) == 0 {
		return "auto commit (github@git-auto-commit)", nil
	}

	commit := strings.ToLower(strings.Join(parts, ", "))
	commit = strings.ReplaceAll(commit, "  ", " ")

	if uint8(len(commit)) > constants.COMMIT_LENGTH {
		if len(funcs) > 0 {
			return strings.ToLower("added " + strings.Join(funcs, ", ") + " functionality"), nil
		}
		if len(added) > 0 {
			return strings.ToLower("including " + strings.Join(added, ", ")), nil
		}
		if len(deleted) > 0 {
			return strings.ToLower("deleted " + strings.Join(deleted, ", ")), nil
		}
		if len(renamed) > 0 {
			return strings.ToLower("renamed " + strings.Join(renamed, ", ")), nil
		}
		if len(changed) > 0 {
			return strings.ToLower("changed " + strings.Join(changed, ", ")), nil
		}
	}

	return commit, nil
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
