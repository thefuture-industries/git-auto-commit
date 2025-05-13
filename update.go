package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func AutoCommitUpdate() {
	root, err := GetGitRoot()
	if err != nil {
		ErrorLogger(err)
		return
	}

	versionFile := filepath.Join(root, ".git", "hooks", VERSION_FILE)

	version, err := os.ReadFile(versionFile)
	if err != nil {
		ErrorLogger(fmt.Errorf("unknown version for auto-commit, please re-install: %w", err))
		return
	}

	resp, err := http.Get(GITHUB_API_REPO_URL + "/releases/latest")
	if err != nil {
		ErrorLogger(fmt.Errorf("could not check latest version: %w", err))
		return
	}
	defer resp.Body.Close()

	var data struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		ErrorLogger(fmt.Errorf("could not parse version info: %w", err))
		return
	}

	if strings.TrimSpace(string(version)) == strings.TrimSpace(data.TagName) {
		fmt.Printf("\033[33m[!] you have the latest version installed %s\033[0m\n", strings.TrimSpace(data.TagName))
		return
	}

	fmt.Printf("updating to version %s...\n", strings.TrimSpace(data.TagName))

	// ps1 || bash
	if runtime.GOOS == "windows" {
		script := fmt.Sprintf("%s/scripts/update-windows-auto-commit.ps1", root)
		cmd := exec.Command("powershell", "-NoProfile", "-ExecutionPolicy", "Bypass", "-File", script)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			ErrorLogger(fmt.Errorf("failed to run update script: %w", err))
			return
		}

		return
	}

	script := fmt.Sprintf("%s/scripts/update-linux-auto-commit.sh", root)
	cmd := exec.Command("bash", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		ErrorLogger(fmt.Errorf("failed to run bash script: %w", err))
		return
	}
}
