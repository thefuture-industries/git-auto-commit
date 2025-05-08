package types

type VariableSignature struct {
	Type  string
	Name  string
	Value string
}

type FunctionSignature struct {
	Name   string
	Params []FunctionParameters
}

type FunctionParameters struct {
	Name string
	Type string
}
