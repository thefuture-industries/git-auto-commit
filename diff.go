package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
)

func GetDiff(file string) (string, error) {
	root, err := GetGitRoot();
	if err != nil {
		return "", err
	}

	cmd := exec.Command("git", "diff", "--cached", "--", fmt.Sprintf("%s/%s", root, file))
	var out bytes.Buffer

	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}

	return out.String(), nil
}

func GetStagedFiles() ([]string, error) {
	cmd := exec.Command("git", "diff", "--cached", "--name-only")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	var files []string
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		files = append(files, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	return files, nil
}
