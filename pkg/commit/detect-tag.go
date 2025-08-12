package commit

import (
	"git-auto-commit/infra/constants"
	"path/filepath"
	"regexp"
	"strings"
)

func DetectTagByFile(filename *string, tag string) string {
	// check type changed
	switch tag {
	case constants.NameStatus_Added:
		return constants.Type_CommitFeat

	case constants.NameStatus_Deleted:
		return constants.Type_CommitRefactor

	case constants.NameStatus_Modified:
		return constants.Type_CommitFix

	case constants.NameStatus_Renamed:
		return constants.Type_CommitRefactor

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
