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
	added, err := utils.NewFileFolder()
	if err != nil {
		return "", err
	}
	addedMessage := strings.Join(added, ", ")

	deleted, err := utils.DeleteFileFolder()
	if err != nil {
		return "", err
	}
	deletedMessage := strings.Join(deleted, ", ")

	renamed, err := utils.RenameFileFolder()
	if err != nil {
		return "", err
	}
	renamedMessage := strings.Join(renamed, ", ")

	n := len(funcs)
	if n == 0 && len(added) == 0 && len(deleted) == 0 && len(renamed) == 0 {
		return "auto commit (github@git-auto-commit)", nil
	}

	if n == 0 && len(added) == 0 && len(deleted) == 0 {
		return strings.ToLower("renamed " + renamedMessage), nil
	}

	if n == 0 && len(added) == 0 {
		return strings.ToLower("deleted " + deletedMessage + ", and renamed " + renamedMessage), nil
	}

	if n == 0 && len(deleted) == 0 {
		return strings.ToLower("including added " + addedMessage + ", and renamed " + renamedMessage), nil
	}

	if len(added) > 0 && len(deleted) == 0 && len(renamed) == 0 {
		return strings.ToLower("added " + strings.Join(funcs[:n-1], ", ") + " and " + funcs[n-1] + " functionality, including " + addedMessage), nil
	}

	if len(added) == 0 && len(deleted) > 0 && len(renamed) == 0 {
		return strings.ToLower("added " + strings.Join(funcs[:n-1], ", ") + " and " + funcs[n-1] + " functionality, and deleted " + deletedMessage), nil
	}

	if len(added) == 0 && len(deleted) == 0 && len(renamed) > 0 {
		return strings.ToLower("added " + strings.Join(funcs[:n-1], ", ") + " and " + funcs[n-1] + " functionality, and renamed " + renamedMessage), nil
	}

	commit := strings.ToLower("added " + strings.Join(funcs[:n-1], ", ") + " and " + funcs[n-1] + " functionality, including " + addedMessage + ", deleted " + deletedMessage + " and renamed " + renamedMessage)
	commitPrepare := strings.ReplaceAll(commit, "  ", " ")
	if uint8(len(commitPrepare)) > constants.COMMIT_LENGTH {
		if n > 0 {
			return strings.ToLower("added " + strings.Join(funcs[:n-1], ", ") + ", and " + funcs[n-1] + " functionality"), nil
		}

		if len(addedMessage) > 0 {
			return strings.ToLower("including " + addedMessage), nil
		}

		if len(deletedMessage) > 0 {
			return strings.ToLower("deleted " + deletedMessage), nil
		}

		if len(renamedMessage) > 0 {
			return strings.ToLower("renamed " + renamedMessage), nil
		}
	}

	return commitPrepare, nil
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
