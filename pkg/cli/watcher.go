package cli

import (
	"errors"
	"git-auto-commit/infra/constants"
	"git-auto-commit/infra/logger"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
)

func (cli *CLI) Watch(path string) {
	cli.GetVersion(false)

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

				directory, err := cli.Git.GetStagedCountDirectory()
				if err != nil {
					logger.ErrorLogger(errors.New("error getting staged files"))
					return
				}

				if directory == "" {
					logger.InfoLogger("No files staged for commit.")
					return
				}

				parser, err := cli.Parser.ParserIndex(directory)
				if err != nil {
					logger.ErrorLogger(err)
					return
				}

				if uint16(len(parser)) >= constants.MAX_COMMIT_LENGTH_WATCHER {
					if err := cli.Git.Commit(parser); err != nil {
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
