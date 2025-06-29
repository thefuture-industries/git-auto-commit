package main

import (
	"testing"
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
	GetStagedFiles = func() ([]string, error) { return getStagedFilesMock() }
	Parser = func(files []string) (string, error) { return parserMock(files) }
	Commit = func(msg string) error { return commitMock(msg) }
	ErrorLogger = func(err error) { errorLoggerMock(err) }
	InfoLogger = func(msg string) { infoLoggerMock(msg) }
	GetVersion = func(show bool) { getVersionMock(show) }

	getStagedFilesMock = func() ([]string, error) { return []string{}, nil }
	infoLoggerMock = func(msg string) {
		if msg != "No files staged for commit." {
			t.Errorf("unexpected info log: %s", msg)
		}
	}

	if calledInfo != "No files staged for commit." {
		t.Errorf("expected info log 'No files staged for commit.', got '%s'", calledInfo)
	}
}
