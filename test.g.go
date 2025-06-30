package main

import (
	"git-auto-commit/achelper"
	"git-auto-commit/achelper/logger"
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
		ErrorLogger:    logger.ErrorLogger,
		InfoLogger:     logger.InfoLogger,
		GetVersion:     achelper.GetVersion,
	}
}

func (m *Mocks) Apply() {
	diff.GetStagedFiles = m.GetStagedFiles
	parser.Parser = m.Parser
	git.Commit = m.Commit
	logger.ErrorLogger = m.ErrorLogger
	logger.InfoLogger = m.InfoLogger
	achelper.GetVersion = m.GetVersion
}
