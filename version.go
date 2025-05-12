package main

func GetVersion(root string) string {
	versionFile := filepath.Join(gitRoot, ".git", "hooks", "auto-commit.version.txt")
    version, err := ioutil.ReadFile(versionFile)
    if err != nil {
        fmt.Println("Version file not found")
        return
    }
    fmt.Printf("Current version: %s\n", string(version))
}
