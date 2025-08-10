package parser

import (
	"fmt"
	"git-auto-commit/pkg/code"
	"git-auto-commit/pkg/file"
	"git-auto-commit/pkg/pkgerror"
	"os"
)

var Parser = func(directory string) (string, error) {
	info, err := os.Stat(directory)
	if err != nil {
		return "", fmt.Errorf("errir getting file info")
	}

	if !info.IsDir() {
		return CreateAutoCommitMsg(&directory, nil, ""), nil
	}

	files := file.GetFilesInDir(directory)
	if len(files) == 0 {
		return "", pkgerror.CreateError(pkgerror.Err_FileNotFound)
	}

	formatted, err := code.FormattedCode(files)
	if err != nil {
		return "", err
	}

	if formatted == "" {
		formattedByRemote, err := code.FormattedByRemote("")
		if err != nil {
			return "", fmt.Errorf("failed to format by remote: %w", err)
		}

		if formattedByRemote != "" {
			formatted = formattedByRemote
		}
	}

	if formatted == "" {
		formattedByBranch, err := code.FormattedByBranch()
		if err != nil {
			return "", fmt.Errorf("failed to format by branch: %w", err)
		}
		
		if formattedByBranch != "" {
			formatted = formattedByBranch
		}
	}

	return formatted, nil
}
