package path

import (
	"errors"
	"os"
)

type Path struct {
	path string
}

func NewPath(path string) *Path {
	return &Path{
		path: path,
	}
}

// TODO:
// - `follow_symlinks` argument.
// - Check for permission-related errors in cases where file exists.
//
// Return true if the path points to an existing file or directory.
// This function normally follows symlinks; to check if a symlink exists, add the argument follow_symlinks=False
func (p *Path) Exists() (bool, error) {
	fileInfo, err := os.Stat(p.path)
	if fileInfo == nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, os.ErrNotExist
		} else if err != nil {
			// This is a little sloppy. Need to specifically handle other errors.
			return false, err
		} else {
			// This is probably impossible right?
			return false, nil
		}
	} else {
		// This should be impossible, right?
		if errors.Is(err, os.ErrNotExist) {
			return true, os.ErrNotExist
		} else if err != nil {
			// This is a little sloppy. Need to specifically handle other errors.
			return true, err
		}
	}
	return true, nil
}
