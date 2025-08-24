package path

import (
	"errors"
	"io/fs"
	"log"
	"os"

	"github.com/rewgs/go-pathlib/purepath"
)

// base forms the basis for the PosixPath and WindowsPath structs.
// It contains an embedded PurePath struct matching the appropriate operating system.
// Any methods that are shared between both flavors are implemented via this struct, whereas
// any methods that differ between flavors are implemented via their respective struct.
type base struct {
	purepath.PurePath
}

// newBase() returns a base with an embedded PurePath struct matching the current runtime.GOOS.
func newBase(pathsegments ...string) base {
	// p := purepath.New(pathsegments...)
	// b := base{p}
	// return b

	return base{
		purepath.New(pathsegments...),
	}
}

func (p base) Exists() bool {
	fileInfo, err := os.Stat(p.AsString())
	if errors.Is(err, os.ErrNotExist) {
		return false
	} else if err != nil && fileInfo == nil {
		log.Fatal(err)
	}
	return true
}

func (p base) MkDir(mode *fs.FileMode, parents bool, existOK bool) error {
	if p.Exists() && !existOK {
		return os.ErrExist
	}

	if mode == nil {
		perm := fs.FileMode(0o777)
		mode = &perm
	}

	if parents {
		err := os.MkdirAll(p.AsString(), *mode)
		if err != nil {
			return err
		}
	} else {
		err := os.Mkdir(p.AsString(), *mode)
		if err != nil {
			return err
		}
	}

	return nil
}
