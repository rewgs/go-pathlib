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

// TODO: Support UNC paths.
//
// PureWindowsPath is the flavor of PurePath which represents Windows filesystem paths.
type PureWindowsPath struct {
	base
}

// NewPureWindowsPath returns a newly-instantiated PureWindowsPath.
func NewPureWindowsPath(pathsegments ...string) PureWindowsPath {
	if platform != "windows" {
		log.Panic()
	}

	if len(pathsegments) == 0 {
		log.Fatal("Cannot create path.")
	}

	pureWindowsPath := PureWindowsPath{}
	pureWindowsPath.filepath = strings.Join(pathsegments, windows.Separator)
	return pureWindowsPath
}

// Anchor returns the concatenation of the drive and root.
func (p PureWindowsPath) Anchor() string {
	drive := p.Drive()
	if drive != "" {
		return drive + windows.Separator
	}
	if unicode.IsLetter(rune(p.filepath[0])) && (string(p.filepath[1]) == "/" || string(p.filepath[1]) == "\\") {
		return p.filepath[:2]
	}
	return ""
}

// AsPosix return a string representation of the path with forward slashes (/).
func (p PureWindowsPath) AsPosix() string {
	return strings.ReplaceAll(p.filepath, windows.Separator, posix.Separator)
}

// Drive returns a string representing the drive letter or name, if any.
func (p PureWindowsPath) Drive() string {
	beginsWith, driveLetter := windows.GetDriveLetter(p.filepath)
	if !beginsWith {
		return ""
	}
	return driveLetter
}

// JoinPath returns a new PurePath which each of the pathsegments combined to the path.
func (p PureWindowsPath) JoinPath(pathsegments ...string) PurePath {
	path := slices.Concat(strings.Split(p.filepath, windows.Separator), pathsegments)
	return NewPurePosixPath(path...)
}

// IsAbsolute returns whether the path is absolute or not. A path is considered absolute if it has both a root and (if the flavour allows) a drive.
func (p PureWindowsPath) IsAbsolute() bool {
	if p.Drive() != "" && p.Root() != "" {
		return true
	}
	return false
}

// TODO:
//
// IsReserved returns true if the path is considered reserved under Windows.
func (p PureWindowsPath) IsReserved() bool {
	panic("PureWindowsPath.IsReserved() not yet implemented.")
}

// TODO:
// Account for the following at https://docs.python.org/3/library/pathlib.html#pathlib.PurePath.parent
// - "You cannot go past an anchor, or empty path."
// - "Note: This is a purely lexical operation..."
//
// The logical parent of the path.
func (p PureWindowsPath) Parent() PurePath {
	return NewPureWindowsPath(path.Dir(p.filepath))
}

// A slice giving access to the path's various components.
func (p PureWindowsPath) Parts() []string {
	return strings.Split(p.filepath, windows.Separator)
}

// A string representing the (local or global) root, if any.
func (p PureWindowsPath) Root() string {
	beginsWith, _ := windows.GetDriveLetter(p.filepath)
	if !beginsWith {
		return ""
	}
	return windows.Separator
}
