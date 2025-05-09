package types

type VariableSignature struct {
	Type  string
	Name  string
	Value string
}

type FunctionSignature struct {
	Name   string
	Params []FunctionParameters
	ReturnType string
}

type FunctionParameters struct {
	Name string
	Type string
}
