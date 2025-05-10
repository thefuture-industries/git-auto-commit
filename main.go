package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-w" || os.Args[1] == "--watch") {
		Watch()
	} else {
		Auto()
	}
}

func Auto() {
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

func Watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		ErrorLogger(err)
		return
	}
	defer watcher.Close()

	root, err := GetGitRoot()
	if err != nil {
		ErrorLogger(err)
		return
	}

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && !strings.HasPrefix(path, ".git") {
			watcher.Add(path)
		}

		return nil
	}) // "." -> root git

	var commitMsg string

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		InfoLogger("Shutdown work watcher...")
		os.Exit(0)
	}()

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				exec.Command("git", "add", ".").Run()

				files, err := GetStagedFiles()
				if err != nil {
					ErrorLogger(fmt.Errorf("error getting staged files: %s", err.Error()))
					return
				}

				if len(files) == 0 {
					InfoLogger("No files staged for commit.")
					return
				}

				parser, err := Parser(files)
				if err != nil {
					ErrorLogger(err)
					return
				}

				commitMsg += parser
				if len(strings.Fields(commitMsg)) >= 300 {
					if err := Commit(commitMsg); err != nil {
						ErrorLogger(err)
						return
					}

					commitMsg = ""
				}
			}
		case err := <-watcher.Errors:
			ErrorLogger(err)
		}

		time.Sleep(1 * time.Second)
	}
}
