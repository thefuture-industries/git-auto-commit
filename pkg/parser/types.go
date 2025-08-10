package parser

import "git-auto-commit/pkg/code"

type Parser struct {
	Code code.CodeInterface
}

type ParserInterface interface {
	// create-commit-msg.go
	CreateAutoCommitMsg(filename, msg *string, changed string) string

	// detect-tag.go
	DetectTagByFile(filename *string, changed string) string

	// parser.go
	ParserIndex(directory string) (string, error)
}
