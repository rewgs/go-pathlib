package posix

import (
	"log"
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

// A string representing the drive letter or name, if any.
func (p *PurePosixPath) Drive() string {
	return ""
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
