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
		ErrorLogger(err)
		return
	}

	versionFile := filepath.Join(root, ".git", "hooks", VERSION_FILE)

	version, err := os.ReadFile(versionFile)
	if err != nil {
		ErrorLogger(fmt.Errorf("unknown version for auto-commit, please re-install"))
		return
	}

	fmt.Println("[git auto-commit] current version:", string(version))

	resp, err := http.Get(GITHUB_API_REPO_URL + "/releases/latest")
	if err != nil {
		ErrorLogger(fmt.Errorf("could not check latest version"))
		return
	}
	defer resp.Body.Close()

	var data struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		ErrorLogger(fmt.Errorf("could not parse version info"))
		return
	}

	if strings.TrimSpace(string(version)) != strings.TrimSpace(data.TagName) {
		fmt.Println("A new version is available: ", strings.TrimSpace(data.TagName))
		fmt.Println("Please update! 'git auto -u'")
	}
}
