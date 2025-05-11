package main

import (
	"regexp"
	"strings"
)

func FormattedImport(diff, lang, filename string) string {
	var importRegex *regexp.Regexp
	var builder strings.Builder

	lines := strings.Split(diff, "\n")

	switch lang {
	case "python":
		importRegex = regexp.MustCompile(`^import\s+(\w+)`)
	case "go":
		inImportBlock := false
		for _, line := range lines {
			if strings.HasPrefix(line, "+") && len(line) > 1 {
				trimmed := strings.TrimSpace(line[1:])

				if strings.HasPrefix(trimmed, "import (") {
					inImportBlock = true
					continue
				}

				if inImportBlock {
					if strings.HasPrefix(trimmed, ")") {
						inImportBlock = false
						continue
					}

					if strings.HasPrefix(trimmed, "\"") && strings.HasSuffix(trimmed, "\"") {
						importName := strings.Trim(trimmed, "\"")
						builder.Reset()
						builder.WriteString("include '")
						builder.WriteString(importName)
						builder.WriteString("' in ")
						builder.WriteString(filename)
						return builder.String()
					}

					// if m := importRegex.FindStringSubmatch(trimmed); m != nil {
					// 	builder.Reset()
					// 	builder.WriteString("include ")
					// 	builder.WriteString("'")
					// 	builder.WriteString(m[1])
					// 	builder.WriteString("'")
					// 	builder.WriteString(" in ")
					// 	builder.WriteString(filename)
					// 	return builder.String()
					// }
				} else if strings.HasPrefix(trimmed, "import ") {
					if m := regexp.MustCompile(`^import\s+\"([^\"]+)\"`).FindStringSubmatch(trimmed); m != nil {
						builder.Reset()
						builder.WriteString("include '")
						builder.WriteString(m[1])
						builder.WriteString("' in ")
						builder.WriteString(filename)
						return builder.String()
					}
				}
			}
		}

		return ""
	case "javascript", "typescript":
		importRegex = regexp.MustCompile(`^import\s+.*from\s+['\"]([^'\"]+)['\"]`)
	case "java":
		importRegex = regexp.MustCompile(`^import\s+([\w\.]+);`)
	case "c", "cpp":
		importRegex = regexp.MustCompile(`^#include\s+[<\"]([^>\"]+)[>\"]`)
	case "csharp":
		importRegex = regexp.MustCompile(`^using\s+([\w\.]+);`)
	default:
		return ""
	}

	for _, line := range lines {
		if strings.HasPrefix(line, "+") {
			l := line[1:]
			if m := importRegex.FindStringSubmatch(strings.TrimSpace(l)); m != nil {
				builder.Reset()
				builder.WriteString("include ")
				builder.WriteString("'")
				builder.WriteString(m[1])
				builder.WriteString("'")
				builder.WriteString(" in ")
				builder.WriteString(filename)
				return builder.String()
			}
		}
	}

	return ""
}
