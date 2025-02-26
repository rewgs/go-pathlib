package utils

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func GetUsername() string {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return currentUser.Username
}

func GetHome() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return home
}

func GetTestsPath() string {
	return filepath.Join(GetHome(), ".local", "share", "go-pathlib", "tests")
}

func MakeDir(dir string) {
	err := os.MkdirAll(dir, 0o755)
	if err != nil {
		log.Fatal(err)
	}
}
