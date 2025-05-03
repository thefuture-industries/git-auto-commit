package main

import (
	"fmt"
	"git-auto-commit/constants"
	"git-auto-commit/pkg"
	"os"
	"os/exec"
	"strings"
)

func getStagedFiles() []string {
	cmd := exec.Command("git", "diff", "--cached", "--name-only")
	out, _ := cmd.Output()
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	return lines
}

func extractFunctions(diff, lang string) []string {
	re, ok := constants.FUNC_PATTERNS[lang]
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
	return "added " + strings.Join(funcs[:n-1], ", ") + " and " + funcs[n-1] + " functionality"
}

func main() {
	files := getStagedFiles()
	allFuncs := []string{}
	for _, file := range files {
		lang := pkg.DetectLang(file)
		if lang == "" {
			continue
		}
		diff := pkg.GetDiff(file)
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
