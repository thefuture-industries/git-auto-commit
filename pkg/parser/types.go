package parser

import "git-auto-commit/pkg/code"

type Parser struct {
	Code code.CodeInterface
}

type ParserInterface interface {
	// parser.go
	ParserIndex(directory string) (string, error)
}
