package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var logFile *os.File

func init() {
	tmpDir := os.TempDir()
	logPath := filepath.Join(tmpDir, "autocommit.log")

	var err error
	logFile, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return
	}
}

func logMessage(prefix, msg, color string) {
	var builder strings.Builder
	builder.WriteString(color)
	builder.WriteString("[git auto-commit] ")
	builder.WriteString(msg)
	builder.WriteString("\033[0m\n")

	fmt.Print(builder.String())

	if logFile != nil {
		fileMsg := fmt.Sprintf("[git auto-commit] %s\n", msg)
		logFile.WriteString(fileMsg)
	}
}

var InfoLogger = func(msg string) {
	logMessage("", msg, "")
}

var GitLogger = func(msg string) {
	logMessage("", msg, "\033[0;34m")
}

var ErrorLogger = func(err error) {
	logMessage("", err.Error(), "\033[0;31m")
}
