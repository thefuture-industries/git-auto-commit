package main

import (
	"fmt"
	"git-auto-commit/types"
	"regexp"
	"strings"
)

func ParseToStructureVariable(line, lang string) *types.VariableSignature {
	switch lang {
	case "python":
		reg := regexp.MustCompile(`^\s*(\w+)\s*=\s*(.+)`)
		m := reg.FindStringSubmatch(line)
		if m == nil {
			return nil
		}

		return &types.VariableSignature{Type: "", Name: m[1], Value: strings.TrimSpace(m[2])}
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

		return &types.VariableSignature{Type: typ, Name: m[2], Value: strings.TrimSpace(m[5])}
	case "go":
		reg := regexp.MustCompile(`^\s*([\w\s,]+):=\s*(.+)`)
		m := reg.FindStringSubmatch(line)
		if m != nil {
			names := strings.Split(m[1], ",")
			value := strings.TrimSpace(m[2])
			return &types.VariableSignature{Type: "", Name: strings.TrimSpace(names[0]), Value: value}
		}

		return nil
	default:
		reg := regexp.MustCompile(`^\s*(\w+)\s+(\w+)\s*=\s*([^;]+);`)

		m := reg.FindStringSubmatch(line)
		if m == nil {
			return nil
		}

		return &types.VariableSignature{Type: m[1], Name: m[2], Value: strings.TrimSpace(m[3])}
	}
}

func FormattedVariables(diff, lang string) string {
	var oldVar, newVar *types.VariableSignature

	lines := strings.Split(diff, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "-") {
			oldVar = ParseToStructureVariable(line[1:], lang)
		} else if strings.HasPrefix(line, "+") {
			newVar = ParseToStructureVariable(line[1:], lang)
		}

		if oldVar != nil && newVar != nil {
			if oldVar.Name == newVar.Name && oldVar.Type != newVar.Type {
				return fmt.Sprintf("changed %s -> %s", oldVar.Type, newVar.Type)
			}

			if oldVar.Type == newVar.Type && oldVar.Value == newVar.Value && oldVar.Name != newVar.Name {
				return fmt.Sprintf("renamed %s -> %s", oldVar.Name, newVar.Name)
			}

			if oldVar.Name == newVar.Name && oldVar.Type == newVar.Type && oldVar.Value != newVar.Value {
				return fmt.Sprintf("changed value in %s", oldVar.Name)
			}

			oldVar, oldVar = nil, nil
		}
	}

	return ""
}
