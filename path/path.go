package path

import (
	"errors"
	"os"
	"runtime"
)

// Path: A representation of a filepath.
type Path interface {
	Exists() (bool, error)
}

type path struct {
	path string
}

// TODO:
// - `follow_symlinks` argument.
// - Check for permission-related errors in cases where file exists.
//
// Return true if the path points to an existing file or directory.
// This function normally follows symlinks; to check if a symlink exists, add the argument follow_symlinks=False
func (p *path) Exists() (bool, error) {
	fileInfo, err := os.Stat(p.path)
	if errors.Is(err, os.ErrNotExist) {
		return false, err
	} else if err != nil && fileInfo == nil {
		// This is a little sloppy, but for now it's fine.
		// TODO: Specifically handle other errors.
		return false, err
	}
	return true, nil
}

// Takes any number of strings, separated by commas.
// Returns either a PurePosixPath or PureWindowsPath, depending on `runtime.GOOS`.
func New(pathsegments ...string) (Path, error) {
	switch platform := runtime.GOOS; platform {
	case "darwin":
		return NewPosixPath(pathsegments...), nil
	case "linux":
		return NewPosixPath(pathsegments...), nil
	case "posix":
		return NewPosixPath(pathsegments...), nil
	case "windows":
		return NewWindowsPath(pathsegments...), nil
	default:
		panic(1)
	}
}
