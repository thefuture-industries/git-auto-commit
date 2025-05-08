package main

import (
	"fmt"
	"git-auto-commit/types"
	"regexp"
	"strings"
)

func ParseToStrcutVariable(line, lang string) *types.Variable {
	switch lang {
	case "python":
		reg := regexp.MustCompile(`^\s*(\w+)\s*=\s*(.+)`)
		m := reg.FindStringSubmatch(line)
		if m == nil {
			return nil
		}

		return &types.Variable{Type: "", Name: m[1], Value: strings.TrimSpace(m[2])}
	case "typescript", "javascript":
		reg := regexp.MustCompile(`^\s*(let|const|var)\s+(\w+)(\s*:\s*(\w+))?\s*=\s*(.+);?`)
		m := reg.FindStringSubmatch(line)
		if m == nil {
			return nil
		}

		typ := ""

		if len(m) > 4 {
			typ = m[4]
		}

		return &types.Variable{Type: typ, Name: m[2], Value: strings.TrimSpace(m[5])}
	case "go":
		reg := regexp.MustCompile(`^\s*([\w\s,]+):=\s*(.+)`)
		m := reg.FindStringSubmatch(line)
		if m != nil {
			names := strings.Split(m[1], ",")
			value := strings.TrimSpace(m[2])
			return &types.Variable{Type: "", Name: strings.TrimSpace(names[0]), Value: value}
		}

		return nil
	default:
		reg := regexp.MustCompile(`^\s*(\w+)\s+(\w+)\s*=\s*([^;]+);`)

		m := reg.FindStringSubmatch(line)
		if m == nil {
			return nil
		}

		return &types.Variable{Type: m[1], Name: m[2], Value: strings.TrimSpace(m[3])}
	}
}

func FormattedVariables(diff, lang string) (string, error) {
	var oldVar, newVar *types.Variable

	lines := strings.Split(diff, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "-") {
			oldVar = ParseToStrcutVariable(line[1:], lang)
		} else if strings.HasPrefix(line, "+") {
			newVar = ParseToStrcutVariable(line[1:], lang)
		}

		if oldVar != nil && newVar != nil {
			if oldVar.Name == newVar.Name && oldVar.Type != newVar.Type {
				return fmt.Sprintf("changed %s -> %s", oldVar.Type, newVar.Type), nil
			}

			if oldVar.Type == newVar.Type && oldVar.Value == newVar.Value && oldVar.Name != newVar.Name {
				return fmt.Sprintf("changed %s -> %s", oldVar.Name, newVar.Name), nil
			}

			if oldVar.Name == newVar.Name && oldVar.Type == newVar.Type && oldVar.Value != newVar.Value {
				return fmt.Sprintf("changed value in %s", oldVar.Name), nil
			}

			oldVar, oldVar = nil, nil
		}
	}

	return "", nil
}
