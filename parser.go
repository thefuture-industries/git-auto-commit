package main

import (
	"fmt"
	"git-auto-commit/diff"
	"path/filepath"
	"strings"
	"sync"
)

func AppendMsg(commitMsg, addition string) string {
	var builder strings.Builder
	builder.Reset()

	if len(commitMsg) == 0 {
		return addition
	}

	builder.WriteString(commitMsg)
	builder.WriteString(" | ")
	builder.WriteString(addition)
	return builder.String()
}

var Parser = func(files []string) (string, error) {
	var (
		payloadMsg string
		mu         sync.Mutex
		wg         sync.WaitGroup
		errChan    = make(chan error, len(files))
	)

	workers := 3
	jobs := make(chan string, len(files))

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for file := range jobs {
				mu.Lock()
				if uint16(len(payloadMsg)) > MAX_COMMIT_LENGTH {
					mu.Unlock()
					continue
				}
				mu.Unlock()

				diff, err := diff.GetDiff(file)
				if err != nil {
					errChan <- fmt.Errorf("error getting diff for %s: %w", file, err)
					continue
				}

				lang := DetectLanguage(file)
				if lang == "" {
					mu.Lock()
					payloadMsg = AppendMsg(payloadMsg, fmt.Sprintf("the '%s' file has been changed", filepath.Base(file)))
					mu.Unlock()
					continue // README.md, etc.
				}

				var fileChanges []string
				for _, formatted := range []string{
					FormattedVariables(diff, lang),
					FormattedFunction(diff, lang),
					FormattedClass(diff, lang),
					FormattedLogic(diff, lang, filepath.Base(file)),
					FormattedImport(diff, lang, filepath.Base(file)),
					FormattedStruct(diff, lang),
					FormattedType(diff, lang),
					FormattedInterface(diff, lang),
					FormattedEnum(diff, lang),
				} {
					if formatted != "" {
						fileChanges = append(fileChanges, formatted)
					} // else -> continue
				}

				if len(fileChanges) > 0 {
					mu.Lock()
					for _, change := range fileChanges {
						nextMsg := AppendMsg(payloadMsg, change)
						if len(nextMsg) > int(MAX_COMMIT_LENGTH) {
							if len(payloadMsg) == 0 {
								if len(change) > int(MAX_COMMIT_LENGTH) {
									change = change[:int(MAX_COMMIT_LENGTH)]
								}

								payloadMsg = change
							}

							break
						}

						payloadMsg = nextMsg
					}
					mu.Unlock()
				}
			}
		}()
	}

	for _, file := range files {
		jobs <- file
	}
	close(jobs)

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return "", err
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
			payloadMsg = AppendMsg(payloadMsg, formattedByRemote)
		} else {
			payloadMsg = AppendMsg(payloadMsg, formattedByBranch)
		}
	}

	return payloadMsg, nil
}
