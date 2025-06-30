package main

import (
	"fmt"
	"io"
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

	var scriptUpdate string
	var scriptUpdateExt string
	if runtime.GOOS == "windows" {
		scriptUpdate = "https://github.com/thefuture-industries/git-auto-commit/raw/main/scripts/update-windows-auto-commit.ps1"
		scriptUpdateExt = ".ps1"
	} else {
		scriptUpdate = "https://github.com/thefuture-industries/git-auto-commit/raw/main/scripts/update-linux-auto-commit.sh"
		scriptUpdateExt = ".sh"
	}

	tmpFile := filepath.Join(os.TempDir(), "auto-commit-update"+scriptUpdateExt)
	err = downloadFile(scriptUpdate, tmpFile)
	if err != nil {
		ErrorLogger(fmt.Errorf("failed to download update script: %v", err))
		return
	}
	defer os.Remove(tmpFile)

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell", "-NoProfile", "-ExecutionPolicy", "Bypass", "-File", tmpFile)
	} else {
		cmd = exec.Command("bash", tmpFile)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		ErrorLogger(fmt.Errorf("failed to run update script: %v", err))
		return
	}
}

func downloadFile(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
