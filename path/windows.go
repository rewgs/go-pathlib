package path

import (
	"strings"

	"github.com/rewgs/go-pathlib/internal/windows"
	"github.com/rewgs/go-pathlib/purepath"
)

type WindowsPath struct {
	path
	*purepath.PureWindowsPath
}

// Takes any number of strings, separated by commas.
func NewWindowsPath(pathsegments ...string) *WindowsPath {
	return &WindowsPath{
		path{
			path: strings.Join(pathsegments, windows.Separator),
		},
		purepath.NewPureWindowsPath(pathsegments...),
	}
}
