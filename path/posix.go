package path

import (
	"log"
	"strings"

	"github.com/rewgs/go-pathlib/internal/posix"
	"github.com/rewgs/go-pathlib/purepath"
)

type PosixPath struct {
	path
	// *purepath.PurePosixPath
	purepath.PurePosixPath
}

// Takes any number of strings, separated by commas.
func NewPosixPath(pathsegments ...string) PosixPath {
	if platform == "windows" {
		log.Panic()
	}

	return PosixPath{
		path{
			path: strings.Join(pathsegments, posix.Separator),
		},
		purepath.NewPurePosixPath(pathsegments...),
	}
}
