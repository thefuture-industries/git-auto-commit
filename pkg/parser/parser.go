package parser

import (
	"fmt"
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
}
