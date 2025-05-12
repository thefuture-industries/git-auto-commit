package main

import (
	"regexp"
	"strings"
)

func FormattedImport(diff, lang, filename string) string {
	var importRegex *regexp.Regexp

	var imports []string

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
						imports = append(imports, importName)
					}
				} else if strings.HasPrefix(trimmed, "import ") {
					if m := regexp.MustCompile(`^import\s+\"([^\"]+)\"`).FindStringSubmatch(trimmed); m != nil {
						imports = append(imports, m[1])
					}
				}
			}
		}

		if len(imports) == 1 {
			return "included '" + imports[0] + "' in " + filename
		} else if len(imports) > 1 {
			quoted := make([]string, len(imports))
			for i, imp := range imports {
				quoted[i] = "'" + imp + "'"
			}

			return "included " + strings.Join(quoted, ", ") + " in " + filename
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
				imports = append(imports, m[1])
			}
		}
	}

	if len(imports) == 1 {
		return "included '" + imports[0] + "' in " + filename
	} else if len(imports) > 1 {
		quoted := make([]string, len(imports))
		for i, imp := range imports {
			quoted[i] = "'" + imp + "'"
		}

		result := "included " + strings.Join(quoted, ", ") + " in " + filename
		for len(result) > int(MAX_COMMIT_LENGTH) && len(quoted) > 1 {
			quoted = quoted[:len(quoted)-1]
			result = "included " + strings.Join(quoted, ", ") + " in " + filename
		}

		if len(result) > int(MAX_COMMIT_LENGTH) && len(quoted) == 1 {
			result = result[:int(MAX_COMMIT_LENGTH)]
		}

		return result
	}

	return ""
}
