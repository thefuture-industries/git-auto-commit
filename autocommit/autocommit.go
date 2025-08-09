package autocommit

import (
	"fmt"
	"git-auto-commit/infra/logger"
	"git-auto-commit/pkg"
	"git-auto-commit/pkg/git"
)

func AutoCommit() {
	GetVersion(false)

	files, err := git.GetStagedFiles()
	if err != nil {
		logger.ErrorLogger(fmt.Errorf("error getting staged files: %s", err.Error()))
		return
	}

	if len(files) == 0 {
		logger.InfoLogger("No files staged for commit.")
		return
	}

	parserMsg, err := pkg.Parser(files)
	if err != nil {
		logger.ErrorLogger(err)
		return
	}

	if err := git.Commit(parserMsg); err != nil {
		logger.ErrorLogger(fmt.Errorf("error committing: %s", err.Error()))
		return
	}
}