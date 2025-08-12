package tests

type parserMocks struct {
	parsedMsg string
	parseErr  error
}

type gitMocks struct {
	stagedDir string
	stagedErr error
	commitErr error
}
