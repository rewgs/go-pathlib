package path

import (
	"errors"
	"log"
	"os"
)

type Shared struct {
	Path string
}

// TODO:
// - `followSymlinks` argument.
// - Check for permission-related errors in cases where file exists.
//
// Returns true if the path points to an existing file or directory.
// This function normally follows symlinks; to check if a symlink exists, add the argument follow_symlinks=False
func (p Shared) Exists() (bool, error) {
	fileInfo, err := os.Stat(p.Path)
	if errors.Is(err, os.ErrNotExist) {
		return false, err
	} else if err != nil && fileInfo == nil {
		// This is a little sloppy, but for now it's fine.
		// TODO: Specifically handle other errors.
		return false, err
	}
	return true, nil
}

func (p Shared) Home() (Path, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return New(home)
}
