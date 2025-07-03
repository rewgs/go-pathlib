package path

import (
	"log"
	"strings"

	"github.com/rewgs/go-pathlib/internal/posix"
	"github.com/rewgs/go-pathlib/purepath"
)

type PosixPath struct {
	Base
	purepath.PurePosixPath
}

// Takes any number of strings, separated by commas.
func NewPosixPath(pathsegments ...string) PosixPath {
	if platform == "windows" {
		log.Panic("NewPosixPath does not support Windows")
	}

	filepath := strings.Join(pathsegments, posix.Separator)

	return PosixPath{
		Base{
			Filepath: filepath,
		},
		purepath.PurePosixPath{
			purepath.Base{
				Filepath: filepath,
			},
		},
	}
}
