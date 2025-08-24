package purepath

import (
	"log"
	"path"
	"slices"
	"strings"

	"github.com/rewgs/go-pathlib/internal/posix"
)

type PurePosixPath struct {
	Base
}

func NewPurePosixPath(pathsegments ...string) PurePosixPath {
	if platform == "windows" {
		log.Panic("NewPosixPath does not support Windows")
	}

	if len(pathsegments) == 0 {
		log.Fatal("Cannot create path.")
	}

	// return PurePosixPath{
	// 	Base{
	// 		Filepath: strings.Join(pathsegments, posix.Separator),
	// 	},
	// }
	purePosixPath := PurePosixPath{}
	purePosixPath.Filepath = strings.Join(pathsegments, posix.Separator)
	return purePosixPath
}

// The concatenation of the drive and root.
func (p PurePosixPath) Anchor() string {
	if !strings.HasPrefix(p.Filepath, posix.Separator) {
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

func (p PurePosixPath) JoinPath(pathsegments ...string) PurePath {
	path := slices.Concat(strings.Split(p.Filepath, posix.Separator), pathsegments)
	return NewPurePosixPath(path...)
}

// TODO:
// Account for the following at https://docs.python.org/3/library/pathlib.html#pathlib.PurePath.parent
// - "You cannot go past an anchor, or empty path."
// - "Note: This is a purely lexical operation..."
//
// The logical parent of the path.
func (p PurePosixPath) Parent() PurePath {
	return NewPurePosixPath(path.Dir(p.Filepath))
}

// A slice giving access to the path's various components.
func (p PurePosixPath) Parts() []string {
	return strings.Split(p.Filepath, posix.Separator)
}

// A string representing the (local or global) root, if any.
func (p PurePosixPath) Root() string {
	if !strings.HasPrefix(p.Filepath, posix.Separator) {
		return ""
	}
	return posix.Separator
}
