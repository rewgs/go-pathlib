package path

import (
	"github.com/rewgs/go-pathlib/purepath"
)

type Windows struct {
	purepath.Posix
}

func NewWindows(path string) *Windows {
}
