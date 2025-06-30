package main

import (
	"git-auto-commit/diff"
	"git-auto-commit/parser"
)

type Mocks struct {
	GetStagedFiles func() ([]string, error)
	Parser         func([]string) (string, error)
	Commit         func(string) error
	ErrorLogger    func(error)
	InfoLogger     func(string)
	GetVersion     func(bool)
}

func SaveMocks() *Mocks {
	return &Mocks{
		GetStagedFiles: diff.GetStagedFiles,
		Parser:         parser.Parser,
		Commit:         git.Commit,
		ErrorLogger:    ErrorLogger,
		InfoLogger:     InfoLogger,
		GetVersion:     GetVersion,
	}
}

func (m *Mocks) Apply() {
	diff.GetStagedFiles = m.GetStagedFiles
	parser.Parser = m.Parser
	Commit = m.Commit
	ErrorLogger = m.ErrorLogger
	InfoLogger = m.InfoLogger
	GetVersion = m.GetVersion
}
