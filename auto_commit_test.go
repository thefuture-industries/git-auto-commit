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
	GetStagedFiles = func() ([]string, error) { return []string{}, nil }
    Parser = func(files []string) (string, error) { return "", nil }
    Commit = func(msg string) error { return nil }
    ErrorLogger = func(err error) { t.Errorf("unexpected error: %v", err) }
    calledInfo := ""
    InfoLogger = func(msg string) { calledInfo = msg }
    GetVersion = func(show bool) {}

	getStagedFilesMock = func() ([]string, error) { return []string{}, nil }
	infoLoggerMock = func(msg string) {
		if msg != "No files staged for commit." {
			t.Errorf("unexpected info log: %s", msg)
		}
	}

	errorLoggerMock = func(err error) { t.Errorf("unexpected error: %v", err) }
	AutoCommit()
}
