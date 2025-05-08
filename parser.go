package main

import "fmt"

func Parser(files []string) (string, error) {
	var commitMsg string = ""

	for _, file := range files {
		diff, err := GetDiff(file)
		if err != nil {
			return "", err
		}

		lang := DetectLanguage(file)
		if lang == "" {
			continue // README.md, etc.
		}

		formattedVar, err := FormattedVariables(diff, lang)
		if err != nil {
			return "", err
		}

		if formattedVar != "" {
			if len(commitMsg) == 0 {
				commitMsg = formattedVar
			} else {
				commitMsg = fmt.Sprintf(" and %s", formattedVar)
			}
		} // else -> continue
	}

	return commitMsg, nil
}
