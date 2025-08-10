package parser

import (
	"git-auto-commit/pkg/file"
	"git-auto-commit/pkg/pkgerror"
	"os"
)

func (p *Parser) ParserIndex(directory string) (string, error) {
	info, err := os.Stat(directory)
	if err != nil {
		return "", pkgerror.CreateError(pkgerror.Err_FailedToReadFile)
	}

	if !info.IsDir() {
		return p.CreateAutoCommitMsg(&directory, nil, ""), nil
	}

	files := file.GetFilesInDir(directory)
	if len(files) == 0 {
		return "", pkgerror.CreateError(pkgerror.Err_FileNotFound)
	}

	formatted, err := p.Code.FormattedCode(files)
	if err != nil {
		return "", pkgerror.CreateError(err)
	}

	if formatted == "" {
		formattedByRemote, err := p.Code.FormattedByRemote("")
		if err != nil {
			return "", pkgerror.CreateError(pkgerror.Err_RemoteUnavailable)
		}

		if formattedByRemote != "" {
			formatted = formattedByRemote
		}
	}

	if formatted == "" {
		formattedByBranch, err := p.Code.FormattedByBranch()
		if err != nil {
			return "", pkgerror.CreateError(pkgerror.Err_BranchNotFound)
		}

		if formattedByBranch != "" {
			formatted = formattedByBranch
		}
	}

	return formatted, nil
}
