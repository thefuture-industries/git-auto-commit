package parser

import (
	"git-auto-commit/infra/constants"
	"path/filepath"
	"regexp"
	"strings"
)

var DetectTagByFile = func(filename *string, changed string) string {
	// check type changed
	switch changed {
	case constants.Ch_TypeAdd:
		return constants.Type_CommitFeat

	case constants.Ch_TypeDelete:
		return constants.Type_CommitRefactor

	case constants.Ch_TypeChanged:
		return constants.Type_CommitFix

	}

	if filename == nil {
		return constants.Type_CommitRefactor
	}

	ext := filepath.Ext(*filename)
	base := filepath.Base(*filename)

	// [docs]
	if ext == ".md" || ext == ".txt" {
		return constants.Type_CommitDocs
	}

	// [style]
	if ext == ".css" || ext == ".scss" || ext == ".sass" || ext == ".less" ||
		regexp.MustCompile("tailwind.*").MatchString(base) ||
		regexp.MustCompile("postcss.*").MatchString(base) {

		return constants.Type_CommitStyle
	}

	// [test]
	if strings.Contains(base, ".test.") || strings.Contains(base, "_test.") {
		return constants.Type_CommitTest
	}

	return ""
}
