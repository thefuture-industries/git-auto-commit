package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Update() {
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

	resp, err := http.Get(GITHUB_REPO_URL + "/releases/latest")
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

	if strings.TrimSpace(string(version)) == strings.TrimSpace(data.TagName) {
		fmt.Println("You have the latest version installed ", strings.TrimSpace(data.TagName))
		return
	}

	fmt.Println("Updating to version ", strings.TrimSpace(data.TagName), "...")

	// Красивый вывод 

	err = os.WriteFile(versionFile, []byte(strings.TrimSpace(data.TagName)), 0644)
	if err != nil {
		ErrorLogger(fmt.Errorf("failed to update version file: %w", err))
		return
	}
	fmt.Println("Successful upgrade to version ", strings.TrimSpace(data.TagName))
}
