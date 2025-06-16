package path

import (
	"fmt"
	"io/fs"
	"runtime"
)

var platform string = runtime.GOOS

// Path: A representation of a filepath.
type Path interface {
	// path methods:

	// TODO:
	// - `followSymlinks` argument.
	// - Check for permission-related errors in cases where file exists.
	//
	// Returns true if the path points to an existing file or directory.
	// This function normally follows symlinks; to check if a symlink exists, add the argument follow_symlinks=False
	Exists() (bool, error)

	Home() (Path, error)

	// Creates a new directory at this given Path.
	//
	// If `mode` is nil, it is set to 0o777.
	//
	// If `parents` is true, any missing parents of this path are created as needed;
	// they are created with the default permissions without taking `mode` into
	// account (mimicking the POSIX `mkdir -p` command).
	//
	// If `existOK` is true, os.ErrExist will not be raised unless the given path
	// already exists in the file system and is not a directory (same behavior as
	// the POSIX `mkdir -p` command).
	MkDir(mode *fs.FileMode, parents bool, existOK bool) error

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

	// NOTE: I suppose these don't need to be present here since
	// they're in PurePath, and Path has an embedded PurePath.
	//
	// pathlib methods:
	// Anchor() string
	// Drive() string
	// IsAbsolute() bool
	// Name() string
	// Parent() purepath.PurePath
	// Parts() []string
	// Root() string
	// Stem() string
	// Suffix() string

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
