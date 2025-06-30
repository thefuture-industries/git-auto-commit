package acpkg

import (
	"git-auto-commit/git"
	"strconv"
	"strings"
)

func FormattedByRemote(token string) (string, error) {
	var builder strings.Builder
	builder.Reset()

	branch, err := GetCurrentBranch()
	if err != nil {
		return "", err
	}

	issue := git.ExtractIssueNumber(branch)
	if issue == "" {
		return "", nil
	}

	owner, repo, err := git.GetOwnerRepository()
	if err != nil {
		return "", err
	}

	issueName, issueNumber, err := GetIssueData(owner, repo, issue, token)
	if err != nil {
		return "", err
	}

	builder.WriteString(issueName)
	builder.WriteString(" (")
	builder.WriteString(strconv.Itoa(int(issueNumber)))
	builder.WriteString(")")

	return builder.String(), nil
}

func FormattedByBranch() (string, error) {
	var builder strings.Builder
	builder.Reset()

	branch, err := GetCurrentBranch()
	if err != nil {
		return "", err
	}

	builder.WriteString("changed the ")
	builder.WriteString("'")
	builder.WriteString(branch)
	builder.WriteString("'")
	builder.WriteString(" branch")
	return builder.String(), err
}
