package git

import (
	"bytes"
	"fmt"
	"git-auto-commit/types"
	"io"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func ExtractIssueNumber(branch string) string {
	re := regexp.MustCompile(`\d+`)
	match := re.FindStringSubmatch(branch)
	if len(match) > 1 {
		fmt.Println(match[1])
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
	var builder strings.Builder
	builder.Reset()

	builder.WriteString("https://api.github.com/repos/")
	builder.WriteString(owner)
	builder.WriteString("/")
	builder.WriteString(repo)
	builder.WriteString("/issues/")
	builder.WriteString(issue)

	req, _ := http.NewRequest("GET", builder.String(), nil)
	if token != "" {
		req.Header.Set("Authorization", "token "+token)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufferPool.Put(buf)

	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return "", 0, err
	}

	var githubIssue types.GithubIssue
	if err := json.Unmarshal(buf.Bytes(), &githubIssue); err != nil {
		return "", 0, err
	}

	return githubIssue.Title, githubIssue.Number, nil
}
