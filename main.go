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
	if len(files) == 0 {
		fmt.Println("No files staged for commit.")
		return
	}

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

	commit, err := pkg.BuildCommitMessage(allFuncs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("[git-auto-commit] commit is: %s\n", commit)

	if err := pkg.Commit(commit); err != nil {
		fmt.Println(err.Error())
	}
}
