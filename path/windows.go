package path

import (
	"log"
	"strings"

	"github.com/rewgs/go-pathlib/internal/windows"
	"github.com/rewgs/go-pathlib/purepath"
)

type WindowsPath struct {
	Base
	purepath.PureWindowsPath
}

// Takes any number of strings, separated by commas.
func NewWindowsPath(pathsegments ...string) WindowsPath {
	if platform != "windows" {
		log.Panic()
	}

	filepath := strings.Join(pathsegments, windows.Separator)

	return WindowsPath{
		Base{Filepath: filepath},
		purepath.PureWindowsPath{
			purepath.Base{Filepath: filepath},
		},
	}
}

// TODO:
// func IsJunction() bool {}
