package constants

const (
	Commit_Docs  = "Performed comprehensive updates and revisions to the documentation"
	Commit_Style = "Applied style improvements and formatting fixes"
	Commit_Test  = "Implemented and refined test cases to ensure code quality"
)

var Ratio_Commit = map[string]string{
	Type_CommitStyle: Commit_Style,
	Type_CommitDocs:  Commit_Docs,
	Type_CommitTest:  Commit_Test,
}
