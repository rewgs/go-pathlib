package purepath

import (
	"log"
	"path"
	"strings"

	"github.com/rewgs/go-pathlib/internal/posix"
)

type PurePosixPath struct {
	path string
}

// Takes any number of strings, separated by commas.
func NewPurePosixPath(pathsegments ...string) *PurePosixPath {
	if len(pathsegments) == 0 {
		log.Fatal("Cannot create path.")
	}
	return &PurePosixPath{
		path: strings.Join(pathsegments, posix.Separator),
	}
}

// The concatenation of the drive and root.
func (p *PurePosixPath) Anchor() string {
	if !strings.HasPrefix(p.path, posix.Separator) {
		return ""
	}
	return posix.Separator
}

// A string representing the drive letter or name, if any.
func (p *PurePosixPath) Drive() string {
	return ""
}

func (p *PurePosixPath) Name() string {
	name := path.Base(p.path)
	if name == "." || name == "/" {
		log.Fatalf("Could not get name from %s", p.path)
	}
	return name
}

// TODO:
// Account for the following at https://docs.python.org/3/library/pathlib.html#pathlib.PurePath.parent
// - "You cannot go past an anchor, or empty path."
// - "Note: This is a purely lexical operation..."
//
// The logical parent of the path.
func (p *PurePosixPath) Parent() PurePath {
	return NewPurePosixPath(path.Dir(p.path))
}

// A slice giving access to the path's various components.
func (p *PurePosixPath) Parts() []string {
	return strings.Split(p.path, posix.Separator)
}

// A string representing the (local or global) root, if any.
func (p *PurePosixPath) Root() string {
	if !strings.HasPrefix(p.path, posix.Separator) {
		return ""
	}
	return posix.Separator
}

func (p *PurePosixPath) Stem() string {
	name := p.Name()
	ext := p.Suffix()

	before, found := strings.CutSuffix(name, ext)
	if !found {
		log.Fatalf("Could not find %s in %s", ext, name)
	}
	return before
}

func (p *PurePosixPath) Suffix() string {
	return path.Ext(p.path)
}
