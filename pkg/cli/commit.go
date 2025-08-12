package cli

import (
	"git-auto-commit/infra/logger"
	"git-auto-commit/pkg/pkgerror"
)

func (cli *CLI) AutoCommit() {
	cli.GetVersion(false)

	directory, err := cli.Git.GetStagedCountDirectory()
	if err != nil {
		logger.ErrorLogger(pkgerror.Err_FailedToGetDiff)
		return
	}

	if directory == "" {
		logger.InfoLogger(pkgerror.Err_NoStagedFiles.Error())
		return
	}

	parserMsg, err := cli.Parser.ParserIndex(directory)
	if err != nil {
		return
	}

	if err := cli.Git.Commit(parserMsg); err != nil {
		return
	}
}
