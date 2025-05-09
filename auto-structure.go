package main

import (
	"fmt"
	"git-auto-commit/types"
	"regexp"
	"strings"
)

func parseStruct(line, lang string) *types.StructureSignature {
	var structRegex *regexp.Regexp

	switch lang {
	case "go":
		structRegex = regexp.MustCompile(`type\s+(\w+)\s+struct\s*{`)
	case "csharp":
		structRegex = regexp.MustCompile(`struct\s+(\w+)\s*{`)
	case "c", "cpp":
		structRegex = regexp.MustCompile(`struct\s+(\w+)\s*{`)
	case "typescript":
		structRegex = regexp.MustCompile(`interface\s+(\w+)\s*{`)
	default:
		return nil
	}

	m := structRegex.FindStringSubmatch(line)
	if m == nil {
		return nil
	}

	return &types.StructureSignature{Name: m[1]}
}

func parseType(line, lang string) *types.TypeSignature {
	var typeRegex *regexp.Regexp
	switch lang {
	case "go":
		typeRegex = regexp.MustCompile(`type\s+(\w+)\s+`)
	case "typescript":
		typeRegex = regexp.MustCompile(`type\s+(\w+)\s*=`)
	case "csharp":
		typeRegex = regexp.MustCompile(`using\s+(\w+)\s*=`)
	default:
		return nil
	}

	m := typeRegex.FindStringSubmatch(line)
	if m == nil {
		return nil
	}

	return &types.TypeSignature{Name: m[1]}
}

func parseInterface(line, lang string) *types.InterfaceSignature {
	var interfaceRegex *regexp.Regexp
	switch lang {
	case "go":
		interfaceRegex = regexp.MustCompile(`type\s+(\w+)\s+interface\s*{`)
	case "typescript", "java", "csharp":
		interfaceRegex = regexp.MustCompile(`interface\s+(\w+)\s*{`)
	default:
		return nil
	}

	m := interfaceRegex.FindStringSubmatch(line)
	if m == nil {
		return nil
	}

	return &types.InterfaceSignature{Name: m[1]}
}

func parseEnum(line, lang string) *types.EnumSignature {
	var enumRegex *regexp.Regexp
	switch lang {
	case "typescript", "java", "csharp", "cpp", "c":
		enumRegex = regexp.MustCompile(`enum\s+(\w+)\s*{`)
	default:
		return nil
	}

	m := enumRegex.FindStringSubmatch(line)
	if m == nil {
		return nil
	}

	return &types.EnumSignature{Name: m[1]}
}

func FormattedStruct(diff, lang string) string {
	var oldStruct, newStruct *types.StructureSignature
	var oldLines, newLines []string

	lines := strings.Split(diff, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "-") {
			oldLines = append(oldLines, line[1:])
		} else if strings.HasPrefix(line, "+") {
			newLines = append(newLines, line[1:])
		}
	}

	oldStruct = parseStruct(strings.Join(oldLines, "\n"), lang)
	newStruct = parseStruct(strings.Join(newLines, "\n"), lang)

	if oldStruct != nil && newStruct == nil {
		return fmt.Sprintf("deleted struct %s", oldStruct.Name)
	}

	if oldStruct == nil && newStruct != nil {
		return fmt.Sprintf("added struct %s", newStruct.Name)
	}

	if oldStruct != nil && newStruct != nil && oldStruct.Name != newStruct.Name {
		return fmt.Sprintf("renamed struct %s -> %s", oldStruct.Name, newStruct.Name)
	}

	return ""
}

func FormattedType(diff, lang string) string {
	var oldType, newType *types.TypeSignature
	var oldLines, newLines []string

	lines := strings.Split(diff, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "-") {
			oldLines = append(oldLines, line[1:])
		} else if strings.HasPrefix(line, "+") {
			newLines = append(newLines, line[1:])
		}
	}

	oldType = parseType(strings.Join(oldLines, "\n"), lang)
	newType = parseType(strings.Join(newLines, "\n"), lang)

	if oldType != nil && newType == nil {
		return fmt.Sprintf("deleted type %s", oldType.Name)
	}

	if oldType == nil && newType != nil {
		return fmt.Sprintf("added type %s", newType.Name)
	}

	if oldType != nil && newType != nil && oldType.Name != newType.Name {
		return fmt.Sprintf("renamed type %s -> %s", oldType.Name, newType.Name)
	}

	return ""
}

func FormattedInterface(diff, lang string) string {
	var oldInterface, newInterface *types.InterfaceSignature
	var oldLines, newLines []string

	lines := strings.Split(diff, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "-") {
			oldLines = append(oldLines, line[1:])
		} else if strings.HasPrefix(line, "+") {
			newLines = append(newLines, line[1:])
		}
	}

	oldInterface = parseInterface(strings.Join(oldLines, "\n"), lang)
	newInterface = parseInterface(strings.Join(newLines, "\n"), lang)

	if oldInterface != nil && newInterface == nil {
		return fmt.Sprintf("deleted interface %s", oldInterface.Name)
	}

	if oldInterface == nil && newInterface != nil {
		return fmt.Sprintf("added interface %s", newInterface.Name)
	}

	if oldInterface != nil && newInterface != nil && oldInterface.Name != newInterface.Name {
		return fmt.Sprintf("renamed interface %s -> %s", oldInterface.Name, newInterface.Name)
	}

	return ""
}

func FormattedEnum(diff, lang string) string {
	var oldEnum, newEnum *types.EnumSignature
	var oldLines, newLines []string

	lines := strings.Split(diff, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "-") {
			oldLines = append(oldLines, line[1:])
		} else if strings.HasPrefix(line, "+") {
			newLines = append(newLines, line[1:])
		}
	}

	oldEnum = parseEnum(strings.Join(oldLines, "\n"), lang)
	newEnum = parseEnum(strings.Join(newLines, "\n"), lang)

	if oldEnum != nil && newEnum == nil {
		return fmt.Sprintf("deleted enum %s", oldEnum.Name)
	}

	if oldEnum == nil && newEnum != nil {
		return fmt.Sprintf("added enum %s", newEnum.Name)
	}

	if oldEnum != nil && newEnum != nil && oldEnum.Name != newEnum.Name {
		return fmt.Sprintf("renamed enum %s -> %s", oldEnum.Name, newEnum.Name)
	}

	return ""
}
