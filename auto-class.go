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
	classRegexJava = regexp.MustCompile(`(?:public\\s+)?class\\s+(\\w+)(?:\\s+extends\\s+(\\w+))?`)
)

func ParseToStructureClass(line, lang string) *types.ClassSignature {
	switch lang {
	case "typescript", "javascript":
		return parseTSJSClass(line)
	case "python":
		return parsePythonClass(line)
	case "cpp":
		return parseCppClass(line)
	case "csharp":
		return parseCSharpClass(line)
	case "go":
		return parseGoStruct(line)
	case "java":
		return parseJavaClass(line)
	default:
		return nil
	}
}

func parseTSJSClass(line string) *types.ClassSignature {
	m := classRegexTSJS.FindStringSubmatch(line)

	name := m[1]
	parent := ""
	if len(m) > 2 {
		parent = m[2]
	}

	methods := parseAccessModifiers(line, "(public|private|protected)\\s+(\\w+)\\s*\\(")
	return &types.ClassSignature{Name: name, Parent: parent, Methods: methods}
}

func parsePythonClass(line string) *types.ClassSignature {
	classRegex := regexp.MustCompile(`class\\s+(\\w+)(?:\\((\\w+)\\))?:`)
	m := classRegex.FindStringSubmatch(line)
	if m == nil {
		return nil
	}

	name := m[1]
	parent := ""
	if len(m) > 2 {
		parent = m[2]
	}

	methods := make(map[string]string)
	methodRegex := regexp.MustCompile(`def\s+(_{0,2}\w+)\s*\(`)
	for _, l := range strings.Split(line, "\n") {
		mm := methodRegex.FindStringSubmatch(l)
		if mm != nil {
			mod := "public"
			if strings.HasPrefix(mm[1], "__") {
				mod = "private"
			} else if strings.HasPrefix(mm[1], "_") {
				mod = "protected"
			}

			methods[mm[1]] = mod
		}
	}

	return &types.ClassSignature{Name: name, Parent: parent, Methods: methods}
}

func parseCppClass(line string) *types.ClassSignature {
	m := classRegexCpp.FindStringSubmatch(line)
	if m == nil {
		return nil
	}

	name := m[1]
	parent := ""
	if len(m) > 3 {
		parent = m[3]
	}

	methods := parseAccessModifiers(line, "(public|private|protected):\\s*\\w+\\s+(\\w+)\\s*\\(")
	return &types.ClassSignature{Name: name, Parent: parent, Methods: methods}
}

func parseCSharpClass(line string) *types.ClassSignature {
	m := classRegexCSharp.FindStringSubmatch(line)
	if m == nil {
		return nil
	}

	name := m[1]
	parent := ""
	if len(m) > 2 {
		parent = m[2]
	}

	methods := parseAccessModifiers(line, "(public|private|protected|internal)\\s+\\w+\\s+(\\w+)\\s*\\(")
	return &types.ClassSignature{Name: name, Parent: parent, Methods: methods}
}

func parseGoStruct(line string) *types.ClassSignature {
	m := classRegexGo.FindStringSubmatch(line)
	if m == nil {
		return nil
	}

	name := m[1]
	return &types.ClassSignature{Name: name, Parent: "", Methods: make(map[string]string)}
}

func parseJavaClass(line string) *types.ClassSignature {
	m := classRegexJava.FindStringSubmatch(line)
	if m == nil {
		return nil
	}

	name := m[1]
	parent := ""
	if len(m) > 2 {
		parent = m[2]
	}

	methods := parseAccessModifiers(line, "(public|private|protected)\\s+(\\w+)\\s*\\(")
	return &types.ClassSignature{Name: name, Parent: parent, Methods: methods}
}

func parseAccessModifiers(line, regex string) map[string]string {
	methods := make(map[string]string)
	methodRegex := regexp.MustCompile(regex)

	for _, l := range strings.Split(line, "\n") {
		mm := methodRegex.FindStringSubmatch(l)
		if mm != nil {
			methods[mm[2]] = mm[1]
		}
	}

	return methods
}

func FormattedClass(diff, lang string) string {
	var oldClass, newClass *types.ClassSignature
	var oldLines, newLines []string

	var results []string

	lines := strings.Split(diff, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "-") {
			oldLines = append(oldLines, line[1:])
		} else if strings.HasPrefix(line, "+") {
			newLines = append(newLines, line[1:])
		}
	}

	oldClass = ParseToStructureClass(strings.Join(oldLines, "\n"), lang)
	newClass = ParseToStructureClass(strings.Join(newLines, "\n"), lang)

	if oldClass != nil && newClass == nil {
		results = append(results, fmt.Sprintf("deleted class %s", oldClass.Name))
	}

	if oldClass != nil && newClass != nil {
		if oldClass.Name != newClass.Name {
			results = append(results, fmt.Sprintf("renamed class %s -> %s", oldClass.Name, newClass.Name))
		}

		if oldClass.Parent != newClass.Parent {
			results = append(results, fmt.Sprintf("the heir was changed to %s", oldClass.Name))
		}

		for m, oldMod := range oldClass.Methods {
			if newMod, ok := newClass.Methods[m]; ok && oldMod != newMod {
				results = append(results, fmt.Sprintf("the access modifier of the %s method has been changed", m))
			}
		}
	}

	return strings.Join(results, ", ")
}
