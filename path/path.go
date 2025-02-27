package path

import (
	"fmt"
	"runtime"

	// "os/user"
	"github.com/rewgs/go-pathlib/purepath"
)

var platform string = runtime.GOOS

// Path: A representation of a filepath.
type Path interface {
	// path methods:
	Exists() (bool, error)
	Home() (Path, error)
	// TODO:
	// Absolute() Path
	// AsURI() Path
	// Chmod() error
	// Cwd() Path
	// ExpandUser() Path
	// FromURI() Path
	// Glob() []Path
	// Group() string // from user.User?
	// HardlinkTo() error
	// IsBlockDevice() bool
	// IsCharDevice() bool
	// IsDir() bool
	// IsFifo() bool
	// IsFile() bool
	// IsMount() bool
	// IsSocket() bool
	// IsSymlink() bool
	// Iterdir() // Not sure what to return here. error for sure, what else?
	// Lchmod() error
	// Lstat() fs.FileInfo
	// MkDir(fs.FileMode) error
	// Open() // Probably returns a Handler?
	// Owner() string // from user.User?
	// ReadBytes() []bytes
	// ReadLink() Path
	// ReadText() string
	// Rename(Path) Path
	// Replace(Path) Path
	// Resolve() Path
	// Rglob() []Path
	// RmDir(Path) error
	// SameFile() bool
	// Stat() fs.FileInfo
	// SymlinkTo() error
	// Touch(fs.FileMode) error
	// Unlink() error
	// Walk() // Not sure what to return here. error for sure, what else?
	// WriteBytes() error // will require a Writer
	// WriteText() error // will require a Writer

	// pathlib methods:
	Anchor() string
	Drive() string
	IsAbsolute() bool
	Name() string
	Parent() purepath.PurePath
	Parts() []string
	Root() string
	Stem() string
	Suffix() string

	// added to purepath for go-pathlib:
	AsString() string
}

// Takes any number of strings, separated by commas.
// Returns either a PurePosixPath or PureWindowsPath, depending on `runtime.GOOS`.
func New(pathsegments ...string) (Path, error) {
	switch platform {
	case "darwin":
		return NewPosixPath(pathsegments...), nil
	case "linux":
		return NewPosixPath(pathsegments...), nil
	case "posix":
		return NewPosixPath(pathsegments...), fmt.Errorf("Operating system not yet implemented or tested: %s\nProceed with caution.\n", platform)
	case "windows":
		return NewWindowsPath(pathsegments...), nil
	default:
		panic(1)
	}
}
