package cli

import (
	"fmt"
	"git-auto-commit/config"
	"git-auto-commit/infra/constants"
	"git-auto-commit/infra/logger"
	"git-auto-commit/pkg/file"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func (cli *CLI) Update() {
	root, err := cli.Git.GetGitRoot()
	if err != nil {
		logger.ErrorLogger(err)
		return
	}

	versionFile := filepath.Join(root, ".git", "hooks", constants.VERSION_FILE)

	version, err := os.ReadFile(versionFile)
	if err != nil {
		logger.ErrorLogger(fmt.Errorf("unknown version for auto-commit, please re-install: %w", err))
		return
	}

	resp, err := http.Get(constants.GITHUB_API_REPO_URL + "/releases/latest")
	if err != nil {
		logger.ErrorLogger(fmt.Errorf("could not check latest version: %w", err))
		return
	}
	defer resp.Body.Close()

	var data struct {
		TagName string `json:"tag_name"`
	}
	if err := config.JSON.NewDecoder(resp.Body).Decode(&data); err != nil {
		logger.ErrorLogger(fmt.Errorf("could not parse version info: %w", err))
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
		scriptUpdate = fmt.Sprintf("%s/raw/main/scripts/%s", constants.GITHUB_REPO_URL, constants.GITHUB_SCRIPT_AUTOCOMMIT_UPDATE_WIN)
		scriptUpdateExt = ".ps1"
	} else {
		scriptUpdate = fmt.Sprintf("%s/raw/main/scripts/%s", constants.GITHUB_REPO_URL, constants.GITHUB_SCRIPT_AUTOCOMMIT_UPDATE_LINUX)
		scriptUpdateExt = ".sh"
	}

	tmpFile := filepath.Join(os.TempDir(), "auto-commit-update"+scriptUpdateExt)
	err = file.DownloadFile(scriptUpdate, tmpFile)
	if err != nil {
		logger.ErrorLogger(fmt.Errorf("failed to download update script: %v", err))
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
		logger.ErrorLogger(fmt.Errorf("failed to run update script: %v", err))
		return
	}
}
