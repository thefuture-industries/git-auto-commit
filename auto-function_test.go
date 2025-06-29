package main

import "testing"

func TestParser_AddedGoFunction(t *testing.T) {
	GetDiff = func(file string) (string, error) {
		return "+func TestParser() {}", nil
	}

	DetectLanguage = 
}
