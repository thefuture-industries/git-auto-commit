package main

import "fmt"

func appendMsg(commitMsg, addition string) string {
	if len(commitMsg) == 0 {
		return addition
	}

	return fmt.Sprintf("%s | %s", commitMsg, addition)
}

func Parser(files []string) (string, error) {
	var commitMsg string = ""

	for _, file := range files {
		if uint16(len(commitMsg)) > MAX_STRING_LENGTH {
			break
		}

		diff, err := GetDiff(file)
		if err != nil {
			return "", err
		}

		lang := DetectLanguage(file)
		if lang == "" {
			commitMsg = appendMsg(commitMsg, fmt.Sprintf("the '%s' file has been changed", file))
			continue // README.md, etc.
		}

		for _, formatted := range []string{
			FormattedVariables(diff, lang),
			FormattedFunction(diff, lang),
			FormattedClass(diff, lang),
			FormattedLogic(diff, lang),
			FormattedStruct(diff, lang),
			FormattedType(diff, lang),
			FormattedInterface(diff, lang),
			FormattedEnum(diff, lang),
		} {
			if formatted != "" {
				commitMsg = appendMsg(commitMsg, formatted)
			} // else -> continue
		}
	}

	if len(commitMsg) == 0 {
		formattedByRemote, err := FormattedByRemote("")
		if err != nil {
			return "", err
		}

		formattedByBranch, err := FormattedByBranch()
		if err != nil {
			return "", err
		}

		if formattedByRemote != "" {
			commitMsg = appendMsg(commitMsg, formattedByRemote)
		} else {
			commitMsg = appendMsg(commitMsg, formattedByBranch)
		}
	}

	return commitMsg, nil
}
