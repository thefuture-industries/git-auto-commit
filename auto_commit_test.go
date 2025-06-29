package main

import (
	"testing"
)

func TestAutoCommit_NoStagedFiles(t *testing.T) {
	calledInfo := ""

	GetStagedFiles = func() ([]string, error) { return []string{}, nil }
    Parser = func(files []string) (string, error) { return "", nil }
    Commit = func(msg string) error { return nil }
    ErrorLogger = func(err error) { t.Errorf("unexpected error: %v", err) }
    InfoLogger = func(msg string) { calledInfo = msg }
    GetVersion = func(show bool) {}

	AutoCommit()

    if calledInfo != "No files staged for commit." {
        t.Errorf("expected info log 'No files staged for commit.', got '%s'", calledInfo)
    }
}
