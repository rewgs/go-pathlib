package posix

import (
	"log"
	"path"
	"strings"
)

const (
	separator string = "/"
)

type PurePosixPath struct {
	path string
}

// Takes any number of strings, separated by commas.
func New(pathsegments ...string) *PurePosixPath {
	if len(pathsegments) == 0 {
		log.Fatal("Cannot create path.")
	}
	return &PurePosixPath{
		path: strings.Join(pathsegments, separator),
	}
}

// Takes a slice of strings.
func NewFromSlice(pathsegments []string) *PurePosixPath {
	if len(pathsegments) == 0 {
		log.Fatal("Cannot create path.")
	}
	return &PurePosixPath{
		path: strings.Join(pathsegments, separator),
	}
}

func NewFromString(path string) *PurePosixPath {
	if len(path) == 0 {
		log.Fatal("Cannot create path.")
	}
	return &PurePosixPath{
		path: path,
	}
}

// The concatenation of the drive and root.
func (p *PurePosixPath) Anchor() string {
	if !strings.HasPrefix(p.path, separator) {
		return ""
	}
	return separator
}

// A string representing the drive letter or name, if any.
func (p *PurePosixPath) Drive() string {
	return ""
}

// TODO:
// Account for the following at https://docs.python.org/3/library/pathlib.html#pathlib.PurePath.parent
// - "You cannot go past an anchor, or empty path."
// - "Note: This is a purely lexical operation..."
//
// The logical parent of the path.
func (p *PurePosixPath) Parent() *PurePosixPath {
	return NewFromString(path.Dir(p.path))
}

// A slice giving access to the path's various components.
func (p *PurePosixPath) Parts() []string {
	return strings.Split(p.path, separator)
}

// A string representing the (local or global) root, if any.
func (p *PurePosixPath) Root() string {
	if !strings.HasPrefix(p.path, separator) {
		return ""
	}
	return separator
}
