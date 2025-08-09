package code

import (
	"git-auto-commit/infra/constants"
	"regexp"
	"strings"
)

var (
	varRegexPython    = regexp.MustCompile(`^\s*(\w+)\s*=\s*(.+)`)
	varRegexTSJS      = regexp.MustCompile(`^\s*(let|const|var)\s+(\w+)(\s*:\s*(\w+))?\s*=\s*(.+);?`)
	varRegexGoFull    = regexp.MustCompile(`^\s*var\s+(\w+)\s+(\w+)\s*=\s*(.+)`)
	varRegexGoShort   = regexp.MustCompile(`^\s*(\w+)\s*:=\s*(.+)`)
	varRegexGoNoValue = regexp.MustCompile(`^\s*var\s+(\w+)\s+(\w+)`)
	defaultVarRegex   = regexp.MustCompile(`^\s*(\w+)\s+(\w+)\s*=\s*([^;]+);`)
)

func ParseToStructureVariable(line, lang string) *VariableSignature {
	switch lang {
	case "python":
		m := varRegexPython.FindStringSubmatch(line)
		if m == nil {
			return nil
		}

		return &VariableSignature{Type: "", Name: m[1], Value: strings.TrimSpace(m[2])}
	case "typescript", "javascript":
		m := varRegexTSJS.FindStringSubmatch(line)
		if m == nil {
			return nil
		}

		typ := ""

		if len(m) > 4 {
			typ = m[4]
		}

		return &VariableSignature{Type: typ, Name: m[2], Value: strings.TrimSpace(m[5])}
	case "go":
		m := varRegexGoFull.FindStringSubmatch(line)
		if m != nil {
			// names := strings.Split(m[1], ",")
			// value := strings.TrimSpace(m[2])
			// return &VariableSignature{Type: "", Name: strings.TrimSpace(names[0]), Value: value}
			return &VariableSignature{Type: m[2], Name: m[1], Value: strings.TrimSpace(m[3])}
		}

		m = varRegexGoShort.FindStringSubmatch(line)
		if m != nil {
			return &VariableSignature{Type: "", Name: m[1], Value: strings.TrimSpace(m[2])}
		}

		m = varRegexGoNoValue.FindStringSubmatch(line)
		if m != nil {
			return &VariableSignature{Type: m[2], Name: m[1], Value: ""}
		}

		return nil
	default:
		m := defaultVarRegex.FindStringSubmatch(line)
		if m == nil {
			return nil
		}

		return &VariableSignature{Type: m[1], Name: m[2], Value: strings.TrimSpace(m[3])}
	}
}

func FormattedVariables(diff, lang string) string {
	var addedVars, renamedVars, changedTypes, changedValues []string
	var oldVar, newVar *VariableSignature

	lines := strings.Split(diff, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "-") {
			oldVar = ParseToStructureVariable(line[1:], lang)
		} else if strings.HasPrefix(line, "+") {
			newVar = ParseToStructureVariable(line[1:], lang)
		}

		if oldVar != nil && newVar != nil {
			if oldVar.Name == newVar.Name && oldVar.Type != newVar.Type {
				changedTypes = append(changedTypes, oldVar.Name+" ("+oldVar.Type+" -> "+newVar.Type+")")
			}

			if oldVar.Type == newVar.Type && oldVar.Value == newVar.Value && oldVar.Name != newVar.Name {
				renamedVars = append(renamedVars, oldVar.Name+" -> "+newVar.Name)
			}

			if oldVar.Name == newVar.Name && oldVar.Type == newVar.Type && oldVar.Value != newVar.Value {
				changedValues = append(changedValues, oldVar.Name)
			}

			oldVar, newVar = nil, nil
		} else if newVar != nil && oldVar == nil {
			addedVars = append(addedVars, newVar.Name)
			newVar = nil
		}
	}

	var results []string
	if len(addedVars) == 1 {
		results = append(results, "added variable "+addedVars[0])
	} else if len(addedVars) > 1 {
		results = append(results, "added variables: "+strings.Join(addedVars, ", "))
	}

	if len(renamedVars) == 1 {
		results = append(results, "renamed variable "+renamedVars[0])
	} else if len(renamedVars) > 1 {
		results = append(results, "renamed variables: "+strings.Join(renamedVars, ", "))
	}

	if len(changedTypes) == 1 {
		results = append(results, "changed type of variable "+changedTypes[0])
	} else if len(changedTypes) > 1 {
		results = append(results, "changed types of variables: "+strings.Join(changedTypes, ", "))
	}

	if len(changedValues) == 1 {
		results = append(results, "changed value in variable "+changedValues[0])
	} else if len(changedValues) > 1 {
		results = append(results, "changed values in variables: "+strings.Join(changedValues, ", "))
	}

	if len(results) == 0 {
		return ""
	}

	parser := strings.Join(results, " | ")
	for len(parser) > int(constants.MAX_COMMIT_LENGTH) && len(results) > 1 {
		results = results[:len(results)-1]
		parser = strings.Join(results, " | ")
	}

	if len(parser) > int(constants.MAX_COMMIT_LENGTH) && len(results) == 1 {
		parser = parser[:int(constants.MAX_COMMIT_LENGTH)]
	}

	return parser
}
