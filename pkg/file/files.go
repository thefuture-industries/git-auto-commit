package file

import (
	"os"
	"path/filepath"
)

func GetFilesInDir(directory string) []string {
	var files []string

	if err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	}); err != nil {
		return []string{}
	}

	return files
}
