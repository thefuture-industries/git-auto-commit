package tests

func (m *gitMocks) GetStagedCountDirectory() (string, error) {
	return m.stagedDir, m.stagedErr
}

func (m *gitMocks) Commit(msg string) error {
	return m.commitErr
}

func (m *gitMocks) GetDiff(file string) (string, error) {
	return "", nil
}

func (m *gitMocks) GetStagedFiles() ([]string, error) {
	return nil, nil
}

func (m *gitMocks) GetGitRoot() (string, error) {
	return "", nil
}

func (m *gitMocks) GetCurrentBranch() (string, error) {
	return "", nil
}

func (m *gitMocks) ExtractIssueNumber(branch string) string {
	return ""
}

func (m *gitMocks) GetOwnerRepository() (string, string, error) {
	return "", "", nil
}

func (m *gitMocks) GetIssueData(owner, repo, issue, token string) (string, uint32, error) {
	return "", 0, nil
}

func (m *parserMocks) ParserIndex(directory string) (string, error) {
	return m.parsedMsg, m.parseErr
}

func (m *parserMocks) CreateAutoCommitMsg(filename, msg *string, changed string) string {
	return ""
}

func (m *parserMocks) DetectTagByFile(filename *string, changed string) string {
	return ""
}
