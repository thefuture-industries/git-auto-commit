package utils

import (
	"git-auto-commit/constants"
	"strings"
)

func ExtractFunctions(diff, lang string) []string {
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
