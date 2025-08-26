package path

import (
	"log"
	"os"
)

// TODO: Implement similar functionality: "If the home directory can't be resolved, RuntimeError is raised."
//
// Home returns a new Path representing the user's home directory.
func Home() Path {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	path := New(home)
	return path
}

// Cwd returns a new Path representing the current directory.
func Cwd() Path {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return New(dir)
}
