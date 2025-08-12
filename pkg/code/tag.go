package code

import (
	"git-auto-commit/infra/constants"
	"git-auto-commit/pkg/commit"
	"path/filepath"
)

func (c *Code) WithTag(files []string, formatted string, added, modified, deleted []string) string {
	// check .ext and max count
	extCount := map[string]int{}
	maxFile := ""
	maxExtCount := 0

	for _, file := range files {
		ext := filepath.Ext(file)
		extCount[ext]++

		if extCount[ext] > maxExtCount {
			maxExtCount = extCount[ext]
			maxFile = file
		}
	}

	ptag := commit.DetectTagByFile(&maxFile, "")
	if ptag != "" {
		return commit.CreateAutoCommitMsg(&maxFile, &formatted, "")
	}

	statusCount := map[string]int{
		constants.NameStatus_Added:    len(added),
		constants.NameStatus_Modified: len(modified),
		constants.NameStatus_Deleted:  len(deleted),
	}

	maxStatus := ""
	maxCount := 0
	for status, count := range statusCount {
		if count > maxCount {
			maxCount = count
			maxStatus = status
		}
	}

	return commit.CreateAutoCommitMsg(&maxFile, &formatted, maxStatus)
}
