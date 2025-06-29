package tests

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
	getStagedFilesMock = func() ([]string, error) { return []string{}, nil }
	parserMock = func(files []string) (string, error) { return "", nil }
	commitMock = func(string) error

	main.AutoCommit()
}
