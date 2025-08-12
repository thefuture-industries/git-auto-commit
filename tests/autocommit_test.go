package tests

import (
	"errors"
	"git-auto-commit/infra/logger"
	"git-auto-commit/pkg/cli"
	"git-auto-commit/pkg/pkgerror"
	"testing"
)

func TestAutoCommit(t *testing.T) {
	tests := []struct {
		name           string
		stagedDir      string
		stagedErr      error
		parseMsg       string
		parseErr       error
		commitErr      error
		expectInfoLog  string
		expectErrorLog bool
	}{
		{
			name:          "No staged files",
			stagedDir:     "",
			stagedErr:     nil,
			expectInfoLog: pkgerror.Err_NoStagedFiles.Error(),
		},
		{
			name:           "Failed to get staged dir",
			stagedErr:      errors.New("fail"),
			expectErrorLog: true,
		},
		{
			name:      "Parser error",
			stagedDir: "some/dir",
			parseErr:  errors.New("parse fail"),
		},
		{
			name:      "Commit error",
			stagedDir: "some/dir",
			parseMsg:  "commit message",
			commitErr: errors.New("commit fail"),
		},
		{
			name:      "Successful commit",
			stagedDir: "some/dir",
			parseMsg:  "commit message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			infoLogged := ""
			errorLogged := false

			gitMocks := &gitMocks{
				stagedDir: tt.stagedDir,
				stagedErr: tt.stagedErr,
				commitErr: tt.commitErr,
			}

			parserMocks := &parserMocks{
				parsedMsg: tt.parseMsg,
				parseErr:  tt.parseErr,
			}

			cli := &cli.CLI{
				Git:    gitMocks,
				Parser: parserMocks,
			}

			infoLogger := logger.InfoLogger
			errorLogger := logger.ErrorLogger
			gitLogger := logger.GitLogger

			defer func() {
				logger.InfoLogger = infoLogger
				logger.ErrorLogger = errorLogger
				logger.GitLogger = gitLogger
			}()

			logger.InfoLogger = func(msg string) {
				infoLogged = msg
			}

			logger.ErrorLogger = func(err error) {
				errorLogged = true
			}

			logger.GitLogger = func(msg string) {}

			cli.AutoCommit()

			if tt.expectInfoLog != "" && infoLogged != tt.expectInfoLog {
				t.Errorf("expected info log '%s', got '%s'", tt.expectInfoLog, infoLogged)
			}

			if tt.expectErrorLog && !errorLogged {
				t.Errorf("expected error log but none was logged")
			}
		})
	}
}
