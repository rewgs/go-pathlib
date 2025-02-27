package purepath

import (
	"log"
	"path"
	"strings"

	"github.com/rewgs/go-pathlib/internal/posix"
)

type PurePosixPath struct {
	Shared
}

func NewPurePosixPath(pathsegments ...string) PurePosixPath {
	if len(pathsegments) == 0 {
		log.Fatal("Cannot create path.")
	}
	return PurePosixPath{
		Shared{
			Path: strings.Join(pathsegments, posix.Separator),
		},
	}
}

// The concatenation of the drive and root.
func (p PurePosixPath) Anchor() string {
	if !strings.HasPrefix(p.Path, posix.Separator) {
		return ""
	}
	return posix.Separator
}

// A string representing the drive letter or name, if any.
func (p PurePosixPath) Drive() string {
	return ""
}

func (p PurePosixPath) IsAbsolute() bool {
	return p.Root() != ""
}

// TODO:
// Account for the following at https://docs.python.org/3/library/pathlib.html#pathlib.PurePath.parent
// - "You cannot go past an anchor, or empty path."
// - "Note: This is a purely lexical operation..."
//
// The logical parent of the path.
func (p PurePosixPath) Parent() PurePath {
	return NewPurePosixPath(path.Dir(p.Path))
}

// A slice giving access to the path's various components.
func (p PurePosixPath) Parts() []string {
	return strings.Split(p.Path, posix.Separator)
}

// A string representing the (local or global) root, if any.
func (p PurePosixPath) Root() string {
	if !strings.HasPrefix(p.Path, posix.Separator) {
		return ""
	}
	return posix.Separator
}
