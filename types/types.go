package types

type VariableSignature struct {
	Type  string
	Name  string
	Value string
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

type ClassSignature struct {
	Name    string
	Parent  string
	Methods map[string]string
}

type SwitchSignature struct {
	Expr  string
	Cases []string
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

type GithubIssue struct {
	Title  string `json:"title"`
	Number uint32 `json:"number"`
}
