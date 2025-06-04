package path

import (
	"log"
	"strings"

	"github.com/rewgs/go-pathlib/internal/windows"
	"github.com/rewgs/go-pathlib/purepath"
)

type WindowsPath struct {
	PathBase
	purepath.PureWindowsPath
}

// Takes any number of strings, separated by commas.
func NewWindowsPath(pathsegments ...string) WindowsPath {
	if platform != "windows" {
		log.Panic()
	}

	filepath := strings.Join(pathsegments, windows.Separator)

	return WindowsPath{
		PathBase{Filepath: filepath},
		purepath.PureWindowsPath{
			purepath.SharedPosixPath{Filepath: filepath},
		},
	}
}

// TODO:
// func IsJunction() bool {}
