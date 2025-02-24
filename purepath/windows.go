package purepath

import (
	"log"
	"path" // TODO: Test if this works with "\\" paths.
	"strings"
	"unicode"
)

const (
	windowsSeparator string = "\\"
)

type Windows struct {
	path string
}

// Takes any number of strings, separated by commas.
func NewWindows(pathsegments ...string) *Windows {
	if len(pathsegments) == 0 {
		log.Fatal("Cannot create path.")
	}
	return &Windows{
		path: strings.Join(pathsegments, windowsSeparator),
	}
}

// Takes a slice of strings.
func NewWindowsFromSlice(pathsegments []string) *Windows {
	if len(pathsegments) == 0 {
		log.Fatal("Cannot create path.")
	}
	return &Windows{
		path: strings.Join(pathsegments, windowsSeparator),
	}
}

func NewWindowsFromString(path string) *Windows {
	if len(path) == 0 {
		log.Fatal("Cannot create path.")
	}
	return &Windows{
		path: path,
	}
}

func (p *Windows) Anchor() string {
	drive := p.Drive()
	if drive != "" {
		return drive + windowsSeparator
	}
	if unicode.IsLetter(rune(p.path[0])) && (string(p.path[1]) == "/" || string(p.path[1]) == "\\") {
		return p.path[:2]
	}
	return ""
}

func (p *Windows) AsPosix() string {
	return strings.ReplaceAll(p.path, windowsSeparator, posixSeparator)
}

// A string representing the drive letter or name, if any.
func (p *Windows) Drive() string {
	beginsWith, driveLetter := driveLetter(p.path)
	if !beginsWith {
		return ""
	}
	return driveLetter
}

func (p *Windows) Name() string {
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
func (p *Windows) Parent() PurePath {
	return NewWindowsFromString(path.Dir(p.path))
}

// A slice giving access to the path's various components.
func (p *Windows) Parts() []string {
	return strings.Split(p.path, windowsSeparator)
}

// A string representing the (local or global) root, if any.
func (p *Windows) Root() string {
	beginsWith, _ := driveLetter(p.path)
	if !beginsWith {
		return ""
	}
	return windowsSeparator
}

func (p *Windows) Stem() string {
	name := p.Name()
	ext := p.Suffix()

	before, found := strings.CutSuffix(name, ext)
	if !found {
		log.Fatalf("Could not find %s in %s", ext, name)
	}
	return before
}

func (p *Windows) Suffix() string {
	return path.Ext(p.path)
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
