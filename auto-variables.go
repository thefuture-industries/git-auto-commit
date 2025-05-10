package main

import (
	"git-auto-commit/types"
	"regexp"
	"strings"
)

var (
	varRegexPython  = regexp.MustCompile(`^\s*(\w+)\s*=\s*(.+)`)
	varRegexTSJS    = regexp.MustCompile(`^\s*(let|const|var)\s+(\w+)(\s*:\s*(\w+))?\s*=\s*(.+);?`)
	varRegexGo      = regexp.MustCompile(`^\s*([\w\s,]+):=\s*(.+)`)
	defaultVarRegex = regexp.MustCompile(`^\s*(\w+)\s+(\w+)\s*=\s*([^;]+);`)
)

func ParseToStructureVariable(line, lang string) *types.VariableSignature {
	switch lang {
	case "python":
		m := varRegexPython.FindStringSubmatch(line)
		if m == nil {
			return nil
		}

		return &types.VariableSignature{Type: "", Name: m[1], Value: strings.TrimSpace(m[2])}
	case "typescript", "javascript":
		m := varRegexTSJS.FindStringSubmatch(line)
		if m == nil {
			return nil
		}

		typ := ""

		if len(m) > 4 {
			typ = m[4]
		}

		return &types.VariableSignature{Type: typ, Name: m[2], Value: strings.TrimSpace(m[5])}
	case "go":
		m := varRegexGo.FindStringSubmatch(line)
		if m != nil {
			names := strings.Split(m[1], ",")
			value := strings.TrimSpace(m[2])
			return &types.VariableSignature{Type: "", Name: strings.TrimSpace(names[0]), Value: value}
		}

		return nil
	default:
		m := defaultVarRegex.FindStringSubmatch(line)
		if m == nil {
			return nil
		}

		return &types.VariableSignature{Type: m[1], Name: m[2], Value: strings.TrimSpace(m[3])}
	}
}

func FormattedVariables(diff, lang string) string {
	var oldVar, newVar *types.VariableSignature
	var builder strings.Builder
	var results []string

	lines := strings.Split(diff, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "-") {
			oldVar = ParseToStructureVariable(line[1:], lang)
		} else if strings.HasPrefix(line, "+") {
			newVar = ParseToStructureVariable(line[1:], lang)
		}

		if oldVar != nil && newVar != nil {
			if oldVar.Name == newVar.Name && oldVar.Type != newVar.Type {
				// results = append(results, fmt.Sprintf("changed type of variable %s -> %s", oldVar.Type, newVar.Type))
				builder.Reset()
				builder.WriteString("changed type of variable ")
				builder.WriteString(oldVar.Type)
				builder.WriteString(" -> ")
				builder.WriteString(newVar.Type)
				results = append(results, builder.String())
			}

			if oldVar.Type == newVar.Type && oldVar.Value == newVar.Value && oldVar.Name != newVar.Name {
				// results = append(results, fmt.Sprintf("renamed variable %s -> %s", oldVar.Name, newVar.Name))
				builder.Reset()
				builder.WriteString("renamed variable ")
				builder.WriteString(oldVar.Name)
				builder.WriteString(" -> ")
				builder.WriteString(newVar.Name)
				results = append(results, builder.String())
			}

			if oldVar.Name == newVar.Name && oldVar.Type == newVar.Type && oldVar.Value != newVar.Value {
				// results = append(results, fmt.Sprintf("changed value in variable %s", oldVar.Name))
				builder.Reset()
				builder.WriteString("changed value in variable ")
				builder.WriteString(oldVar.Name)
				results = append(results, builder.String())
			}

			oldVar, oldVar = nil, nil
		} else if newVar != nil && oldVar == nil {
			builder.Reset()
			builder.WriteString("added variable ")
			builder.WriteString(newVar.Name)
			if newVar.Type != "" {
				builder.WriteString(" of type ")
				builder.WriteString(newVar.Type)
			}

			results = append(results, builder.String())
			newVar = nil
		}
	}

	if len(results) == 0 {
		return ""
	}

	builder.Reset()
	for i, result := range results {
		if i > 0 {
			builder.WriteString(", ")
		}

		builder.WriteString(result)
	}

	return builder.String()
}
