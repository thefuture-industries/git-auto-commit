package code

import (
	"git-auto-commit/pkg/git"
)

type Code struct {
	Git git.GitInterface
}

type CodeInterface interface {
	// code.go
	FormattedCode(files []string) (string, error)

	// comment.go

	// remote.go
	FormattedByRemote(token string) (string, error)
	FormattedByBranch() (string, error)
}

type FunctionSignature struct {
	Name       string
	Params     []FunctionParameters
	ReturnType string
}

type FunctionParameters struct {
	Name string
	Type string
}

type StructureSignature struct {
	Name string
}

type TypeSignature struct {
	Name string
}

type EnumSignature struct {
	Name   string
	Values []string
}

type InterfaceSignature struct {
	Name    string
	Methods []string
}

type VariableSignature struct {
	Type  string
	Name  string
	Value string
}

type ClassSignature struct {
	Name    string
	Parent  string
	Methods map[string]string
}

type SwitchSignature struct {
	Expr  string
	Cases []string
}
