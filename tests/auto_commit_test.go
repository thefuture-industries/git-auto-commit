package tests

import (
	"testing"
	"git-auto-commit"
)

var (
	getStagedFilesMock func() ([]string, error)
	parserMock         func([]string) (string, error)
	commitMock         func(string) error
	errorLoggerMock    func(error)
	infoLoggerMock     func(string)
	getVersionMock     func(bool)
)

func TestAutoCommit_NoStagedFiles(t *testing.T) {
	calledInfo := ""

	getStagedFilesMock = func() ([]string, error) { return []string{}, nil }
	parserMock = func(files []string) (string, error) { return "", nil }
	commitMock = func(msg string) error { return nil }
	errorLoggerMock = func(err error) { t.Errorf("unexpected error: %v", err) }
	infoLoggerMock = func(msg string) { calledInfo = msg }
	getVersionMock = func(show bool) {}

	main.AutoCommit()

	if calledInfo != "No files staged for commit." {
		t.Errorf("expected info log 'No files staged for commit.', got '%s'", calledInfo)
	}
}
