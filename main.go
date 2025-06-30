package main

import (
	"fmt"
	"git-auto-commit/diff"
	"git-auto-commit/git"
	"os"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-w" || os.Args[1] == "--watch") {
		path, err := git.GetGitRoot()
		if err != nil {
			ErrorLogger(err)
			return
		}

		if len(os.Args) > 2 {
			path = fmt.Sprintf("%s/%s", path, os.Args[2])
		}

		WatchCommit(path)
	} else if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		GetVersion(true)
	} else if len(os.Args) > 1 && (os.Args[1] == "-u" || os.Args[1] == "--update") {
		AutoCommitUpdate()
	} else {
		AutoCommit()
	}
}

func AutoCommit() {
	GetVersion(false)

	files, err := diff.GetStagedFiles()
	if err != nil {
		ErrorLogger(fmt.Errorf("error getting staged files: %s", err.Error()))
		return
	}

	if len(files) == 0 {
		InfoLogger("No files staged for commit.")
		return
	}

	parserMsg, err := Parser(files)
	if err != nil {
		ErrorLogger(err)
		return
	}

	if err := Commit(parserMsg); err != nil {
		ErrorLogger(fmt.Errorf("error committing: %s", err.Error()))
		return
	}
}
