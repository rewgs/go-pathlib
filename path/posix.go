package path

import (
	"log"
	"strings"

	"github.com/rewgs/go-pathlib/internal/posix"
	"github.com/rewgs/go-pathlib/purepath"
)

type PosixPath struct {
	Shared
	purepath.PurePosixPath
}

// Takes any number of strings, separated by commas.
func NewPosixPath(pathsegments ...string) PosixPath {
	if platform == "windows" {
		log.Panic()
	}

	path := strings.Join(pathsegments, posix.Separator)

	return PosixPath{
		Shared{Path: path},
		purepath.PurePosixPath{
			purepath.Shared{Path: path},
		},
	}
}
