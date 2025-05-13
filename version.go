package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func GetVersion() {
	root, err := GetGitRoot()
	if err != nil {
		ErrorLogger(fmt.Errorf("could not get git root: %w", err))
		return
	}

	versionFile := filepath.Join(root, ".git", "hooks", VERSION_FILE)

	version, err := os.ReadFile(versionFile)
	if err != nil {
		ErrorLogger(fmt.Errorf("unknown version for auto-commit, please re-install: %w", err))
		return
	}

	fmt.Println("[git auto-commit] current version:", strings.TrimSpace(string(version)))

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

	if strings.TrimSpace(string(version)) != strings.TrimSpace(data.TagName) {
		fmt.Printf("\033[33m[!] a new version is available: %s\033[0m\n", strings.TrimSpace(data.TagName))
		fmt.Printf("\033[33m[!] please update! 'git auto -u'\033[0m\n")
	}
}
