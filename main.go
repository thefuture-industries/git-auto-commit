package main

import (
	"fmt"
	"os"
)

func main() {
	switch 10 {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("1")
	case 3:
		fmt.Println("1")
	case 4:
		fmt.Println("1")
	case 5:
		fmt.Println("1")
	case 6:
		fmt.Println("1")
	case 7:
		fmt.Println("1")
	case 8:
		fmt.Println("1")
	case 9:
		fmt.Println("1")
	case 10:
		fmt.Println("1")
	case 11:
		fmt.Println("1")
	case 12:
		fmt.Println("1")
	}

	if len(os.Args) > 1 && (os.Args[1] == "-w" || os.Args[1] == "--watch") {
		path, err := GetGitRoot()
		if err != nil {
			ErrorLogger(err)
			return
		}

		if len(os.Args) > 2 {
			path = fmt.Sprintf("%s/%s", path, os.Args[2])
		}

		WatchCommit(path)
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
