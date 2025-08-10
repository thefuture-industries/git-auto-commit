package parser

import (
	"fmt"
	"git-auto-commit/pkg/file"
	"git-auto-commit/pkg/pkgerror"
	"os"
	"path/filepath"
)

func (p *Parser) ParserIndex(directory string) (string, error) {
	info, err := os.Stat(directory)
	if err != nil {
		return "", pkgerror.CreateError(pkgerror.Err_FailedToReadFile)
	}

	// check this is file
	if !info.IsDir() {
		return p.CreateAutoCommitMsg(&directory, nil, ""), nil
	}

	files := file.GetFilesInDir(directory)
	if len(files) == 0 {
		return "", pkgerror.CreateError(pkgerror.Err_FileNotFound)
	}

	// check .ext and max count
	extCount := map[string]int{}
	maxExt := ""
	maxCount := 0

	for _, file := range files {
		ext := filepath.Ext(file)
		extCount[ext]++

		if extCount[ext] > maxCount {
			maxCount = extCount[ext]
			maxExt = ext
		}
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

	maxExt = fmt.Sprintf("file%s", maxExt)
	return p.CreateAutoCommitMsg(&maxExt, &formatted, ""), nil
}
