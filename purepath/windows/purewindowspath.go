package windows

import (
	"log"
	"path"
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

func NewFromString(path string) *PureWindowsPath {
	if len(path) == 0 {
		log.Fatal("Cannot create path.")
	}
	return &PureWindowsPath{
		path: path,
	}
}

func (p *PureWindowsPath) Anchor() string {
	drive := p.Drive()
	if drive != "" {
		return drive + separator
	}
	if unicode.IsLetter(rune(p.path[0])) && (string(p.path[1]) == "/" || string(p.path[1]) == "\\") {
		return p.path[:2]
	}
	return ""
}

// A string representing the drive letter or name, if any.
func (p *PureWindowsPath) Drive() string {
	beginsWith, driveLetter := driveLetter(p.path)
	if !beginsWith {
		return ""
	}
	return driveLetter
}

// TODO:
// Account for the following at https://docs.python.org/3/library/pathlib.html#pathlib.PurePath.parent
// - "You cannot go past an anchor, or empty path."
// - "Note: This is a purely lexical operation..."
//
// The logical parent of the path.
func (p *PureWindowsPath) Parent() *PureWindowsPath {
	return NewFromString(path.Dir(p.path))
}

// A slice giving access to the path's various components.
func (p *PureWindowsPath) Parts() []string {
	return strings.Split(p.path, separator)
}

// A string representing the (local or global) root, if any.
func (p *PureWindowsPath) Root() string {
	beginsWith, _ := driveLetter(p.path)
	if !beginsWith {
		return ""
	}
	return separator
}

// Checks if a path begins with a drive letter, and returns it if true.
func driveLetter(path string) (bool, string) {
	if !unicode.IsLetter(rune(path[0])) {
		return false, ""
	}
	if string(path[1]) != ":" {
		return false, ""
	}
	if string(path[2]) != "/" && string(path[2]) != "\\" {
		return false, ""
	}
	return true, path[:2]
}
