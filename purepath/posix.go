package purepath

import (
	"log"
	"path"
	"slices"
	"strings"

	"github.com/rewgs/go-pathlib/internal/posix"
)

// PurePosixPath is the flavor of PurePath which represents non-Windows filesystem paths.
type PurePosixPath struct {
	base
}

// NewPurePosixPath returns a newly-instantiated PurePosixPath.
func NewPurePosixPath(pathsegments ...string) PurePosixPath {
	if platform == "windows" {
		log.Panic("NewPosixPath does not support Windows")
	}

	if len(pathsegments) == 0 {
		log.Fatal("Cannot create path.")
	}

	purePosixPath := PurePosixPath{}
	purePosixPath.filepath = strings.Join(pathsegments, posix.Separator)
	return purePosixPath
}

// Anchor returns the concatenation of the drive and root.
func (p PurePosixPath) Anchor() string {
	if !strings.HasPrefix(p.filepath, posix.Separator) {
		return ""
	}
	return posix.Separator
}

// Drive returns the string representing the drive letter or name, if any.
func (p PurePosixPath) Drive() string {
	return ""
}

// IsAbsolute returns whether the path is absolute or not. A path is considered absolute if it has both a root and (if the flavour allows) a drive.
func (p PurePosixPath) IsAbsolute() bool {
	return p.Root() != ""
}

// IsReserved always returns False, as Posix does not implement reserved paths.
func (p PurePosixPath) IsReserved() bool {
	return false
}

// JoinPath returns a new PurePath which each of the pathsegments combined to the path.
func (p PurePosixPath) JoinPath(pathsegments ...string) PurePath {
	path := slices.Concat(strings.Split(p.filepath, posix.Separator), pathsegments)
	return NewPurePosixPath(path...)
}

// TODO:
// Account for the following at https://docs.python.org/3/library/pathlib.html#pathlib.PurePath.parent
// - "You cannot go past an anchor, or empty path."
// - "Note: This is a purely lexical operation..."
//
// The logical parent of the path.
func (p PurePosixPath) Parent() PurePath {
	return NewPurePosixPath(path.Dir(p.filepath))
}

// A slice giving access to the path's various components.
func (p PurePosixPath) Parts() []string {
	return strings.Split(p.filepath, posix.Separator)
}

// A string representing the (local or global) root, if any.
func (p PurePosixPath) Root() string {
	if !strings.HasPrefix(p.filepath, posix.Separator) {
		return ""
	}
	return posix.Separator
}
