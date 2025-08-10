package parser

import (
	"fmt"
	"git-auto-commit/infra/constants"
)

func (p *Parser) CreateAutoCommitMsg(filename, msg *string, changed string) string {
	ext := p.DetectTagByFile(filename, changed)

	msgCommit, ok := constants.Ratio_Commit[ext]
	if ok {
		return ext + " " + msgCommit
	}

	if msg != nil {
		return ext + " " + *msg
	}

	return fmt.Sprintf("%s Processed file - %s is refactored", constants.Type_CommitRefactor, *filename)
}
