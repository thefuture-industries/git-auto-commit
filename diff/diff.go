package diff

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
	"sync"

	"git-auto-commit/git"
)

var diffBufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

var GetDiff = func(file string) (string, error) {
	var builder strings.Builder
	builder.Reset()

	root, err := git.GetGitRoot()
	if err != nil {
		return "", err
	}

	builder.WriteString(root)
	builder.WriteString("/")
	builder.WriteString(file)

	cmd := exec.Command("git", "diff", "--cached", "--", builder.String())

	buf := diffBufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer diffBufferPool.Put(buf)

	cmd.Stdout = buf
	if err := cmd.Run(); err != nil {
		return "", err
	}

	return buf.String(), nil
}

var GetStagedFiles = func() ([]string, error) {
	cmd := exec.Command("git", "diff", "--cached", "--name-only")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	var files []string
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		files = append(files, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	return files, nil
}
