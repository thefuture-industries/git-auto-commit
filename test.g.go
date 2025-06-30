package main

import (
	"git-auto-commit/achelper"
	"git-auto-commit/diff"
	"git-auto-commit/git"
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
		ErrorLogger:    achelper.ErrorLogger,
		InfoLogger:     achelper.InfoLogger,
		GetVersion:     GetVersion,
	}
}

func (m *Mocks) Apply() {
	diff.GetStagedFiles = m.GetStagedFiles
	parser.Parser = m.Parser
	git.Commit = m.Commit
	achelper.ErrorLogger = m.ErrorLogger
	achelper.InfoLogger = m.InfoLogger
	GetVersion = m.GetVersion
}
