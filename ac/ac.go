package acpkg

import (
	"fmt"
	"git-auto-commit/achelper"
	"git-auto-commit/achelper/logger"
	"git-auto-commit/diff"
	"git-auto-commit/git"
	"git-auto-commit/parser"
)

func AutoCommit() {
	achelper.GetVersion(false)

	files, err := diff.GetStagedFiles()
	if err != nil {
		logger.ErrorLogger(fmt.Errorf("error getting staged files: %s", err.Error()))
		return
	}

	if len(files) == 0 {
		logger.InfoLogger("No files staged for commit.")
		return
	}

	parserMsg, err := parser.Parser(files)
	if err != nil {
		logger.ErrorLogger(err)
		return
	}

	if err := git.Commit(parserMsg); err != nil {
		logger.ErrorLogger(fmt.Errorf("error committing: %s", err.Error()))
		return
	}
}
