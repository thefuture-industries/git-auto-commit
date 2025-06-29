package main

func SaveAndRestoreMocks() func() {
	origGetStagedFiles := GetStagedFiles
	origParser := Parser
	origCommit := Commit
	origErrorLogger := ErrorLogger
	origInfoLogger := InfoLogger
	origGetVersion := GetVersion

	return func() {
		GetStagedFiles = origGetStagedFiles
		Parser = origParser
		Commit = origCommit
		ErrorLogger = origErrorLogger
		InfoLogger = origInfoLogger
		GetVersion = origGetVersion
	}
}
