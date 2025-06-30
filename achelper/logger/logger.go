package logger

import (
	"fmt"
	"strings"
)

var InfoLogger = func(msg string) {
	var builder strings.Builder
	builder.Reset()
	builder.WriteString("[git auto-commit] ")
	builder.WriteString(msg)
	fmt.Println(builder.String())
}

var GitLogger = func(msg string) {
	var builder strings.Builder
	builder.Reset()
	builder.WriteString("\033[0;34m[git auto-commit] ")
	builder.WriteString(msg)
	builder.WriteString("\033[0m\n")
	fmt.Print(builder.String())
}

var ErrorLogger = func(err error) {
	var builder strings.Builder
	builder.Reset()
	builder.WriteString("\033[0;31m[git auto-commit] ")
	builder.WriteString(err.Error())
	builder.WriteString("\033[0m\n")
	fmt.Print(builder.String())
}
