package path

import (
	"log"
	"os"
)

func Home() Path {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	path := New(home)
	return path
}

func Cwd() Path {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return New(dir)
}
