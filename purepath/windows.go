package purepath

import (
	"log"
	"path" // TODO: Test if this works with "\\" paths.
	"strings"
	"unicode"

	"github.com/rewgs/go-pathlib/internal/posix"
	"github.com/rewgs/go-pathlib/internal/windows"
)

type PureWindowsPath struct {
	path string
}

// Takes any number of strings, separated by commas.
func NewPureWindowsPath(pathsegments ...string) *PureWindowsPath {
	if len(pathsegments) == 0 {
		log.Fatal("Cannot create path.")
	}
	return &PureWindowsPath{
		path: strings.Join(pathsegments, windows.Separator),
	}
}

func (p *PureWindowsPath) Anchor() string {
	drive := p.Drive()
	if drive != "" {
		return drive + windows.Separator
	}
	if unicode.IsLetter(rune(p.path[0])) && (string(p.path[1]) == "/" || string(p.path[1]) == "\\") {
		return p.path[:2]
	}
	return ""
}

func (p *PureWindowsPath) AsPosix() string {
	return strings.ReplaceAll(p.path, windows.Separator, posix.Separator)
}

// A string representing the drive letter or name, if any.
func (p *PureWindowsPath) Drive() string {
	beginsWith, driveLetter := driveLetter(p.path)
	if !beginsWith {
		return ""
	}
	return driveLetter
}

func (p *PureWindowsPath) Name() string {
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
func (p *PureWindowsPath) Parent() PurePath {
	return NewPureWindowsPath(path.Dir(p.path))
}

// A slice giving access to the path's various components.
func (p *PureWindowsPath) Parts() []string {
	return strings.Split(p.path, windows.Separator)
}

// A string representing the (local or global) root, if any.
func (p *PureWindowsPath) Root() string {
	beginsWith, _ := driveLetter(p.path)
	if !beginsWith {
		return ""
	}
	return windows.Separator
}

func (p *PureWindowsPath) Stem() string {
	name := p.Name()
	ext := p.Suffix()

	before, found := strings.CutSuffix(name, ext)
	if !found {
		log.Fatalf("Could not find %s in %s", ext, name)
	}
	return before
}

func (p *PureWindowsPath) Suffix() string {
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
