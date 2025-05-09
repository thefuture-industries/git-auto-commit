package main

import (
	"encoding/json"
	"fmt"
	"git-auto-commit/types"
	"io"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
)

func ExtractIssueNumber(branch string) string {
	re := regexp.MustCompile(`(\\d+)`)
	match := re.FindStringSubmatch(branch)
	if len(match) > 1 {
		return match[1]
	}

	return ""
}

func GetOwnerRepository() (string, string, error) {
	cmd := exec.Command("git", "remote", "get-url", "origin")
	out, err := cmd.Output()
	if err != nil {
		return "", "", err
	}

	url := strings.TrimSpace(string(out))
	regex := regexp.MustCompile(`[:/]([^/:]+)/([^/]+?)(?:\.git)?$`)

	match := regex.FindStringSubmatch(url)
	if len(match) == 3 {
		return match[1], match[2], nil
	}

	return "", "", fmt.Errorf("could not parse owner/repository from remote url: %s", url)
}

func GetIssueData(owner, repo, issue, token string) (string, uint32, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%s", owner, repo, issue)
	req, _ := http.NewRequest("GET", url, nil)
	if token != "" {
		req.Header.Set("Authorization", "token "+token)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var githubIssue types.GithubIssue
	if err := json.Unmarshal(body, &githubIssue); err != nil {
		return "", 0, err
	}

	return githubIssue.Title, githubIssue.Number, nil
}
