package utils

import (
	"log"
	"os"
	"os/user"
	// "path/filepath"
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

// func GetTestsPath() string {
// 	return filepath.Join(GetHome(), ".local", "share", "go-pathlib", "tests")
// }

// Even though go-pathlib implements as MkDir function, this needs to exist so that directories can
// be made in tests without relying on the library that is being tested.
func MakeDir(dir string) {
	err := os.MkdirAll(dir, 0o755)
	if err != nil {
		log.Fatal(err)
	}
}
