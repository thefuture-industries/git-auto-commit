package main

import (
	"errors"
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

func TestAutoCommit_ErrorGettingFiles(t *testing.T) {
	GetStagedFiles = func() ([]string, error) { return nil, errors.New("fail") }
	Parser = func(files []string) (string, error) { return "", nil }
	Commit = func(msg string) error { return nil }
	InfoLogger = func(msg string) {}
	GetVersion = func(show bool) {}

	var calledErr string
	ErrorLogger = func(err error) { calledErr = err.Error() }

	InfoLogger = func(msg string) { calledErr = msg }

	AutoCommit()

	expected := "error getting staged files: fail"
	if calledErr != expected {
		t.Errorf("expected error log '%s', got '%s'", expected, calledErr)
	}
}
