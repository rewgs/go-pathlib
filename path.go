package path

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
)

type Path struct {
	path      string
	separator string
}

func NewPath(path string) (*Path, error) {
	p := Path{
		path: path,
	}

	// TODO: Implement and test other operating systems listed here:
	// https://go.dev/doc/install/source#environment
	switch os := runtime.GOOS; os {
	case "darwin":
		p.separator = "/"
	case "linux":
		p.separator = "/"
	case "windows":
		p.separator = "\\"
	default:
		return nil, fmt.Errorf("operating system not yet implemented or tested: %s", os)
	}

	return &p, nil
}

// TODO:
// - `follow_symlinks` argument.
// - Check for permission-related errors in cases where file exists.
//
// Return true if the path points to an existing file or directory.
// This function normally follows symlinks; to check if a symlink exists, add the argument follow_symlinks=False
func (p *Path) Exists() (bool, error) {
	fileInfo, err := os.Stat(p.path)
	if errors.Is(err, os.ErrNotExist) {
		return false, err
	} else if err != nil && fileInfo == nil {
		// This is a little sloppy. Need to specifically handle other errors.
		return false, err
	}
	return true, nil
}
