package autocommit

import (
	"fmt"
	"git-auto-commit/infra/logger"
	"git-auto-commit/pkg/git"
	"git-auto-commit/pkg/parser"
)

func AutoCommit() {
	GetVersion(false)

	directory, err := git.GetStagedCountDirectory()
	if err != nil {
		logger.ErrorLogger(fmt.Errorf("error getting staged files: %s", err.Error()))
		return
	}

	if directory == "" {
		logger.InfoLogger("No files staged for commit.")
		return
	}

	parserMsg, err := parser.Parser(directory)
	if err != nil {
		logger.ErrorLogger(err)
		return
	}

	if err := git.Commit(parserMsg); err != nil {
		logger.ErrorLogger(fmt.Errorf("error committing: %s", err.Error()))
		return
	}
}
