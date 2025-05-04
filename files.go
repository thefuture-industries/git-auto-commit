package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func AdFileFolder() ([]string, error) {
	cmd := exec.Command("git", "diff", "--cached", "--name-status")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var add []string
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "A") {
			add = append(add, strings.TrimSpace(line[2:]))
		}
	}

	return add, nil
}

func DelFileFolder() ([]string, error) {
	cmd := exec.Command("git", "diff", "--cached", "--name-status")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var del []string
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "D") {
			del = append(del, strings.TrimSpace(line[2:]))
		}
	}

	return del, nil
}

func RnFileFolder() ([]string, error) {
	cmd := exec.Command("git", "diff", "--cached", "--name-status")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var rn []string
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "R") {
			parts := strings.Fields(line)
			oldFileName := parts[1]
			newFileName := parts[2]
			rn = append(rn, fmt.Sprintf("%s -> %s", oldFileName, newFileName))
		}
	}

	return rn, nil
}

func ChFileFolder() ([]string, error) {
	cmd := exec.Command("git", "status", "--porcelain")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	var ch []string
	lines := strings.Split(out.String(), "\n")
	for _, line := range lines {
		if len(line) < 4 {
			continue
		}
		ch = append(ch, strings.TrimSpace(line[3:]))
	}

	return ch, nil
}
