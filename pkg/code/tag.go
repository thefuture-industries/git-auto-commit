package code

import (
	"git-auto-commit/pkg/commit"
	"path/filepath"
)

func (c *Code) WithTag(files []string, formatted string) string {
	// check .ext and max count
	extCount := map[string]int{}
	maxFile := ""
	maxCount := 0

	for _, file := range files {
		ext := filepath.Ext(file)
		extCount[ext]++

		if extCount[ext] > maxCount {
			maxCount = extCount[ext]
			maxFile = file
		}
	}

	return commit.CreateAutoCommitMsg(&maxFile, &formatted, "")
}
