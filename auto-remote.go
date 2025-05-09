package main

import "fmt"

func FormattedByRemote(token string) (string, error) {
	branch, err := GetCurrentBranch()
	if err != nil {
		return "", err
	}

	issue := ExtractIssueNumber(branch)
	if issue == "" {
		return "", nil
	}

	owner, repo, err := GetOwnerRepository()
	if err != nil {
		return "", err
	}

	issueName, issueNumber, err := GetIssueData(owner, repo, issue, token)
	if err != nil {
		return "", err
	}

	commitMsg := fmt.Sprintf("%s (%d)", issueName, issueNumber)

	return commitMsg, nil
}

func FormattedByBranch() (string, error) {
	branch, err := GetCurrentBranch()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("changed the '%s' branch", branch), err
}
