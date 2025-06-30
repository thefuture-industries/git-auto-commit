package main

import (
	"fmt"
	"git-auto-commit/achelper"
	"git-auto-commit/diff"
	"git-auto-commit/git"
	"git-auto-commit/parser"
	"os"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-w" || os.Args[1] == "--watch") {
		path, err := git.GetGitRoot()
		if err != nil {
			achelper.ErrorLogger(err)
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
		AutoCommit()
	}
}

func AutoCommit() {
	achelper.GetVersion(false)

	files, err := diff.GetStagedFiles()
	if err != nil {
		achelper.ErrorLogger(fmt.Errorf("error getting staged files: %s", err.Error()))
		return
	}

	if len(files) == 0 {
		achelper.InfoLogger("No files staged for commit.")
		return
	}

	parserMsg, err := parser.Parser(files)
	if err != nil {
		achelper.ErrorLogger(err)
		return
	}

	if err := git.Commit(parserMsg); err != nil {
		achelper.ErrorLogger(fmt.Errorf("error committing: %s", err.Error()))
		return
	}
}
