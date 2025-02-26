package purepath

import (
	"log"
	"path" // TODO: Test if this works with "\\" paths.
	"slices"
	"strings"
	"unicode"

	"github.com/rewgs/go-pathlib/internal/posix"
	"github.com/rewgs/go-pathlib/internal/windows"
)

type PureWindowsPath struct {
	Shared
}

func NewPureWindowsPath(pathsegments ...string) PureWindowsPath {
	if len(pathsegments) == 0 {
		log.Fatal("Cannot create path.")
	}
	return PureWindowsPath{
		Shared{
			Path: strings.Join(pathsegments, windows.Separator),
		},
	}
}

func (p PureWindowsPath) Anchor() string {
	drive := p.Drive()
	if drive != "" {
		return drive + windows.Separator
	}
	if unicode.IsLetter(rune(p.Path[0])) && (string(p.Path[1]) == "/" || string(p.Path[1]) == "\\") {
		return p.Path[:2]
	}
	return ""
}

func (p PureWindowsPath) AsPosix() string {
	return strings.ReplaceAll(p.Path, windows.Separator, posix.Separator)
}

// A string representing the drive letter or name, if any.
func (p PureWindowsPath) Drive() string {
	beginsWith, driveLetter := windows.GetDriveLetter(p.Path)
	if !beginsWith {
		return ""
	}
	return driveLetter
}

func (p PureWindowsPath) JoinPath(pathsegments ...string) PurePath {
	path := slices.Concat(strings.Split(p.Path, windows.Separator), pathsegments)
	return NewPurePosixPath(path...)
}

func (p PureWindowsPath) IsAbsolute() bool {
	if p.Drive() != "" && p.Root() != "" {
		return true
	}
	return false
}

// TODO:
// Account for the following at https://docs.python.org/3/library/pathlib.html#pathlib.PurePath.parent
// - "You cannot go past an anchor, or empty path."
// - "Note: This is a purely lexical operation..."
//
// The logical parent of the path.
func (p PureWindowsPath) Parent() PurePath {
	return NewPureWindowsPath(path.Dir(p.Path))
}

// A slice giving access to the path's various components.
func (p PureWindowsPath) Parts() []string {
	return strings.Split(p.Path, windows.Separator)
}

// A string representing the (local or global) root, if any.
func (p PureWindowsPath) Root() string {
	beginsWith, _ := windows.GetDriveLetter(p.Path)
	if !beginsWith {
		return ""
	}
	return windows.Separator
}
