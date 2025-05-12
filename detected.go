package main

import (
	"path/filepath"
	"strings"
)

func DetectLanguage(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))

	switch ext {
	case ".go":
		return "go"
	case ".py":
		return "python"
	case ".js":
		return "javascript"
	case ".jsx":
		return "javascript"
	case ".ts":
		return "typescript"
	case ".tsx":
		return "typescript"
	case ".cpp":
		return "cpp"
	case ".c", ".h":
		return "c"
	case ".java":
		return "java"
	case ".cs":
		return "csharp"
	case ".rs":
		return "rust"
	case ".scala":
		return "scala"
	default:
		return ""
	}
}
