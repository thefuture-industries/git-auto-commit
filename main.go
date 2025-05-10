package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-w" || os.Args[1] == "--watch") {
		pathRoot, err := GetGitRoot()
		if err != nil {
			ErrorLogger(err)
			return
		}



		WatchCommit("")
	} else {
		AutoCommit()
	}
}

func AutoCommit() {
	files, err := GetStagedFiles()
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
