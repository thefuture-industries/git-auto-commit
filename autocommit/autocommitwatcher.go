package autocommit

import (
	"fmt"
	"git-auto-commit/infra/constants"
	"git-auto-commit/infra/logger"
	"git-auto-commit/pkg/git"
	"git-auto-commit/pkg/parser"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
)

func Watch(path string) {
	GetVersion(false)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.ErrorLogger(err)
		return
	}
	defer watcher.Close()

	logger.InfoLogger("Started commit watcher...")

	if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && !strings.HasPrefix(path, ".git") {
			if err := watcher.Add(path); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		logger.ErrorLogger(err)
		return
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		logger.InfoLogger("shutdown work watcher...")
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
					logger.ErrorLogger(err)
					return
				}

				directory, err := git.GetStagedCountDirectory()
				if err != nil {
					logger.ErrorLogger(fmt.Errorf("error getting staged files: %s", err.Error()))
					return
				}

				if directory == "" {
					logger.InfoLogger("No files staged for commit.")
					return
				}

				parser, err := parser.Parser(directory)
				if err != nil {
					logger.ErrorLogger(err)
					return
				}

				if uint16(len(parser)) >= constants.MAX_COMMIT_LENGTH_WATCHER {
					if err := git.Commit(parser); err != nil {
						logger.ErrorLogger(err)
					}
				}
			}
		case err := <-watcher.Errors:
			logger.ErrorLogger(err)
		}

		time.Sleep(constants.COMMIT_TIME)
	}
}
