package main

import "fmt"

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
			continue // README.md, etc.
		}

		formattedVar := FormattedVariables(diff, lang)
		if formattedVar != "" {
			if len(commitMsg) == 0 {
				commitMsg = formattedVar
			} else {
				commitMsg += fmt.Sprintf(" | %s", formattedVar)
			}
		} // else -> continue

		formattedFunc := FormattedFunction(diff, lang)
		if formattedFunc != "" {
			if len(commitMsg) == 0 {
				commitMsg = formattedFunc
			} else {
				commitMsg += fmt.Sprintf(" | %s", formattedFunc)
			}
		} // else -> continue

		formattedClass := FormattedClass(diff, lang)
		if formattedClass != "" {
			if len(commitMsg) == 0 {
				commitMsg = formattedClass
			} else {
				commitMsg += fmt.Sprintf(" | %s", formattedClass)
			}
		} // else -> continue

		formattedLogic := FormattedLogic(diff, lang)
		if formattedLogic != "" {
			if len(commitMsg) == 0 {
				commitMsg = formattedLogic
			} else {
				commitMsg += fmt.Sprintf(" | %s", formattedLogic)
			}
		} // else -> continue
	}

	return commitMsg, nil
}
