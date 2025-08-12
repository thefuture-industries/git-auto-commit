package code

import (
	"bufio"
	"git-auto-commit/infra/constants"
	"os/exec"
	"path/filepath"
	"strings"
)

var ExecCommand = exec.Command

func (c *Code) FormattedCode(files []string) (string, error) {
	args := append([]string{"diff", "--cached", "--name-status"}, files...)
	cmd := ExecCommand("git", args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	if err := cmd.Start(); err != nil {
		return "", err
	}

	var added, modified, deleted []string

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}

		status, file := parts[0], parts[1]
		switch status {
		case constants.NameStatus_Added:
			added = append(added, file)

		case constants.NameStatus_Modified:
			modified = append(modified, file)

		case constants.NameStatus_Deleted:
			deleted = append(deleted, file)

		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	if err := cmd.Wait(); err != nil {
		return "", err
	}

	msg := c.build(added, modified, deleted)
	return c.WithTag(files, msg), nil
}

func (c *Code) build(added, modified, deleted []string) string {
	var parts []string

	if len(added) > 0 {
		parts = append(parts, "implemented "+c.summarize(added))
	}

	if len(modified) > 0 {
		parts = append(parts, "updated "+c.summarize(modified))
	}

	if len(deleted) > 0 {
		parts = append(parts, "removed "+c.summarize(deleted))
	}

	msg := strings.Join(parts, ", ")
	if len(msg) > 0 {
		msg = strings.ToUpper(msg[:1]) + msg[1:]
	}

	return msg
}

func (c *Code) summarize(files []string) string {
	folders := map[string]struct{}{}
	for _, file := range files {
		directory := strings.Split(filepath.ToSlash(file), "/")[0]

		if directory == "." || directory == "" {
			directory = strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
		}

		folders[directory] = struct{}{}
	}

	names := make([]string, 0, len(folders))
	for name := range folders {
		if desc, ok := constants.EXPANDING_NOTATION_FOLDERS[name]; ok {
			names = append(names, desc)
		} else {
			names = append(names, name)
		}
	}

	if len(names) == 1 {
		return names[0]
	}

	if len(names) == 2 {
		return names[0] + " and " + names[1]
	}

	return strings.Join(names[:len(names)-1], ", ") + " and " + names[len(names)-1]
}
