package utils

import (
	"log"
	"os"
	"path/filepath"
)

func GetTestsPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(home, ".local", "share", "go-pathlib", "tests")
}

func MakeTestsDir(dir string) {
	err := os.MkdirAll(dir, 0o755)
	if err != nil {
		log.Fatal(err)
	}
}
