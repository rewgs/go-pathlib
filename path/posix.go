package path

import (
	"strings"

	"github.com/rewgs/go-pathlib/internal/posix"
	"github.com/rewgs/go-pathlib/purepath"
)

type PosixPath struct {
	path
	*purepath.PurePosixPath
}

// Takes any number of strings, separated by commas.
func NewPosixPath(pathsegments ...string) *PosixPath {
	return &PosixPath{
		path{
			path: strings.Join(pathsegments, posix.Separator),
		},
		purepath.NewPurePosixPath(pathsegments...),
	}
}
