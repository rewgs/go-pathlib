package path

import (
	"log"
	"strings"

	"github.com/rewgs/go-pathlib/internal/windows"
	"github.com/rewgs/go-pathlib/purepath"
)

type WindowsPath struct {
	path
	// *purepath.PureWindowsPath
	purepath.PureWindowsPath
}

// Takes any number of strings, separated by commas.
func NewWindowsPath(pathsegments ...string) WindowsPath {
	if platform != "windows" {
		log.Panic()
	}

	return WindowsPath{
		path{
			path: strings.Join(pathsegments, windows.Separator),
		},
		purepath.NewPureWindowsPath(pathsegments...),
	}
}

// TODO:
// func IsJunction() bool {}
