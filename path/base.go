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
//
// Absolute makes the path absolute, without normalization or resolving symlinks. Returns a new Path.
func (b base) Absolute() Path {
	panic("Path.Absolute() not yet implemented.")
}

// TODO: Impelement similar functionality: "ValueError is raised if the path isn’t absolute."
// TODO:
//
// AsURI represents the path as a ‘file’ URI.
func (b base) AsURI() Path {
	panic("Path.AsURI() not yet implemented.")
}

// TODO:
// - Utilize add `followSymlinks bool` argument.
// - Check for permission-related errors in cases where file exists.
//
// Exists returns true if the path points to an existing file or directory.
func (b base) Exists() bool {
	fileInfo, err := os.Stat(b.AsString())
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
func (b base) MkDir(mode *fs.FileMode, parents bool, existOK bool) error {
	if b.Exists() && !existOK {
		return os.ErrExist
	}

	if mode == nil {
		perm := fs.FileMode(0o777)
		mode = &perm
	}

	if parents {
		err := os.MkdirAll(b.AsString(), *mode)
		if err != nil {
			return err
		}
	} else {
		err := os.Mkdir(b.AsString(), *mode)
		if err != nil {
			return err
		}
	}

	return nil
}

// RmDir removes this directory.
// Unlike Python's Path.rmdir(), the directory does not have to be empty.
func (b base) RmDir() {
	err := os.RemoveAll(b.AsString())
	if err != nil {
		log.Fatal(err)
	}
}

///////////////////////////////////////////////////////////////////////////////

// TODO:
func (b base) Chmod(mode *fs.FileMode, followSymlinks bool) error {
	panic("Path.Chmod() not yet implemented.")
}

// TODO:
func (b base) ExpandUser() Path {
	panic("Path.ExpandUser() not yet implemented.")
}

// TODO:
func (b base) FromURI() Path {
	panic("Path.FromURI() not yet implemented.")
}

// TODO:
func (b base) Glob() []Path {
	panic("Path.Glob() not yet implemented.")
}

// TODO:
//
// from user.User?
func (b base) Group() string {
	panic("Path.Group() not yet implemented.")
}

// TODO:
func (b base) HardlinkTo() error {
	panic("Path.HardlinkTo() not yet implemented.")
}

// TODO:
func (b base) IsBlockDevice() bool {
	panic("Path.IsBlockDevice() not yet implemented.")
}

// TODO:
func (b base) IsCharDevice() bool {
	panic("Path.IsCharDevice() not yet implemented.")
}

// TODO:
func (b base) IsDir() bool {
	panic("Path.IsDir() not yet implemented.")
}

// TODO:
func (b base) IsFifo() bool {
	panic("Path.IsFifo() not yet implemented.")
}

// TODO:
func (b base) IsFile() bool {
	panic("Path.IsFile() not yet implemented.")
}

// TODO:
func (b base) IsMount() bool {
	panic("Path.IsMount() not yet implemented.")
}

// TODO:
func (b base) IsSocket() bool {
	panic("Path.IsSocket() not yet implemented.")
}

// TODO:
func (b base) IsSymlink() bool {
	panic("Path.IsSymlink() not yet implemented.")
}

// TODO:
//
// Not sure what to return here. error for sure, what else?
func (b base) Iterdir() error {
	panic("Path.Iterdir() not yet implemented.")
}

// TODO:
func (b base) Lchmod() error {
	panic("Path.Lchmod() not yet implemented.")
}

// TODO:
func (b base) Lstat() fs.FileInfo {
	panic("Path.Lstat() not yet implemented.")
}

// TODO:
//
// Probably returns a Handler?
func (b base) Open() {
	panic("Path.Open() not yet implemented.")
}

// TODO:
//
// from user.User?
func (b base) Owner() string {
	panic("Path.Owner() not yet implemented.")
}

// TODO:
func (b base) ReadBytes() []byte {
	panic("Path.ReadBytes() not yet implemented.")
}

// TODO:
func (b base) ReadLink() Path {
	panic("Path.ReadLink() not yet implemented.")
}

// TODO:
func (b base) ReadText() string {
	panic("Path.ReadText() not yet implemented.")
}

// TODO:
func (b base) Rename(p Path) Path {
	panic("Path.Rename() not yet implemented.")
}

// TODO:
func (b base) Replace(p Path) Path {
	panic("Path.Replace() not yet implemented.")
}

// TODO:
func (b base) Resolve() Path {
	panic("Path.Resolve() not yet implemented.")
}

// TODO:
func (b base) Rglob() []Path {
	panic("Path.Rglob() not yet implemented.")
}

// TODO:
func (b base) SameFile() bool {
	panic("Path.SameFile() not yet implemented.")
}

// TODO:
func (b base) Stat() fs.FileInfo {
	panic("Path.Stat() not yet implemented.")
}

// TODO:
func (b base) SymlinkTo() error {
	panic("Path.SymlinkTo() not yet implemented.")
}

// TODO:
func (b base) Touch(mode fs.FileMode) error {
	panic("Path.Touch() not yet implemented.")
}

// TODO:
func (b base) Unlink() error {
	panic("Path.Unlink() not yet implemented.")
}

// TODO:
//
// Not sure what to return here. error for sure, what else?
func (b base) Walk() error {
	panic("Path.Walk() not yet implemented.")
}

// TODO:
//
// will require a Writer
func (b base) WriteBytes() error {
	panic("Path.WriteBytes() not yet implemented.")
}

// TODO:
//
// will require a Writer
func (b base) WriteText() error {
	panic("Path.WriteText() not yet implemented.")
}
