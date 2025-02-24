package purepath

import (
	"log"
	"path"
	"strings"
)

const (
	posixSeparator string = "/"
)

type Posix struct {
	path string
}

// Takes any number of strings, separated by commas.
func NewPosix(pathsegments ...string) *Posix {
	if len(pathsegments) == 0 {
		log.Fatal("Cannot create path.")
	}
	return &Posix{
		path: strings.Join(pathsegments, posixSeparator),
	}
}

// Takes a slice of strings.
func NewPosixFromSlice(pathsegments []string) *Posix {
	if len(pathsegments) == 0 {
		log.Fatal("Cannot create path.")
	}
	return &Posix{
		path: strings.Join(pathsegments, posixSeparator),
	}
}

func NewPosixFromString(path string) *Posix {
	if len(path) == 0 {
		log.Fatal("Cannot create path.")
	}
	return &Posix{
		path: path,
	}
}

// The concatenation of the drive and root.
func (p *Posix) Anchor() string {
	if !strings.HasPrefix(p.path, posixSeparator) {
		return ""
	}
	return posixSeparator
}

// A string representing the drive letter or name, if any.
func (p *Posix) Drive() string {
	return ""
}

func (p *Posix) Name() string {
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
func (p *Posix) Parent() PurePath {
	return NewPosixFromString(path.Dir(p.path))
}

// A slice giving access to the path's various components.
func (p *Posix) Parts() []string {
	return strings.Split(p.path, posixSeparator)
}

// A string representing the (local or global) root, if any.
func (p *Posix) Root() string {
	if !strings.HasPrefix(p.path, posixSeparator) {
		return ""
	}
	return posixSeparator
}

func (p *Posix) Stem() string {
	name := p.Name()
	ext := p.Suffix()

	before, found := strings.CutSuffix(name, ext)
	if !found {
		log.Fatalf("Could not find %s in %s", ext, name)
	}
	return before
}

func (p *Posix) Suffix() string {
	return path.Ext(p.path)
}
