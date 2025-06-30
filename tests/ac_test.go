package tests

import (
	"errors"
	"git-auto-commit/achelper"
	"git-auto-commit/achelper/logger"
	"git-auto-commit/diff"
	"git-auto-commit/git"
	"git-auto-commit/parser"
	"testing"
)

func TestAutoCommit_NoStagedFiles(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

	calledInfo := ""
	diff.GetStagedFiles = func() ([]string, error) { return []string{}, nil }
	parser.Parser = func(files []string) (string, error) { return "", nil }
	git.Commit = func(msg string) error { return nil }
	logger.ErrorLogger = func(err error) { t.Errorf("unexpected error: %v", err) }
	logger.InfoLogger = func(msg string) { calledInfo = msg }
	achelper.GetVersion = func(show bool) {}

	ac.AutoCommit()

	if calledInfo != "No files staged for commit." {
		t.Errorf("expected info log 'No files staged for commit.', got '%s'", calledInfo)
	}
}

func TestAutoCommit_ErrorGettingFiles(t *testing.T) {
	mocks := SaveMocks()
	defer mocks.Apply()

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
