package main

import (
	"fmt"
	"strings"
)

func InfoLogger(msg string) {
	var builder strings.Builder
	builder.Reset()
	builder.WriteString("[git auto-commit] ")
	builder.WriteString(msg)
	fmt.Println(builder.String())
}

func GitLogger(msg string) {
	var builder strings.Builder
	builder.Reset()
	builder.WriteString("\033[0;34m[git auto-commit] ")
	builder.WriteString(msg)
	builder.WriteString("\033[0m\n")
	fmt.Print(builder.String())
}

func ErrorLogger(err error) {
	var builder strings.Builder
	builder.Reset()
	builder.WriteString("\033[0;31m[git auto-commit] ")
	builder.WriteString(err.Error())
	builder.WriteString("\033[0m\n")
	fmt.Print(builder.String())
}
