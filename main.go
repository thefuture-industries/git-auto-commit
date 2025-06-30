package main

import (
	"fmt"
	"git-auto-commit/achelper"
	"git-auto-commit/achelper/logger"
	"git-auto-commit/ac"
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

		achelper.WatchCommit(path)
	} else if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		achelper.GetVersion(true)
	} else if len(os.Args) > 1 && (os.Args[1] == "-u" || os.Args[1] == "--update") {
		achelper.AutoCommitUpdate()
	} else {
		ac.AutoCommit()
	}
}
