package main

import (
	"fmt"
	"sync"
)

func appendMsg(commitMsg, addition string) string {
	if len(commitMsg) == 0 {
		return addition
	}

	return fmt.Sprintf("%s | %s", commitMsg, addition)
}

func Parser(files []string) (string, error) {
	var (
		payloadMsg string
		mu         sync.Mutex
		wg         sync.WaitGroup
		errChan    = make(chan error, len(files))
	)

	workers := 3
	jobs := make(chan string, len(files))

	for _, file := range files {
		if uint16(len(payloadMsg)) > MAX_COMMIT_LENGTH {
			break
		}

		diff, err := GetDiff(file)
		if err != nil {
			return "", err
		}

		lang := DetectLanguage(file)
		if lang == "" {
			payloadMsg = appendMsg(payloadMsg, fmt.Sprintf("the '%s' file has been changed", file))
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
				payloadMsg = appendMsg(payloadMsg, formatted)
			} // else -> continue
		}
	}

	if len(payloadMsg) == 0 {
		formattedByRemote, err := FormattedByRemote("")
		if err != nil {
			return "", err
		}

		formattedByBranch, err := FormattedByBranch()
		if err != nil {
			return "", err
		}

		if formattedByRemote != "" {
			payloadMsg = appendMsg(payloadMsg, formattedByRemote)
		} else {
			payloadMsg = appendMsg(payloadMsg, formattedByBranch)
		}
	}

	return payloadMsg, nil
}
