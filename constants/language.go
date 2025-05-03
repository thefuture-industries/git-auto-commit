package constants

import "regexp"

var LANGUAGE_MAP = map[string]string{
	".py":    "python",
	".go":    "golang",
	".js":    "javascript",
	".ts":    "typescript",
	".cpp":   "cpp",
	".c":     "c",
	".java":  "java",
	".cs":    "csharp",
	".rs":    "rust",
	".scala": "scala",
}

var FUNC_PATTERNS = map[string]*regexp.Regexp{
	"python":     regexp.MustCompile(`^\+\s*def\s+(\w+)\s*\(`),
	"golang":     regexp.MustCompile(`^\+\s*func\s+(\w+)\s*\(`),
	"javascript": regexp.MustCompile(`^\+\s*(?:function\s+)?(\w+)\s*\(`),
	"cpp":        regexp.MustCompile(`^\+\s*(?:[\w:<>,\s\*]+)\s+(\w+)\s*\([^)]*\)\s*\{`),
	"c":          regexp.MustCompile(`^\+\s*(?:[\w\s\*]+)\s+(\w+)\s*\([^)]*\)\s*\{`),
	"java":       regexp.MustCompile(`^\+\s*(?:public|private|protected)?\s+[\w<>\[\]]+\s+(\w+)\s*\(`),
	"csharp":     regexp.MustCompile(`^\+\s*(?:public|private|protected|internal|protected internal)?\s+[\w<>,\s\*]+\s+(\w+)\s*\(`),
	"rust":       regexp.MustCompile(`^\+\s*fn\s+(\w+)\s*\(`),
	"scala":      regexp.MustCompile(`^\+\s*def\s+(\w+)\s*\(`),
}
