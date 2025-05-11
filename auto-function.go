package main

import (
	"git-auto-commit/types"
	"regexp"
	"strings"
)

var (
	functionRegexGo        = regexp.MustCompile(`func\s+(\w+)\s*\(([^)]*)\)`)
	functionRegexPython    = regexp.MustCompile(`def\s+(\w+)\s*\(([^)]*)\)`)
	functionRegexTSJS      = regexp.MustCompile(`function\s+(\w+)\s*\(([^)]*)\)(:\s*(\w+))?`)
	functionRegexTSJSConst = regexp.MustCompile(`(const|let|var)\s+(\w+)\s*=\s*(?:function)?\s*\(([^)]*)\)\s*=>?`)
	functionRegexCJava     = regexp.MustCompile(`(\w+)\s+(\w+)\s*\(([^)]*)\)`)
	functionRegexCSharp    = regexp.MustCompile(`(public|private|protected|internal)?\s*(static)?\s*(\w+)\s+(\w+)\s*\(([^)]*)\)`)
)

func Test() {}

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
	var addedFuncs, deletedFuncs, renamedFuncs, changedParams, changedTypes []string
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

				if oldFunc != nil {
					deletedFuncs = append(deletedFuncs, oldFunc.Name)
				}

				// if oldFunc != nil {
				// 	builder.Reset()
				// 	builder.WriteString("deleted function ")
				// 	builder.WriteString(oldFunc.Name)
				// 	results = append(results, builder.String())
				// }
			}
		} else if strings.HasPrefix(line, "+") {
			newFunc = ParseToStructureFunction(line[1:], lang)

			if oldFunc == nil && newFunc != nil {
				addedFuncs = append(addedFuncs, newFunc.Name)
				// builder.Reset()
				// builder.WriteString("added function ")
				// builder.WriteString(newFunc.Name)
				// results = append(results, builder.String())
			}
		} else {
			oldFunc, newFunc = nil, nil
			continue
		}

		if oldFunc != nil && newFunc != nil {
			if oldFunc.Name != newFunc.Name {
				renamedFuncs = append(renamedFuncs, oldFunc.Name+" -> "+newFunc.Name)
				// builder.Reset()
				// builder.WriteString("renamed function ")
				// builder.WriteString(oldFunc.Name)
				// builder.WriteString(" -> ")
				// builder.WriteString(newFunc.Name)
				// results = append(results, builder.String())
			}

			if len(oldFunc.Params) == len(newFunc.Params) {
				for i := range oldFunc.Params {
					if oldFunc.Params[i].Name != newFunc.Params[i].Name && oldFunc.Params[i].Type == newFunc.Params[i].Type {
						changedParams = append(changedParams, oldFunc.Name+" function")
						// builder.Reset()
						// builder.WriteString("changed parameter in ")
						// builder.WriteString(oldFunc.Name)
						// builder.WriteString(" function")
						// results = append(results, builder.String())
					}

					if oldFunc.Params[i].Name == newFunc.Params[i].Name && oldFunc.Params[i].Type != newFunc.Params[i].Type {
						changedTypes = append(changedTypes, oldFunc.Params[i].Name+" in "+oldFunc.Name+" function")
						// builder.Reset()
						// builder.WriteString("changed type ")
						// builder.WriteString(oldFunc.Params[i].Type)
						// builder.WriteString(" -> ")
						// builder.WriteString(newFunc.Params[i].Type)
						// results = append(results, builder.String())
					}
				}
			}

			oldFunc, newFunc = nil, nil
		}
	}

	var results []string
	if len(addedFuncs) == 1 {
		results = append(results, "added function "+addedFuncs[0])
	} else if len(addedFuncs) > 1 {
		results = append(results, "added functions: "+strings.Join(addedFuncs, ", "))
	}

	if len(deletedFuncs) == 1 {
		results = append(results, "deleted function "+deletedFuncs[0])
	} else if len(deletedFuncs) > 1 {
		results = append(results, "deleted functions: "+strings.Join(deletedFuncs, ", "))
	}

	if len(renamedFuncs) == 1 {
		results = append(results, "renamed function "+renamedFuncs[0])
	} else if len(renamedFuncs) > 1 {
		results = append(results, "renamed functions: "+strings.Join(renamedFuncs, ", "))
	}

	if len(changedParams) == 1 {
		results = append(results, "changed parameter in "+changedParams[0])
	} else if len(changedParams) > 1 {
		results = append(results, "changed parameters in functions: "+strings.Join(changedParams, ", "))
	}

	if len(changedTypes) == 1 {
		results = append(results, "changed parameter type "+changedTypes[0])
	} else if len(changedTypes) > 1 {
		results = append(results, "changed parameter types: "+strings.Join(changedTypes, ", "))
	}

	if len(results) == 0 {
		return ""
	}

	return strings.Join(results, " | ")
}

func parseGoFunction(line string) *types.FunctionSignature {
	m := functionRegexGo.FindStringSubmatch(line)
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
	m := functionRegexPython.FindStringSubmatch(line)
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
	m := functionRegexTSJS.FindStringSubmatch(line)
	if m != nil {
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

	m = functionRegexTSJSConst.FindStringSubmatch(line)
	if m != nil {
		name := m[2]

		params := []types.FunctionParameters{}
		for _, p := range strings.Split(m[3], ",") {
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

		return &types.FunctionSignature{Name: name, Params: params, ReturnType: ""}
	}

	return nil
}

func parseCJavaFunction(line string) *types.FunctionSignature {
	m := functionRegexCJava.FindStringSubmatch(line)
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
	m := functionRegexCSharp.FindStringSubmatch(line)
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
