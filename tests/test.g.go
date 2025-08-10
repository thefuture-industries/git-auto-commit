package tests

import (
	"git-auto-commit/autocommit"
	"git-auto-commit/infra/logger"
	"git-auto-commit/pkg/git"
	"git-auto-commit/pkg/parser"
)

type Mocks struct {
	GetStagedFiles func() ([]string, error)
	Parser         func(string) (string, error)
	Commit         func(string) error
	ErrorLogger    func(error)
	InfoLogger     func(string)
	GetVersion     func(bool)
}

func SaveMocks() *Mocks {
	return &Mocks{
		GetStagedFiles: git.GetStagedFiles,
		Parser:         parser.Parser,
		Commit:         git.Commit,
		ErrorLogger:    logger.ErrorLogger,
		InfoLogger:     logger.InfoLogger,
	}
}

func (m *Mocks) Apply() {
	git.GetStagedFiles = m.GetStagedFiles
	parser.Parser = m.Parser
	git.Commit = m.Commit
	logger.ErrorLogger = m.ErrorLogger
	logger.InfoLogger = m.InfoLogger
	autocommit.GetVersion = m.GetVersion
}
