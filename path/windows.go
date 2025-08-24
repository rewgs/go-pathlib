package path

import (
	"errors"
	"io/fs"
	"log"
	"os"

	"github.com/rewgs/go-pathlib/purepath"
)

type WindowsPath struct {
	purepath.PureWindowsPath
}

// Takes any number of strings, separated by commas.
func NewWindowsPath(pathsegments ...string) WindowsPath {
	return WindowsPath{
		purepath.NewPureWindowsPath(),
	}
}

func (p WindowsPath) Exists() bool {
	fileInfo, err := os.Stat(p.Filepath)
	if errors.Is(err, os.ErrNotExist) {
		return false
	} else if err != nil && fileInfo == nil {
		log.Fatal(err)
	}
	return true
}

func (p WindowsPath) MkDir(mode *fs.FileMode, parents bool, existOK bool) error {
	if p.Exists() && !existOK {
		return os.ErrExist
	}

	if mode == nil {
		perm := fs.FileMode(0o777)
		mode = &perm
	}

	if parents {
		err := os.MkdirAll(p.Filepath, *mode)
		if err != nil {
			return err
		}
	} else {
		err := os.Mkdir(p.Filepath, *mode)
		if err != nil {
			return err
		}
	}

	return nil
}

// TODO:
// func IsJunction() bool {}
