// purepath provides purely computation operations with I/O.
package purepath

import (
	"runtime"
)

// A generic interfaces that represents the system's path flavour (instantiating
// it creates either a PurePosixPath or PureWindowsPath).
type PurePath interface {
	Anchor() string
	Drive() string
	Parent() PurePath
	Parts() []string
	Root() string
}

// Takes any number of strings, separated by commas.
// Returns either a PurePosixPath or PureWindowsPath, depending on `runtime.GOOS`.
func New(pathsegments ...string) (PurePath, error) {
	switch platform := runtime.GOOS; platform {
	case "darwin":
		return NewPosixFromSlice(pathsegments), nil
	case "linux":
		return NewPosixFromSlice(pathsegments), nil
	case "posix":
		return NewPosixFromSlice(pathsegments), nil
	case "windows":
		return NewWindowsFromSlice(pathsegments), nil
	default:
		panic(1)
	}
}
