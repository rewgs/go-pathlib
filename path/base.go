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
//
// Chmod changes the file mode and permissions.chmod().
func (b base) Chmod(mode *fs.FileMode, followSymlinks bool) error {
	panic("Path.Chmod() not yet implemented.")
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

// TODO: Impelement similar functionality: "If a home directory can’t be resolved, RuntimeError is raised."
// TODO:
//
// ExpanderUser returns a new Path with expanded ~ and ~user constructs, as returned by os.UserHomeDir().
func (b base) ExpandUser() Path {
	panic("Path.ExpandUser() not yet implemented.")
}

// TODO:
//
// FromURI returns a new Path from parsing a ‘file’ URI
func (b base) FromURI(uri string) Path {
	panic("Path.FromURI() not yet implemented.")
}

// TODO:
//
// Glob globs the given relative pattern in the directory represented by this path, yielding all matching files (of any kind).
// By default, or when the caseSensitive argument is false, this method matches paths using platform-specific casing rules:
// typically, case-sensitive on POSIX, and case-insensitive on Windows. Set caseSensitive to true or false to override this behaviour.
func (b base) Glob(pattern string, caseSensitive bool, recurseSymlinks bool) []Path {
	panic("Path.Glob() not yet implemented.")
}

// TODO: Impelement similar functionality: "KeyError is raised if the file’s group identifier (GID) isn’t found in the system database."
// TODO:
//
// Group returns the name of the group owning the file.
//
// from user.User?
func (b base) Group(followSymlinks bool) string {
	panic("Path.Group() not yet implemented.")
}

// TODO:
//
// HardlinkTo makes this Path a hard link to target file.
func (b base) HardlinkToPath(target Path) error {
	panic("Path.HardlinkTo() not yet implemented.")
}

// TODO:
//
// HardlinkTo makes this Path a hard link to target file.
func (b base) HardlinkToString(target string) error {
	panic("Path.HardlinkTo() not yet implemented.")
}

// TODO:
//
// IsBlockDevice returns true if the Path points to a block device (or a symbolic link pointing to a block device),
// false if it points to another kind of file.
// Also returns false if the Path doesn’t exist or is a broken symlink;
// other errors (such as permission errors) are propagated.
func (b base) IsBlockDevice() (bool, error) {
	panic("Path.IsBlockDevice() not yet implemented.")
}

// TODO:
//
// IsCharDevice returns true if the Path points to a character device (or a symbolic link pointing to a character device),
// false if it points to another kind of file.
// Also returns false if the Path doesn’t exist or is a broken symlink;
// other errors (such as permission errors) are propagated.
func (b base) IsCharDevice() (bool, error) {
	panic("Path.IsCharDevice() not yet implemented.")
}

// TODO:
//
// IsDir returns true if the Path points to a directory, false if it points to another kind of file.
// Also returns false if the Path doesn’t exist or is a broken symlink;
// other errors (such as permission errors) are propagated.
func (b base) IsDir(followSymlinks bool) (bool, error) {
	panic("Path.IsDir() not yet implemented.")
}

// TODO:
//
// IsFIFO returns true if the Path points to a FIFO (or a symbolic link pointing to a FIFO),
// false if it points to another kind of file.
// also returns false if the Path doesn’t exist or is a broken symlink;
// other errors (such as permission errors) are propagated.
func (b base) IsFIFO() (bool, error) {
	panic("Path.IsFifo() not yet implemented.")
}

// TODO:
//
// IsFile returns true if the Path points to a regular file, false if it points to another kind of file.
// Also returns false if the Path doesn’t exist or is a broken symlink;
// other errors (such as permission errors) are propagated.
func (b base) IsFile(followSymlinks bool) (bool, error) {
	panic("Path.IsFile() not yet implemented.")
}

// TODO:
//
// IsMount returns true if the Path is a mount point (a point in a file system where a different
// file system has been mounted).
// On POSIX, the function checks whether Path’s parent, path/.., is on a different device than path,
// or whether path/.. and path point to the same i-node on the same device — this should detect mount
// points for all Unix and POSIX variants.
// On Windows, a mount point is considered to be a drive letter root (e.g. c:\), a UNC share (e.g. \\server\share),
// or a mounted filesystem directory.
func (b base) IsMount() bool {
	panic("Path.IsMount() not yet implemented.")
}

// TODO:
//
// IsSocket returns true if the path points to a Unix socket (or a symbolic link pointing to a Unix socket),
// false if it points to another kind of file.
// Also returns false if the Path doesn’t exist or is a broken symlink;
// other errors (such as permission errors) are propagated.
func (b base) IsSocket() (bool, error) {
	panic("Path.IsSocket() not yet implemented.")
}

// TODO:
//
// IsSymlinks returns true if the path points to a symbolic link, false otherwise.
// Also returns false if the path doesn’t exist; other errors (such as permission errors) are propagated.
func (b base) IsSymlink() (bool, error) {
	panic("Path.IsSymlink() not yet implemented.")
}

// TODO: Translate Python's concept of yield into Go.
// TODO:
//
// Iterdir yields Path objects of the Path's contents, as long as Path is a directory.
//
// More from the original Python pathlib doc comment -- not sure how much of this will apply:
// The children are yielded in arbitrary order, and the special entries '.' and '..' are not included.
// If a file is removed from or added to the directory after creating the iterator,
// it is unspecified whether a path object for that file is included.
// If the path is not a directory or otherwise inaccessible, OSError is raised.
func (b base) Iterdir() []Path {
	panic("Path.Iterdir() not yet implemented.")
}

// TODO:
//
// Lchmod is like Path.Chmod(), but if the Path points to a symbolic link, the symbolic link’s mode is changed
// rather than its target’s.
func (b base) Lchmod(mode fs.FileMode) error {
	panic("Path.Lchmod() not yet implemented.")
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
