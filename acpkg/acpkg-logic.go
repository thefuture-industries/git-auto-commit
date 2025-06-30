package main

import (
	"git-auto-commit/constants"
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

func extractIfBlocks(lines []string, lang string, isNew bool) []string {
	var blocks []string
	var ifRegex *regexp.Regexp

	switch lang {
	case "python":
		ifRegex = regexp.MustCompile(`if\s+([^:]+):`)
	case "go":
		ifRegex = regexp.MustCompile(`^\s*if\s+(.*)`)
	case "c", "cpp", "java", "csharp", "typescript", "javascript":
		ifRegex = regexp.MustCompile(`if\s*\(([^)]+)\)`)
	default:
		ifRegex = regexp.MustCompile(`if`)
	}

	for _, line := range lines {
		if (isNew && strings.HasPrefix(line, "+")) || (!isNew && strings.HasPrefix(line, "-")) {
			l := line[1:]
			if m := ifRegex.FindStringSubmatch(l); m != nil {
				expr := "if"
				if len(m) > 1 {
					expr = strings.TrimSpace(m[1])
				}

				blocks = append(blocks, expr)
			}
		}
	}

	return blocks
}

func describeCondition(expr string) string {
	expr = strings.TrimSpace(expr)
	if idx := strings.Index(expr, "{"); idx != -1 {
		expr = strings.TrimSpace(expr[:idx])
	}

	replacements := []struct {
		pattern *regexp.Regexp
		replace string
	}{
		{regexp.MustCompile(`(\w+)\s*==\s*"?(\w+)"?`), "if $1 is equal to $2"},
		{regexp.MustCompile(`(\w+)\s*!=\s*"?(\w+)"?`), "if $1 is not equal to $2"},
		{regexp.MustCompile(`(\w+)\s*<\s*(\d+)`), "if $1 is less than $2"},
		{regexp.MustCompile(`(\w+)\s*>\s*(\d+)`), "if $1 is greater than $2"},
	}

	// {regexp.MustCompile(`(.+?)\s*==\s*"?(.+?)"?$`), "if $1 is equal to $2"},
	// 	{regexp.MustCompile(`(.+?)\s*!=\s*"?(.+?)"?$`), "if $1 is not equal to $2"},
	// 	{regexp.MustCompile(`(.+?)\s*<\s*(.+?)$`), "if $1 is less than $2"},
	// 	{regexp.MustCompile(`(.+?)\s*>\s*(.+?)$`), "if $1 is greater than $2"},

	for _, r := range replacements {
		if r.pattern.MatchString(expr) {
			return r.pattern.ReplaceAllString(expr, r.replace)
		}
	}

	expr = strings.ReplaceAll(expr, "&&", " and ")
	expr = strings.ReplaceAll(expr, "||", " or ")

	return "added condition logic: " + expr
}

func FormattedLogic(line, lang, filename string) string {
	lines := strings.Split(line, "\n")
	var builder strings.Builder

	oldSwitches := extractSwitchBlocks(lines, lang, false)
	newSwitches := extractSwitchBlocks(lines, lang, true)

	oldIfs := extractIfBlocks(lines, lang, false)
	newIfs := extractIfBlocks(lines, lang, true)

	if len(newIfs) > 0 && len(oldIfs) == 0 {
		builder.Reset()

		for i, cond := range newIfs {
			if i > 0 {
				builder.WriteString("; ")
			}

			builder.WriteString(describeCondition(cond))
		}

		result := builder.String()
		for len(result) > int(constants.MAX_COMMIT_LENGTH) && len(newIfs) > 1 {
			newIfs = newIfs[:len(newIfs)-1]
			builder.Reset()

			for i, cond := range newIfs {
				if i > 0 {
					builder.WriteString("; ")
				}

				builder.WriteString(describeCondition(cond))
			}

			result = builder.String()
		}

		if len(result) > int(constants.MAX_COMMIT_LENGTH) && len(newIfs) == 1 {
			result = result[:int(constants.MAX_COMMIT_LENGTH)]
		}

		return result
	}

	if len(oldSwitches) > 0 && len(newSwitches) == 0 {
		makeResult := func(switches []types.SwitchSignature) string {
			var b strings.Builder
			b.WriteString("removed switch on ")
			for i, sw := range switches {
				if i > 0 {
					b.WriteString("; ")
				}
				b.WriteString("'" + sw.Expr + "'")
				if len(sw.Cases) > 0 {
					b.WriteString(" with cases: ")
					b.WriteString(strings.ReplaceAll(strings.Join(sw.Cases, ", "), "\"", "'"))
				}
			}

			return b.String()
		}

		result := makeResult(oldSwitches)
		for len(result) > int(constants.MAX_COMMIT_LENGTH) && len(oldSwitches) > 1 {
			oldSwitches = oldSwitches[:len(oldSwitches)-1]
			result = makeResult(oldSwitches)
		}

		if len(result) > int(constants.MAX_COMMIT_LENGTH) && len(oldSwitches) == 1 {
			result = result[:int(constants.MAX_COMMIT_LENGTH)]
		}

		return result
	}

	if len(newSwitches) > 0 && len(oldSwitches) == 0 {
		makeResult := func(switches []types.SwitchSignature) string {
			var b strings.Builder
			b.WriteString("added switch on ")
			for i, sw := range switches {
				if i > 0 {
					b.WriteString("; ")
				}
				b.WriteString("'" + sw.Expr + "'")
				if len(sw.Cases) > 0 {
					b.WriteString(" with cases: ")
					b.WriteString(strings.ReplaceAll(strings.Join(sw.Cases, ", "), "\"", "'"))
				}
			}

			return b.String()
		}

		result := makeResult(newSwitches)
		for len(result) > int(constants.MAX_COMMIT_LENGTH) && len(newSwitches) > 1 {
			newSwitches = newSwitches[:len(newSwitches)-1]
			result = makeResult(newSwitches)
		}

		if len(result) > int(constants.MAX_COMMIT_LENGTH) && len(newSwitches) == 1 {
			result = result[:int(constants.MAX_COMMIT_LENGTH)]
		}

		return result
	}

	if len(oldSwitches) > 0 && len(newSwitches) > 0 {
		osw := oldSwitches[0]
		nsw := newSwitches[0]

		makeResult := func(osw, nsw types.SwitchSignature) string {
			var b strings.Builder
			b.WriteString("changed switch from '")
			b.WriteString(osw.Expr)
			b.WriteString("' (cases: ")
			b.WriteString(strings.Join(osw.Cases, ", "))
			b.WriteString(") to '")
			b.WriteString(nsw.Expr)
			b.WriteString("' (cases: ")
			b.WriteString(strings.ReplaceAll(strings.Join(nsw.Cases, ", "), "\"", "'"))
			b.WriteString(")")

			return b.String()
		}

		result := makeResult(osw, nsw)
		for len(result) > int(constants.MAX_COMMIT_LENGTH) && (len(osw.Cases) > 1 || len(nsw.Cases) > 1) {
			if len(osw.Cases) > len(nsw.Cases) {
				osw.Cases = osw.Cases[:len(osw.Cases)-1]
			} else {
				nsw.Cases = nsw.Cases[:len(nsw.Cases)-1]
			}

			result = makeResult(osw, nsw)
		}

		if len(result) > int(constants.MAX_COMMIT_LENGTH) && len(osw.Cases) == 1 && len(nsw.Cases) == 1 {
			result = result[:int(constants.MAX_COMMIT_LENGTH)]
		}

		return result
	}

	return ""
}
