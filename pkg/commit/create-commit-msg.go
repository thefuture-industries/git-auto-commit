package commit

import (
	"fmt"
	"git-auto-commit/infra/constants"
)

func CreateAutoCommitMsg(filename, msg *string, tag string) string {
	ext := DetectTagByFile(filename, tag)

	msgCommit, ok := constants.Ratio_Commit[ext]
	if ok {
		return ext + " " + msgCommit
	}

	if msg != nil {
		return ext + " " + *msg
	}

	return fmt.Sprintf("%s Processed file - %s is refactored", constants.Type_CommitRefactor, *filename)
}
