package tests

import (
	"git-auto-commit/autocommit"
	"git-auto-commit/infra/logger"
	"git-auto-commit/pkg"
	"git-auto-commit/pkg/git"
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
		GetStagedFiles: git.GetStagedFiles,
		Parser:         pkg.Parser,
		Commit:         git.Commit,
		ErrorLogger:    logger.ErrorLogger,
		InfoLogger:     logger.InfoLogger,
	}
}

func (m *Mocks) Apply() {
	git.GetStagedFiles = m.GetStagedFiles
	pkg.Parser = m.Parser
	git.Commit = m.Commit
	logger.ErrorLogger = m.ErrorLogger
	logger.InfoLogger = m.InfoLogger
	autocommit.GetVersion = m.GetVersion
}
