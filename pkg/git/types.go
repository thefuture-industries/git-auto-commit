package git

type Git struct{}

type GitInterface interface {
	// deff.go
	GetDiff(file string) (string, error)
	GetStagedCountDirectory() (string, error)
	GetStagedFiles() ([]string, error)
	
	// commit.go
	Commit(commitMsg string) error

	// git.go
	GetGitRoot() (string, error)
	GetCurrentBranch() (string, error)

	// issues.go
	ExtractIssueNumber(branch string) string
	GetOwnerRepository() (string, string, error)
	GetIssueData(owner, repo, issue, token string) (string, uint32, error)
}


type GithubIssue struct {
	Title  string `json:"title"`
	Number uint32 `json:"number"`
}
