package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func NewFileFolder() ([]string, error) {
	cmd := exec.Command("git", "diff", "--cached", "--name-status")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var added []string
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "A") {
			added = append(added, strings.TrimSpace(line[2:]))
		}
	}

	return added, nil
}

func DeleteFileFolder() ([]string, error) {
	cmd := exec.Command("git", "diff", "--cached", "--name-status")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var deleted []string
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "D") {
			deleted = append(deleted, strings.TrimSpace(line[2:]))
		}
	}

	return deleted, nil
}

func RenameFileFolder() ([]string, error) {
	cmd := exec.Command("git", "diff", "--cached", "--name-status")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var renamed []string
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "R") {
			parts := strings.Fields(line)
			oldFileName := parts[1]
			newFileName := parts[2]
			renamed = append(renamed, fmt.Sprintf("%s -> %s", oldFileName, newFileName))
		}
	}

	return renamed, nil
}
