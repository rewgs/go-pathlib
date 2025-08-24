package path

import (
	"log"
	"os"
)

// Base forms the basis for the PosixPath and WindowsPath structs.
// Any methods that are shared between both flavors are implemented via this struct, whereas
// any methods that differ between flavors are implemented via their respective struct.
type Base struct {
}

func Home() Path {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	path := New(home)
	return path
}

// NOTE: Duplicating this in posix and windows.
//
// func (p Base) Exists() bool {
// 	fileInfo, err := os.Stat(p.Filepath)
// 	if errors.Is(err, os.ErrNotExist) {
// 		return false
// 	} else if err != nil && fileInfo == nil {
// 		log.Fatal(err)
// 	}
// 	return true
// }

// NOTE: Duplicating this in posix and windows.
//
// func (p Base) MkDir(mode *fs.FileMode, parents bool, existOK bool) error {
// 	if p.Exists() && !existOK {
// 		return os.ErrExist
// 	}
//
// 	if mode == nil {
// 		perm := fs.FileMode(0o777)
// 		mode = &perm
// 	}
//
// 	if parents {
// 		err := os.MkdirAll(p.Filepath, *mode)
// 		if err != nil {
// 			return err
// 		}
// 	} else {
// 		err := os.Mkdir(p.Filepath, *mode)
// 		if err != nil {
// 			return err
// 		}
// 	}
//
// 	return nil
// }
