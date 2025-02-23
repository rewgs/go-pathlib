package windows

import (
	"log"
	"strings"
	"unicode"
)

const (
	separator string = "\\"
)

type PureWindowsPath struct {
	path string
}

// Takes any number of strings, separated by commas.
func New(pathsegments ...string) *PureWindowsPath {
	if len(pathsegments) == 0 {
		log.Fatal("Cannot create path.")
	}
	return &PureWindowsPath{
		path: strings.Join(pathsegments, separator),
	}
}

// Takes a slice of strings.
func NewFromSlice(pathsegments []string) *PureWindowsPath {
	if len(pathsegments) == 0 {
		log.Fatal("Cannot create path.")
	}
	return &PureWindowsPath{
		path: strings.Join(pathsegments, separator),
	}
}

// A string representing the drive letter or name, if any.
func (p *PureWindowsPath) Drive() string {
	if !unicode.IsLetter(p.path[0]) {
		return ""
	}
	if p.path[1] != ":" || p.path[2] != "/" || p.path[2] != "\\" {
		return ""
	}
	return p.path[:3]
}

// A slice giving access to the path's various components.
func (p *PureWindowsPath) Parts() []string {
	return strings.Split(p.path, separator)
}

// TODO:
// A string representing the (local or global) root, if any.
func (p *PureWindowsPath) Root() string {
}
