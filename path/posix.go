package path

import (
	"github.com/rewgs/go-pathlib/purepath"
)

type Posix struct {
	purepath.Posix
}

func NewPosix(path string) *Posix {
}

// TODO:
// - `follow_symlinks` argument.
// - Check for permission-related errors in cases where file exists.
//
// Return true if the path points to an existing file or directory.
// This function normally follows symlinks; to check if a symlink exists, add the argument follow_symlinks=False
// func (p *Posix) Exists() (bool, error) {
// 	fileInfo, err := os.Stat(p.path)
// 	if errors.Is(err, os.ErrNotExist) {
// 		return false, err
// 	} else if err != nil && fileInfo == nil {
// 		// This is a little sloppy, but for now it's fine.
// 		// TODO: Specifically handle other errors.
// 		return false, err
// 	}
// 	return true, nil
// }
