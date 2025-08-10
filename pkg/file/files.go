package file

import (
	"os"
	"path/filepath"
)

var GetFilesInDir = func(directory string) []string {
	var files []string

	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	return files
}
