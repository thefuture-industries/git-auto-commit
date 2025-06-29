package main

import "testing"

func TestParser_AddedGoFunction(t *testing.T) {
	GetDiff = func(file string) (string, error) {
		return "+func TestParser() {}", nil
	}

	DetectLanguage = func(filename string) string {
		return "go"
	}

	msg, err := Parser([]string{"test"})
}
