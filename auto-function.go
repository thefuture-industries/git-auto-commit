package main

import (
	"fmt"
	"git-auto-commit/types"
	"regexp"
	"strings"
)

func ParseToStructureFunction(line, lang string) *types.FunctionSignature {
	functionRegex := regexp.MustCompile(`func\s+(\w+)\s*\(([^)]*)\)`)
	m := functionRegex.FindStringSubmatch(line)
	if m == nil {
		return nil
	}

	name := m[1]
	paramsString := m[2]
	params := []types.FunctionParameters{}

	for _, p := range strings.Split(paramsString, ",") {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}

		parts := strings.Fields(p)
		if len(parts) == 2 {
			params = append(params, types.FunctionParameters{Name: parts[0], Type: parts[1]})
		}
	}

	return &types.FunctionSignature{Name: name, Params: params}
}

func FormattedFunction(diff, lang string) string {
	var oldFunc, newFunc *types.FunctionSignature

	lines := strings.Split(diff, "\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if strings.HasPrefix(line, "-") {
			oldFunc = ParseToStructureFunction(line[1:], lang)

			if i+1 < len(lines) && strings.HasPrefix(lines[i+1], "+") {
				newFunc = ParseToStructureFunction(lines[i+1][1:], lang)
				i++
			} else {
				newFunc = nil
			}
		} else if strings.HasPrefix(line, "+") {
			newFunc = ParseToStructureFunction(line[1:], lang)

			if oldFunc == nil && newFunc != nil {
				return fmt.Sprintf("added function %s", newFunc.Name)
			}
		} else {
			oldFunc, newFunc = nil, nil
			continue
		}

		if oldFunc != nil && newFunc != nil {
			if oldFunc.Name != newFunc.Name {
				return fmt.Sprintf("renamed function %s -> %s", oldFunc.Name, newFunc.Name)
			}

			if len(oldFunc.Params) == len(newFunc.Params) {
				for i := range oldFunc.Params {
					if oldFunc.Params[i].Name != newFunc.Params[i].Name && oldFunc.Params[i].Type == newFunc.Params[i].Type {
						return fmt.Sprintf("changed parameter in %s function", oldFunc.Name)
					}

					if oldFunc.Params[i].Name == newFunc.Params[i].Name && oldFunc.Params[i].Type != newFunc.Params[i].Type {
						return fmt.Sprintf("changed type '%s %s' -> '%s %s'", oldFunc.Params[i].Name, oldFunc.Params[i].Type, newFunc.Params[i].Name, newFunc.Params[i].Type)
					}
				}
			}

			oldFunc, newFunc = nil, nil
		}
	}

	return ""
}

func Test(name string) {}
