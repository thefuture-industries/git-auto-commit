package code

import (
	"strconv"
	"strings"
)

func (c *Code) FormattedByRemote(token string) (string, error) {
	var builder strings.Builder
	builder.Reset()

	branch, err := c.Git.GetCurrentBranch()
	if err != nil {
		return "", err
	}

	issue := c.Git.ExtractIssueNumber(branch)
	if issue == "" {
		return "", nil
	}

	owner, repo, err := c.Git.GetOwnerRepository()
	if err != nil {
		return "", err
	}

	issueName, issueNumber, err := c.Git.GetIssueData(owner, repo, issue, token)
	if err != nil {
		return "", err
	}

	builder.WriteString(issueName)
	builder.WriteString(" (")
	builder.WriteString(strconv.Itoa(int(issueNumber)))
	builder.WriteString(")")

	return builder.String(), nil
}

func (c *Code) FormattedByBranch() (string, error) {
	var builder strings.Builder
	builder.Reset()

	branch, err := c.Git.GetCurrentBranch()
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
