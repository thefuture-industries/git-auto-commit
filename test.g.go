package main

type Mocks struct {
    GetStagedFiles func() ([]string, error)
    Parser         func([]string) (string, error)
    Commit         func(string) error
    ErrorLogger    func(error)
    InfoLogger     func(string)
    GetVersion     func(bool)
}

func SaveMocks
