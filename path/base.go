package path

import (
	"log"
	"os"
)

// Base forms the basis for the PosixPath and WindowsPath structs.
// Any methods that are shared between both flavors are implemented via this struct, whereas
// any methods that differ between flavors are implemented via their respective struct.
type Base struct {
}

func Home() Path {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	path := New(home)
	return path
}
