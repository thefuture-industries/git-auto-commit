package code

import (
	"git-auto-commit/infra/constants"
	"os/exec"
	"path/filepath"
	"strings"
)

var ExecCommand = exec.Command

func (c *Code) FormattedCode(files []string) (string, error) {
	args := append([]string{"diff", "--cached", "--name-status"}, files...)

	stdout, err := ExecCommand("git", args...).Output()
	if err != nil {
		return "", err
	}

	var added, modified, deleted []string

	lines := strings.Split(strings.TrimSpace(string(stdout)), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		sep := strings.IndexByte(line, '\t')
		if sep == -1 {
			sep = strings.IndexByte(line, ' ')
			if sep == -1 {
				continue
			}
		}

		status := strings.TrimSpace(line[:sep])
		file := strings.TrimSpace(line[sep+1:])

		switch status {
		case constants.NameStatus_Added:
			added = append(added, file)

		case constants.NameStatus_Modified:
			modified = append(modified, file)

		case constants.NameStatus_Deleted:
			deleted = append(deleted, file)

		}
	}

	msg := c.build(added, modified, deleted)
	return c.WithTag(files, msg, added, modified, deleted), nil
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
	seen := make(map[string]struct{}, len(files))
	names := make([]string, 0, len(files))

	for _, file := range files {
		var dir string

		if slash := strings.IndexByte(file, '/'); slash != -1 {
			dir = file[:slash]
		} else {
			dir = strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
		}

		if dir == "." || dir == "" {
			dir = strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
		}

		if _, exists := seen[dir]; !exists {
			seen[dir] = struct{}{}
			if desc, ok := constants.EXPANDING_NOTATION_FOLDERS[dir]; ok {
				names = append(names, desc)
			} else {
				names = append(names, dir)
			}
		}
	}

	switch len(names) {
	case 0:
		return ""

	case 1:
		return names[0]

	case 2:
		return names[0] + " and " + names[1]

	default:
		return strings.Join(names[:len(names)-1], ", ") + " and " + names[len(names)-1]

	}
}
