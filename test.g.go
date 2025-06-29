package main

type Mocks struct {
	GetStagedFiles func() ([]string, error)
	Parser         func([]string) (string, error)
	Commit         func(string) error
	ErrorLogger    func(error)
	InfoLogger     func(string)
	GetVersion     func(bool)
}

func SaveMocks() *Mocks {
	return &Mocks{
		GetStagedFiles: GetStagedFiles,
		Parser:         Parser,
		Commit:         Commit,
		ErrorLogger:    ErrorLogger,
		InfoLogger:     InfoLogger,
		GetVersion:     GetVersion,
	}
}

func (m *Mocks) Apply() {
	GetStagedFiles = m.GetStagedFiles
	Parser = m.Parser
	Commit = m.Commit
	ErrorLogger = m.ErrorLogger
	InfoLogger = m.InfoLogger
	GetVersion = m.GetVersion
}
