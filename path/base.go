package path

import (
	"errors"
	"io/fs"
	"os"
)

type Base struct {
	Filepath string
}

func (p Base) Exists() (bool, error) {
	fileInfo, err := os.Stat(p.Filepath)
	if errors.Is(err, os.ErrNotExist) {
		return false, err
	} else if err != nil && fileInfo == nil {
		// This is a little sloppy, but for now it's fine.
		// TODO: Specifically handle other errors.
		return false, err
	}
	return true, nil
}

func (p Base) Home() (Path, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	return New(home)
}

func (p Base) MkDir(mode *fs.FileMode, parents bool, existOK bool) error {
	exists, err := p.Exists()

	// os.ErrNotExist is okay here, so it doesn't need to be returned or dealt with.
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	if exists && !existOK {
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
