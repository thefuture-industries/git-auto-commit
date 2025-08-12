package main

import (
	"fmt"
	"git-auto-commit/infra/logger"
	"git-auto-commit/pkg/cli"
	"git-auto-commit/pkg/code"
	"git-auto-commit/pkg/git"
	"git-auto-commit/pkg/parser"
	"os"
)

func main() {
	cli := &cli.CLI{
		Git:    &git.Git{},
		Parser: &parser.Parser{Code: &code.Code{}},
	}

	if len(os.Args) > 1 && (os.Args[1] == "-w" || os.Args[1] == "--watch") {
		path, err := cli.Git.GetGitRoot()
		if err != nil {
			logger.ErrorLogger(err)
			return
		}

		if len(os.Args) > 2 {
			path = fmt.Sprintf("%s/%s", path, os.Args[2])
		}

		cli.Watch(path)
	} else if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		cli.GetVersion(true)
	} else if len(os.Args) > 1 && (os.Args[1] == "-u" || os.Args[1] == "--update") {
		cli.Update()
	} else {
		cli.AutoCommit()
	}
}
