package git

type GithubIssue struct {
	Title  string `json:"title"`
	Number uint32 `json:"number"`
}
