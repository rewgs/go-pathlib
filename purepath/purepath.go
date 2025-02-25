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
	IsAbsolute() bool
	Name() string
	Parent() PurePath
	Parts() []string
	Root() string
	Stem() string
	Suffix() string
	// TODO:
	// FullMatch() bool
	// IsRelativeTo() bool
	// IsReserved() bool
	// JoinPath() PurePath
	// Match() bool
	// Parents() []PurePath
	// RelativeTo() PurePath
	// Suffixes() []string
	// WithName() PurePath
	// WithSegments() PurePath
	// WithStem() PurePath
	// WithSuffix() PurePath
}

// Takes any number of strings, separated by commas.
// Returns either a PurePosixPath or PureWindowsPath, depending on `runtime.GOOS`.
func New(pathsegments ...string) (PurePath, error) {
	switch platform := runtime.GOOS; platform {
	case "darwin":
		return NewPurePosixPath(pathsegments...), nil
	case "linux":
		return NewPurePosixPath(pathsegments...), nil
	case "posix":
		return NewPurePosixPath(pathsegments...), nil
	case "windows":
		return NewPureWindowsPath(pathsegments...), nil
	default:
		panic(1)
	}
}
