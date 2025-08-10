package git

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

var diffBufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func (g *Git) GetDiff(file string) (string, error) {
	var builder strings.Builder
	builder.Reset()

	root, err := g.GetGitRoot()
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

func (g *Git) GetStagedCountDirectory() (string, error) {
	cmd := exec.Command("git", "diff", "--cached", "--numstat")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	if err := cmd.Start(); err != nil {
		return "", err
	}

	directoryChanges := make(map[string]int)
	rootFileChanges := make(map[string]int)

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}

		adds, errAdds := strconv.Atoi(fields[0])
		dels, errDels := strconv.Atoi(fields[1])
		if errAdds != nil || errDels != nil {
			continue
		}

		file := fields[2]
		changes := adds + dels

		lastSlash := strings.LastIndex(file, "/")
		if lastSlash != -1 {
			dir := file[:lastSlash]
			directoryChanges[dir] += changes
		} else {
			rootFileChanges[file] += changes
		}
	}

	if err := cmd.Wait(); err != nil {
		return "", err
	}

	var maxDirectory string
	var maxDirectoryChanges int

	for dir, count := range directoryChanges {
		if count > maxDirectoryChanges {
			maxDirectoryChanges = count
			maxDirectory = dir
		}
	}

	var maxRootFile string
	var maxRootChanges int

	for file, count := range rootFileChanges {
		if count > maxRootChanges {
			maxRootChanges = count
			maxRootFile = file
		}
	}

	if maxRootChanges > maxDirectoryChanges {
		return maxRootFile, nil
	}

	if maxDirectory != "" {
		return maxDirectory, nil
	}

	return "", fmt.Errorf("")
}

func (g *Git) GetStagedFiles() ([]string, error) {
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
