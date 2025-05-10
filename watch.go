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

func WatchCommit(path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		ErrorLogger(err)
		return
	}
	defer watcher.Close()
	

	InfoLogger("Started commit watcher...")

	root, err := GetGitRoot()
	if err != nil {
		ErrorLogger(err)
		return
	}

	if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && !strings.HasPrefix(path, ".git") {
			if err := watcher.Add(path); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		ErrorLogger(err)
		return
	} // "." -> root git OR "/user"

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
			if strings.Contains(filepath.ToSlash(event.Name), "/.git/") {
				continue
			}

			if event.Op&fsnotify.Write == fsnotify.Write {
				// check need is commit
				cmd := exec.Command("git", "diff", "--quiet")
				if err := cmd.Run(); err == nil {
					continue
				}

				if err := exec.Command("git", "add", ".").Run(); err != nil {
					ErrorLogger(err)
					return
				}

				files, err := GetStagedFiles()
				if err != nil {
					ErrorLogger(fmt.Errorf("error getting staged files: %s", err.Error()))
					return
				}

				if len(files) == 0 {
					InfoLogger("No files staged for commit.")
				}

				parser, err := Parser(files)
				if err != nil {
					ErrorLogger(err)
					return
				}

				if err := Commit(parser); err != nil {
					ErrorLogger(err)
				}
			}
		case err := <-watcher.Errors:
			ErrorLogger(err)
		}

		time.Sleep(3 * time.Second)
	}
}
