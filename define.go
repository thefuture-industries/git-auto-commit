package main

import "time"

const (
	MAX_LINE_LENGTH           uint16        = 1024
	MAX_COMMIT_LENGTH         uint16        = 300
	MAX_COMMIT_LENGTH_WATCHER uint16        = 60
	COMMIT_TIME               time.Duration = 17 * time.Second
)
