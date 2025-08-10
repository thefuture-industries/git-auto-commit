package main

import (
	"fmt"
	"git-auto-commit/autocommit"
	"git-auto-commit/infra/logger"
	"git-auto-commit/pkg/git"
	"os"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)

	if len(args) > 1 && (args[0] == "-w" || args[0] == "--watch") {
		path, err := git.GetGitRoot()
		if err != nil {
			logger.ErrorLogger(err)
			return
		}

		if len(args) > 2 {
			path = fmt.Sprintf("%s/%s", path, args[1])
		}

		autocommit.Watch(path)
	} else if len(args) > 1 && (args[0] == "-v" || args[0] == "--version") {
		autocommit.GetVersion(true)
	} else if len(args) > 1 && (args[0] == "-u" || args[0] == "--update") {
		autocommit.Update()
	} else {
		autocommit.AutoCommit()
	}
}
