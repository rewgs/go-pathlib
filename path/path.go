// path provides "concrete paths," which inherit from PurePaths but also provide I/O operations.

package path

import (
	"fmt"
	"io/fs"
	"runtime"

	"github.com/rewgs/go-pathlib/purepath"
)

var platform string = runtime.GOOS

// Path represents a concrete path and thus provides I/O operations.
type Path interface {
	// path methods:
	Absolute() Path
	AsURI() Path
	Chmod(mode *fs.FileMode, followSymlinks bool) error
	Exists() bool
	ExpandUser() Path
	FromURI(uri string) Path
	Glob(pattern string, caseSensitive bool, recurseSymlinks bool) []Path
	Group(followSymlinksl bool) string
	HardlinkToPath(target Path) error
	HardlinkToString(target string) error
	IsBlockDevice() (bool, error)
	IsCharDevice() bool
	IsDir() bool
	IsFifo() bool
	IsFile() bool
	IsMount() bool
	IsSocket() bool
	IsSymlink() bool
	Iterdir() error
	Lchmod() error
	Lstat() fs.FileInfo
	MkDir(mode *fs.FileMode, parents bool, existOK bool) error
	Open()
	Owner() string
	ReadBytes() []byte
	ReadLink() Path
	ReadText() string
	Rename(p Path) Path
	Replace(p Path) Path
	Resolve() Path
	Rglob() []Path
	RmDir()
	SameFile() bool
	Stat() fs.FileInfo
	SymlinkTo() error
	Touch(mode fs.FileMode) error
	Unlink() error
	Walk() error
	WriteBytes() error
	WriteText() error

	// embedded pathlib methods:
	Anchor() string
	AsString() string
	Drive() string
	IsAbsolute() bool
	JoinPath(...string) purepath.PurePath
	Name() string
	Parent() purepath.PurePath
	Parts() []string
	Root() string
	Stem() string
	Suffix() string
}

// New returns either a PosixPath or WindowsPath, depending on `runtime.GOOS`.
// Takes any number of strings, separated by commas.
func New(pathsegments ...string) Path {
	switch platform {
	case "darwin":
		return NewPosixPath(pathsegments...)
	case "linux":
		return NewPosixPath(pathsegments...)
	case "posix":
		fmt.Printf("WARNING: operating system not yet implemented or tested: %s\nProceed with caution.\n", platform)
		return NewPosixPath(pathsegments...)
	case "windows":
		return NewWindowsPath(pathsegments...)
	default:
		panic(1)
	}
}

// Added for go-pathlib
func NewFromPurePath(p purepath.PurePath) Path {
	parts := p.Parts()
	return New(parts...)
}
