package acpkg

import (
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
	var addedStruct, deletedStruct, renamedStruct []string
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
		deletedStruct = append(deletedStruct, oldStruct.Name)
	}

	if oldStruct == nil && newStruct != nil {
		addedStruct = append(addedStruct, newStruct.Name)
	}

	if oldStruct != nil && newStruct != nil && oldStruct.Name != newStruct.Name {
		renamedStruct = append(renamedStruct, oldStruct.Name+" -> "+newStruct.Name)
	}

	if len(addedStruct) == 1 {
		return "added structure " + addedStruct[0]
	} else if len(addedStruct) > 1 {
		quoted := make([]string, len(addedStruct))
		for i, structName := range addedStruct {
			quoted[i] = "'" + structName + "'"
		}

		return "added structs: " + strings.Join(quoted, ", ")
	}

	if len(deletedStruct) == 1 {
		return "deleted structure " + deletedStruct[0]
	} else if len(deletedStruct) > 1 {
		quoted := make([]string, len(deletedStruct))
		for i, structName := range deletedStruct {
			quoted[i] = "'" + structName + "'"
		}

		return "deleted structs: " + strings.Join(quoted, ", ")
	}

	if len(renamedStruct) == 1 {
		return "renamed structure " + renamedStruct[0]
	} else if len(renamedStruct) > 1 {
		quoted := make([]string, len(renamedStruct))
		for i, structName := range renamedStruct {
			quoted[i] = "'" + structName + "'"
		}

		return "renamed structs: " + strings.Join(quoted, ", ")
	}

	return ""
}

func FormattedType(diff, lang string) string {
	var oldType, newType *types.TypeSignature
	var addedType, deletedType, renamedType []string
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
		deletedType = append(deletedType, oldType.Name)
	}

	if oldType == nil && newType != nil {
		addedType = append(addedType, newType.Name)
	}

	if oldType != nil && newType != nil && oldType.Name != newType.Name {
		renamedType = append(renamedType, oldType.Name+" -> "+newType.Name)
	}

	if len(addedType) == 1 {
		return "added type " + addedType[0]
	} else if len(addedType) > 1 {
		quoted := make([]string, len(addedType))
		for i, typeName := range addedType {
			quoted[i] = "'" + typeName + "'"
		}

		return "added types: " + strings.Join(quoted, ", ")
	}

	if len(deletedType) == 1 {
		return "deleted type " + deletedType[0]
	} else if len(deletedType) > 1 {
		quoted := make([]string, len(deletedType))
		for i, typeName := range deletedType {
			quoted[i] = "'" + typeName + "'"
		}

		return "deleted types: " + strings.Join(quoted, ", ")
	}

	if len(renamedType) == 1 {
		return "renamed type " + renamedType[0]
	} else if len(renamedType) > 1 {
		quoted := make([]string, len(renamedType))
		for i, typeName := range renamedType {
			quoted[i] = "'" + typeName + "'"
		}

		return "renamed types: " + strings.Join(quoted, ", ")
	}

	return ""
}

func FormattedInterface(diff, lang string) string {
	var oldInterface, newInterface *types.InterfaceSignature
	var addedInterface, deletedInterface, renamedInterface []string
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
		deletedInterface = append(deletedInterface, oldInterface.Name)
	}

	if oldInterface == nil && newInterface != nil {
		addedInterface = append(addedInterface, newInterface.Name)
	}

	if oldInterface != nil && newInterface != nil && oldInterface.Name != newInterface.Name {
		renamedInterface = append(renamedInterface, oldInterface.Name+" -> "+newInterface.Name)
	}

	if len(addedInterface) == 1 {
		return "added interface " + addedInterface[0]
	} else if len(addedInterface) > 1 {
		quoted := make([]string, len(addedInterface))
		for i, interfaceName := range addedInterface {
			quoted[i] = "'" + interfaceName + "'"
		}

		return "added interfaces: " + strings.Join(quoted, ", ")
	}

	if len(deletedInterface) == 1 {
		return "deleted interface " + deletedInterface[0]
	} else if len(deletedInterface) > 1 {
		quoted := make([]string, len(deletedInterface))
		for i, interfaceName := range deletedInterface {
			quoted[i] = "'" + interfaceName + "'"
		}

		return "deleted interfaces: " + strings.Join(quoted, ", ")
	}

	if len(renamedInterface) == 1 {
		return "renamed interface " + renamedInterface[0]
	} else if len(renamedInterface) > 1 {
		quoted := make([]string, len(renamedInterface))
		for i, interfaceName := range renamedInterface {
			quoted[i] = "'" + interfaceName + "'"
		}

		return "renamed interfaces: " + strings.Join(quoted, ", ")
	}

	return ""
}

func FormattedEnum(diff, lang string) string {
	var oldEnum, newEnum *types.EnumSignature
	var addedEnum, deletedEnum, renamedEnum []string
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
		deletedEnum = append(deletedEnum, oldEnum.Name)
	}

	if oldEnum == nil && newEnum != nil {
		addedEnum = append(addedEnum, newEnum.Name)
	}

	if oldEnum != nil && newEnum != nil && oldEnum.Name != newEnum.Name {
		renamedEnum = append(renamedEnum, oldEnum.Name+" -> "+newEnum.Name)
	}

	if len(addedEnum) == 1 {
		return "added enum " + addedEnum[0]
	} else if len(addedEnum) > 1 {
		quoted := make([]string, len(addedEnum))
		for i, enumName := range addedEnum {
			quoted[i] = "'" + enumName + "'"
		}

		return "added enums: " + strings.Join(quoted, ", ")
	}

	if len(deletedEnum) == 1 {
		return "deleted enum " + deletedEnum[0]
	} else if len(deletedEnum) > 1 {
		quoted := make([]string, len(deletedEnum))
		for i, enumName := range deletedEnum {
			quoted[i] = "'" + enumName + "'"
		}

		return "deleted enums: " + strings.Join(quoted, ", ")
	}

	if len(renamedEnum) == 1 {
		return "renamed enum " + renamedEnum[0]
	} else if len(renamedEnum) > 1 {
		quoted := make([]string, len(renamedEnum))
		for i, enumName := range renamedEnum {
			quoted[i] = "'" + enumName + "'"
		}

		return "renamed enums: " + strings.Join(quoted, ", ")
	}

	return ""
}
