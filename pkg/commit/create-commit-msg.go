package commit

import (
	"fmt"
	"git-auto-commit/infra/constants"
	"strings"
)

func CreateAutoCommitMsg(filename, msg *string, tag string) string {
	ext := DetectTagByFile(filename, tag)

	msgCommit, ok := constants.Ratio_Commit[ext]
	if ok {
		return strings.TrimSpace(ext + " " + msgCommit)
	}

	if msg != nil {
		if ext != "" {
			return strings.TrimSpace(ext + " " + *msg)
		}

		return strings.TrimSpace(*msg)
	}

	return fmt.Sprintf("%s Processed file - %s is refactored", constants.Type_CommitRefactor, *filename)
}
