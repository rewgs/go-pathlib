package path

import (
	"log"
	"strings"

	"github.com/rewgs/go-pathlib/internal/windows"
	"github.com/rewgs/go-pathlib/purepath"
)

type WindowsPath struct {
	Shared
	purepath.PureWindowsPath
}

// Takes any number of strings, separated by commas.
func NewWindowsPath(pathsegments ...string) WindowsPath {
	if platform != "windows" {
		log.Panic()
	}

	path := strings.Join(pathsegments, windows.Separator)

	return WindowsPath{
		Shared{Path: path},
		purepath.PureWindowsPath{
			purepath.Shared{Path: path},
		},
	}
}

// TODO:
// func IsJunction() bool {}
