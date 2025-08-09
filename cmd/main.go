package main

import (
	"fmt"
	"git-auto-commit/autocommit"
	"git-auto-commit/infra/logger"
	"git-auto-commit/pkg/git"
	"os"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-w" || os.Args[1] == "--watch") {
		path, err := git.GetGitRoot()
		if err != nil {
			logger.ErrorLogger(err)
			return
		}

		if len(os.Args) > 2 {
			path = fmt.Sprintf("%s/%s", path, os.Args[2])
		}

		autocommit.Watch(path)
	} else if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		autocommit.GetVersion(true)
	} else if len(os.Args) > 1 && (os.Args[1] == "-u" || os.Args[1] == "--update") {
		autocommit.Update()
	} else {
		autocommit.AutoCommit()
	}
}
