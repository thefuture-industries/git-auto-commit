package main

import (
	"fmt"
	"git-auto-commit/types"
	"regexp"
	"strings"
)

func ParseToStrcutVariable(line string) *types.Variable {
	var reg = regexp.MustCompile(`^\s*(\w+)\s+(\w+)\s*=\s*([^;]+);`)

	m := reg.FindStringSubmatch(line)
	if m == nil {
		return nil
	}

	return &types.Variable{Type: m[1], Name: m[2], Value: strings.TrimSpace(m[3])}
}

func FormattedVariables(diff string) (string, error) {
	var oldVar, newVar *types.Variable

	lines := strings.Split(diff, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "-") {
			oldVar = ParseToStrcutVariable(line[1:])
		} else if strings.HasPrefix(line, "+") {
			newVar = ParseToStrcutVariable(line[1:])
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
