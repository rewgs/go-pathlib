package path

import (
	"runtime"
)

// Path: A representation of a filepath.
type Path interface {
	Exists() (bool, error)
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
