package main

import "time"

const (
	MAX_LINE_LENGTH           uint16        = 1024
	MAX_COMMIT_LENGTH         uint16        = 300
	MAX_COMMIT_LENGTH_WATCHER uint16        = 25
	COMMIT_TIME               time.Duration = 15 * time.Second

	GITHUB_API_REPO_URL string = "https://api.github.com/repos/thefuture-industries/git-auto-commit"
	VERSION_FILE        string = "auto-commit.version.txt"
	BINARY_AUTO_COMMIT  string = "auto-commit"
	GITHUB_REPO_URL     string = "https://github.com/thefuture-industries/git-auto-commit"
)
