package main

import (
	"fmt"
	"git-auto-commit/types"
	"regexp"
	"strings"
)

var (
	classRegexTSJS = regexp.MustCompile(`class\s+(\w+)(?:\s+extends\s+(\w+))?`)
	classRegexPython = regexp.MustCompile(`class\\s+(\\w+)(?:\\((\\w+)\\))?:`)
	classRegexCpp = regexp.MustCompile(`class\\s+(\\w+)(?:\\s*:\\s*(public|protected|private)\\s+(\\w+))?`)
	classRegexCSharp = regexp.MustCompile(`(?:public\\s+)?class\\s+(\\w+)(?:\\s*:\\s*(\\w+))?`)
	classRegexGo = regexp.MustCompile(`type\\s+(\\w+)\\s+struct\\s*{`)
	functionRegexGo = regexp.MustCompile(`func\s+(\w+)\s*\(([^)]*)\)`)
	functionRegexPython = regexp.MustCompile(`def\s+(\w+)\s*\(([^)]*)\)(:\s*(\w+))?`)
	functionRegexTSJS = regexp.MustCompile(`function\s+(\w+)\s*\(([^)]*)\)(:\s*(\w+))?`)
	functionRegexCJava = regexp.MustCompile(`(\w+)\s+(\w+)\s*\(([^)]*)\)`)
	functionRegexCSharp = regexp.MustCompile(`(public|private|protected|internal)?\s*(static)?\s*(\w+)\s+(\w+)\s*\(([^)]*)\)`)
)

func ParseToStructureFunction(line, lang string) *types.FunctionSignature {
	switch lang {
	case "go":
		return parseGoFunction(line)
	case "python":
		return parsePythonFunction(line)
	case "typescript", "javascript":
		return parseTSJSFunction(line)
	case "c", "cpp", "java":
		return parseCJavaFunction(line)
	case "csharp":
		return parseCSharpFunction(line)
	default:
		return nil
	}
}

func FormattedFunction(diff, lang string) string {
	var oldFunc, newFunc *types.FunctionSignature

	var results []string

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

				if oldFunc != nil {
					results = append(results, fmt.Sprintf("deleted function %s", oldFunc.Name))
				}
			}
		} else if strings.HasPrefix(line, "+") {
			newFunc = ParseToStructureFunction(line[1:], lang)

			if oldFunc == nil && newFunc != nil {
				results = append(results, fmt.Sprintf("added function %s", newFunc.Name))
			}
		} else {
			oldFunc, newFunc = nil, nil
			continue
		}

		if oldFunc != nil && newFunc != nil {
			if oldFunc.Name != newFunc.Name {
				results = append(results, fmt.Sprintf("renamed function %s -> %s", oldFunc.Name, newFunc.Name))
			}

			if len(oldFunc.Params) == len(newFunc.Params) {
				for i := range oldFunc.Params {
					if oldFunc.Params[i].Name != newFunc.Params[i].Name && oldFunc.Params[i].Type == newFunc.Params[i].Type {
						results = append(results, fmt.Sprintf("changed parameter in %s function", oldFunc.Name))
					}

					if oldFunc.Params[i].Name == newFunc.Params[i].Name && oldFunc.Params[i].Type != newFunc.Params[i].Type {
						results = append(results, fmt.Sprintf("changed type '%s %s' -> '%s %s'", oldFunc.Params[i].Name, oldFunc.Params[i].Type, newFunc.Params[i].Name, newFunc.Params[i].Type))
					}
				}
			}

			oldFunc, newFunc = nil, nil
		}
	}

	return strings.Join(results, ", ")
}

func parseGoFunction(line string) *types.FunctionSignature {
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

func parsePythonFunction(line string) *types.FunctionSignature {
	functionRegex := regexp.MustCompile(`def\s+(\w+)\s*\(([^)]*)\)`)
	m := functionRegex.FindStringSubmatch(line)
	if m == nil {
		return nil
	}

	name := m[1]
	params := []types.FunctionParameters{}

	for _, p := range strings.Split(m[2], ",") {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}

		parts := strings.Split(p, ":")
		if len(parts) == 2 {
			params = append(params, types.FunctionParameters{Name: strings.TrimSpace(parts[0]), Type: strings.TrimSpace(parts[1])})
		} else {
			params = append(params, types.FunctionParameters{Name: p, Type: ""})
		}
	}

	return &types.FunctionSignature{Name: name, Params: params}
}

func parseTSJSFunction(line string) *types.FunctionSignature {
	functionRegex := regexp.MustCompile(`function\s+(\w+)\s*\(([^)]*)\)(:\s*(\w+))?`)
	m := functionRegex.FindStringSubmatch(line)
	if m == nil {
		return nil
	}

	name := m[1]
	returnType := ""
	if len(m) > 4 {
		returnType = m[4]
	}

	params := []types.FunctionParameters{}
	for _, p := range strings.Split(m[2], ",") {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}

		parts := strings.Split(p, ":")
		if len(parts) == 2 {
			params = append(params, types.FunctionParameters{Name: strings.TrimSpace(parts[0]), Type: strings.TrimSpace(parts[1])})
		} else {
			params = append(params, types.FunctionParameters{Name: p, Type: ""})
		}
	}

	return &types.FunctionSignature{Name: name, Params: params, ReturnType: returnType}
}

func parseCJavaFunction(line string) *types.FunctionSignature {
	functionRegex := regexp.MustCompile(`(\w+)\s+(\w+)\s*\(([^)]*)\)`)
	m := functionRegex.FindStringSubmatch(line)
	if m == nil {
		return nil
	}

	returnType := m[1]
	name := m[2]

	params := []types.FunctionParameters{}
	for _, p := range strings.Split(m[3], ",") {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}

		parts := strings.Fields(p)
		if len(parts) == 2 {
			params = append(params, types.FunctionParameters{Name: parts[1], Type: parts[0]})
		} else if len(parts) == 1 {
			params = append(params, types.FunctionParameters{Name: parts[0], Type: ""})
		}
	}

	return &types.FunctionSignature{Name: name, Params: params, ReturnType: returnType}
}

func parseCSharpFunction(line string) *types.FunctionSignature {
	functionRegex := regexp.MustCompile(`(public|private|protected|internal)?\s*(static)?\s*(\w+)\s+(\w+)\s*\(([^)]*)\)`)

	m := functionRegex.FindStringSubmatch(line)
	if m == nil {
		return nil
	}

	name := m[4]
	params := []types.FunctionParameters{}
	for _, p := range strings.Split(m[5], ",") {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}

		parts := strings.Fields(p)
		if len(parts) == 2 {
			params = append(params, types.FunctionParameters{Name: parts[1], Type: parts[0]})
		} else if len(parts) == 1 {
			params = append(params, types.FunctionParameters{Name: parts[0], Type: ""})
		}
	}

	return &types.FunctionSignature{Name: name, Params: params, ReturnType: m[3]}
}
