package main

import (
	"fmt"
	"git-auto-commit/types"
	"regexp"
	"strings"
)

func extractSwitchBlocks(lines []string, lang string, isNew bool) []types.SwitchSignature {
	var blocks []types.SwitchSignature
	var switchRegex, caseRegex *regexp.Regexp

	switch lang {
	case "python":
		switchRegex = regexp.MustCompile(`match\s+([\w\d_]+)\s*:`)
		caseRegex = regexp.MustCompile(`case\s+([^:]+):`)
	case "go":
		switchRegex = regexp.MustCompile(`switch\s*([\w\d_]+)?\s*{`)
		caseRegex = regexp.MustCompile(`case\s+([^:]+):`)
	case "c", "cpp", "java", "csharp":
		switchRegex = regexp.MustCompile(`switch\s*\(([^)]+)\)`)
		caseRegex = regexp.MustCompile(`case\s+([^:]+):`)
	case "typescript", "javascript":
		switchRegex = regexp.MustCompile(`switch\s*\(([^)]+)\)`)
		caseRegex = regexp.MustCompile(`case\s+([^:]+):`)
	default:
		switchRegex = regexp.MustCompile(`switch`)
		caseRegex = regexp.MustCompile(`case`)
	}

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if (isNew && strings.HasPrefix(line, "+")) || (!isNew && strings.HasPrefix(line, "-")) {
			l := line[1:]
			if m := switchRegex.FindStringSubmatch(l); m != nil {
				expr := "switch"

				if len(m) > 1 && m[1] != "" {
					expr = strings.TrimSpace(m[1])
				}

				cases := []string{}
				for j := i + 1; j < len(lines); j++ {
					cl := lines[j]
					if (isNew && strings.HasPrefix(cl, "+")) || (!isNew && strings.HasPrefix(cl, "-")) {
						cln := cl[1:]
						if caseRegex.MatchString(cln) {
							cm := caseRegex.FindStringSubmatch(cln)
							if len(cm) > 1 {
								cases = append(cases, strings.TrimSpace(cm[1]))
							}
						}

						if switchRegex.MatchString(cln) {
							break
						}
					} else {
						break
					}
				}

				blocks = append(blocks, types.SwitchSignature{Expr: expr, Cases: cases})
			}
		}
	}

	return blocks
}

func FormattedLogic(line, lang string) string {
	lines := strings.Split(line, "\n")
	oldSwitches := extractSwitchBlocks(lines, lang, false)
	newSwitches := extractSwitchBlocks(lines, lang, true)

	if len(oldSwitches) > 0 && len(newSwitches) == 0 {
		var parts []string
		for _, sw := range oldSwitches {
			cases := ""

			if len(sw.Cases) > 0 {
				cases = fmt.Sprintf(" (cases: %s)", strings.Join(sw.Cases, ", "))
			}

			parts = append(parts, fmt.Sprintf("%s%s", sw.Expr, cases))
		}

		return fmt.Sprintf("deleted switch: %s", strings.Join(parts, "; "))
	}

	if len(newSwitches) > 0 && len(oldSwitches) == 0 {
		var parts []string
		for _, sw := range newSwitches {
			cases := ""

			if len(sw.Cases) > 0 {
				cases = fmt.Sprintf(" (cases: %s)", strings.Join(sw.Cases, ", "))
			}

			parts = append(parts, fmt.Sprintf("%s%s", sw.Expr, cases))
		}

		return fmt.Sprintf("added switch: %s", strings.Join(parts, "; "))
	}

	if len(oldSwitches) > 0 && len(newSwitches) > 0 {
		osw := oldSwitches[0]
		nsw := newSwitches[0]
		if osw.Expr != nsw.Expr || strings.Join(osw.Cases, ",") != strings.Join(nsw.Cases, ",") {
			return fmt.Sprintf("changed logic switch '%s (cases: %s)' -> '%s (cases: %s)'", osw.Expr, strings.Join(osw.Cases, ", "), nsw.Expr, strings.Join(nsw.Cases, ", "))
		}
	}

	return ""
}
