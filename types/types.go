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
