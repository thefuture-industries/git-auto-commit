package main

import (
	"fmt"
	"git-auto-commit/pkg"
	"git-auto-commit/utils"
	"os/exec"
	"strings"
)

func getStagedFiles() []string {
	cmd := exec.Command("git", "diff", "--cached", "--name-only")
	out, _ := cmd.Output()
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	return lines
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
		funcs := utils.ExtractFunctions(diff, lang)
		allFuncs = append(allFuncs, funcs...)
	}

	msg := pkg.BuildCommitMessage(allFuncs)

	if err := pkg.Commit(msg); err != nil {
		fmt.Println(err.Error())
	}
}
