// purepath provides purely computation operations with I/O.
package purepath

import (
	"fmt"
	"runtime"
)

var platform string = runtime.GOOS

// PurePath provides purely computational operations without I/O.
type PurePath interface {
	// pathlib methods:
	Anchor() string
	Drive() string
	IsAbsolute() bool
	JoinPath(...string) PurePath
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
	// Match() bool
	// Parents() []PurePath
	// RelativeTo() PurePath
	// Suffixes() []string
	// WithName() PurePath
	// WithSegments() PurePath
	// WithStem() PurePath
	// WithSuffix() PurePath

	// added to purepath for go-pathlib:
	AsString() string
}

// Takes any number of strings, separated by commas.
// Returns either a PurePosixPath or PureWindowsPath, depending on `runtime.GOOS`.
func New(pathsegments ...string) PurePath {
	switch platform {
	case "darwin":
		return NewPurePosixPath(pathsegments...)
	case "linux":
		return NewPurePosixPath(pathsegments...)
	case "posix":
		fmt.Printf("WARNING: operating system not yet implemented or tested: %s\nProceed with caution.\n", platform)
		return NewPurePosixPath(pathsegments...)
	case "windows":
		return NewPureWindowsPath(pathsegments...)
	default:
		panic(1)
	}
}
