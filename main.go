package main

import "fmt"

func main() {
	// Изменения
	files, err := GetStagedFiles()
	if err != nil {
		fmt.Println("Error getting staged files:", err)
		return
	}

	if len(files) == 0 {
		fmt.Println("No files staged for commit.")
		return
	}

	// Парсер
	parserMsg, err := Parser(files)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(parserMsg)
}
