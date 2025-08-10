package autocommit

import (
	"git-auto-commit/infra/logger"
	"git-auto-commit/pkg/git"
	"git-auto-commit/pkg/parser"
	"git-auto-commit/pkg/pkgerror"
)

func AutoCommit() {
	GetVersion(false)

	directory, err := git.GetStagedCountDirectory()
	if err != nil {
		logger.ErrorLogger(pkgerror.Err_FailedToGetDiff)
		return
	}

	if directory == "" {
		logger.InfoLogger(pkgerror.Err_NoStagedFiles.Error())
		return
	}

	parserMsg, err := parser.Parser(directory)
	if err != nil {
		return
	}

	if err := git.Commit(parserMsg); err != nil {
		return
	}
}
