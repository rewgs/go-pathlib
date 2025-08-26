package path

import (
	"errors"
	"io/fs"
	"log"
	"os"

	"github.com/rewgs/go-pathlib/purepath"
)

// base forms the basis for the PosixPath and WindowsPath structs.
// It contains an embedded PurePath struct matching the appropriate operating system.
// Any methods that are shared between both flavors are implemented via this struct, whereas
// any methods that differ between flavors are implemented via their respective struct.
type base struct {
	purepath.PurePath
}

// newBase returns a base with an embedded PurePath struct matching the current runtime.GOOS.
func newBase(pathsegments ...string) base {
	return base{
		purepath.New(pathsegments...),
	}
}

// TODO:
// - Utilize add `followSymlinks bool` argument.
// - Check for permission-related errors in cases where file exists.
//
// Exists returns true if the path points to an existing file or directory.
func (p base) Exists() bool {
	fileInfo, err := os.Stat(p.AsString())
	if errors.Is(err, os.ErrNotExist) {
		return false
	} else if err != nil && fileInfo == nil {
		log.Fatal(err)
	}
	return true
}

// MkDir creates a new directory at this given Path.
// If mode is nil, it is set to 0o777.
// If parents is true, any missing parents of this path are created as needed; they are created with
// the default permissions without taking `mode` into account (mimicking the POSIX `mkdir -p` command).
// If existOK is true, os.ErrExist will not be raised unless the given path already exists in the
// file system and is not a directory (same behavior as the POSIX `mkdir -p` command).
func (p base) MkDir(mode *fs.FileMode, parents bool, existOK bool) error {
	if p.Exists() && !existOK {
		return os.ErrExist
	}

	if mode == nil {
		perm := fs.FileMode(0o777)
		mode = &perm
	}

	if parents {
		err := os.MkdirAll(p.AsString(), *mode)
		if err != nil {
			return err
		}
	} else {
		err := os.Mkdir(p.AsString(), *mode)
		if err != nil {
			return err
		}
	}

	return nil
}

// RmDir removes this directory.
// Unlike Python's Path.rmdir(), the directory does not have to be empty.
func (p base) RmDir() {
	err := os.RemoveAll(p.AsString())
	if err != nil {
		log.Fatal(err)
	}
}

func Absolute() Path {
	panic("Path.Absolute() not yet implemented.")
}

func AsURI() Path {
	panic("Path.AsURI() not yet implemented.")
}

func Chmod() error {
	panic("Path.Chmod() not yet implemented.")
}

func ExpandUser() Path {
	panic("Path.ExpandUser() not yet implemented.")
}

func FromURI() Path {
	panic("Path.FromURI() not yet implemented.")
}

func Glob() []Path {
	panic("Path.Glob() not yet implemented.")
}

// from user.User?
func Group() string {
	panic("Path.Group() not yet implemented.")
}

func HardlinkTo() error {
	panic("Path.HardlinkTo() not yet implemented.")
}

func IsBlockDevice() bool {
	panic("Path.IsBlockDevice() not yet implemented.")
}

func IsCharDevice() bool {
	panic("Path.IsCharDevice() not yet implemented.")
}

func IsDir() bool {
	panic("Path.IsDir() not yet implemented.")
}

func IsFifo() bool {
	panic("Path.IsFifo() not yet implemented.")
}

func IsFile() bool {
	panic("Path.IsFile() not yet implemented.")
}

func IsMount() bool {
	panic("Path.IsMount() not yet implemented.")
}

func IsSocket() bool {
	panic("Path.IsSocket() not yet implemented.")
}

func IsSymlink() bool {
	panic("Path.IsSymlink() not yet implemented.")
}

// Not sure what to return here. error for sure, what else?
func Iterdir() {
	panic("Path.Iterdir() not yet implemented.")
}

func Lchmod() error {
	panic("Path.Lchmod() not yet implemented.")
}

func Lstat() fs.FileInfo {
	panic("Path.Lstat() not yet implemented.")
}

// Probably returns a Handler?
func Open() {
	panic("Path.Open() not yet implemented.")
}

// from user.User?
func Owner() string {
	panic("Path.Owner() not yet implemented.")
}

func ReadBytes() []byte {
	panic("Path.ReadBytes() not yet implemented.")
}

func ReadLink() Path {
	panic("Path.ReadLink() not yet implemented.")
}

func ReadText() string {
	panic("Path.ReadText() not yet implemented.")
}

func Rename(Path) Path {
	panic("Path.Rename() not yet implemented.")
}

func Replace(Path) Path {
	panic("Path.Replace() not yet implemented.")
}

func Resolve() Path {
	panic("Path.Resolve() not yet implemented.")
}

func Rglob() []Path {
	panic("Path.Rglob() not yet implemented.")
}

func SameFile() bool {
	panic("Path.SameFile() not yet implemented.")
}

func Stat() fs.FileInfo {
	panic("Path.Stat() not yet implemented.")
}

func SymlinkTo() error {
	panic("Path.SymlinkTo() not yet implemented.")
}

func Touch(fs.FileMode) error {
	panic("Path.Touch() not yet implemented.")
}

func Unlink() error {
	panic("Path.Unlink() not yet implemented.")
}

// Not sure what to return here. error for sure, what else?
func Walk() {
	panic("Path.Walk() not yet implemented.")
}

// will require a Writer
func WriteBytes() error {
	panic("Path.WriteBytes() not yet implemented.")
}

// will require a Writer
func WriteText() error {
	panic("Path.WriteText() not yet implemented.")
}
