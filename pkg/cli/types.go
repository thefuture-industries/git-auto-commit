package cli

import (
	"git-auto-commit/pkg/git"
	"git-auto-commit/pkg/parser"
)

type CLI struct {
	Git    git.GitInterface
	Parser parser.ParserInterface
}
