package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var extLangMap = map[string]string{
	".py":   "python",
	".go":   "golang",
	".js":   "javascript",
	".ts":   "typescript",
	".cpp":  "cpp",
	".c":    "c",
	".java": "java",
}

var funcPatterns = map[string]*regexp.Regexp{
	"python":     regexp.MustCompile(`^\+\s*def\s+(\w+)\s*\(`),
	"golang":     regexp.MustCompile(`^\+\s*func\s+(\w+)\s*\(`),
	"javascript": regexp.MustCompile(`^\+\s*(?:function\s+)?(\w+)\s*\(`),
	"cpp":        regexp.MustCompile(`^\+\s*(?:[\w:<>,\s\*]+)\s+(\w+)\s*\([^)]*\)\s*\{`),
	"c":          regexp.MustCompile(`^\+\s*(?:[\w\s\*]+)\s+(\w+)\s*\([^)]*\)\s*\{`),
	"java":       regexp.MustCompile(`^\+\s*(?:public|private|protected)?\s+[\w<>\[\]]+\s+(\w+)\s*\(`),
}

func detectLang(filename string) string {
	ext := filepath.Ext(filename)
	return extLangMap[ext]
}

func getDiff(file string) string {
	cmd := exec.Command("git", "diff", "--cached", file)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	return out.String()
}

func getStagedFiles() []string {
	cmd := exec.Command("git", "diff", "--cached", "--name-only")
	out, _ := cmd.Output()
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	return lines
}

func extractFunctions(diff, lang string) []string {
	re, ok := funcPatterns[lang]
	if !ok {
		return nil
	}
	lines := strings.Split(diff, "\n")
	funcs := make(map[string]bool)
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) > 1 {
			funcs[match[1]] = true
		}
	}
	names := []string{}
	for name := range funcs {
		names = append(names, name)
	}
	return names
}

func buildCommitMessage(funcs []string) string {
	n := len(funcs)
	if n == 0 {
		return "auto commit"
	}
	if n == 1 {
		return "added " + funcs[0] + " function"
	}
	return "added " + strings.Join(funcs[:n-1], ", ") + " and " + funcs[n-1] + " functions"
}

func main() {
	files := getStagedFiles()
	allFuncs := []string{}
	for _, file := range files {
		lang := detectLang(file)
		if lang == "" {
			continue
		}
		diff := getDiff(file)
		funcs := extractFunctions(diff, lang)
		allFuncs = append(allFuncs, funcs...)
	}

	msg := buildCommitMessage(allFuncs)

	cmd := exec.Command("git", "commit", "-m", msg)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to commit:", err)
	}
}
