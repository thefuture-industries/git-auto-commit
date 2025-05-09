package main

import "fmt"

func InfoLogger(msg string) {
	fmt.Printf("[git auto-commit] %s\n", msg)
}

func GitLogger(msg string) {
	fmt.Printf("\033[0;34m[git auto-commit] %s\033[0m\n", msg)
}

func ErrorLogger(err error) {
	fmt.Printf("\033[0;31m[git auto-commit] %s\033[0m\n", err.Error())
}
