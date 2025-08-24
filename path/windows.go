package path

// type WindowsPath struct {
// 	purepath.PureWindowsPath
// }

// Takes any number of strings, separated by commas.
// func NewWindowsPath(pathsegments ...string) WindowsPath {
// 	return WindowsPath{
// 		purepath.NewPureWindowsPath(),
// 	}
// }

type WindowsPath struct {
	base
}

func NewWindowsPath(pathsegments ...string) WindowsPath {
	return WindowsPath{
		newBase(pathsegments...),
	}
}

// func (p WindowsPath) Exists() bool {
// 	fileInfo, err := os.Stat(p.AsString())
// 	if errors.Is(err, os.ErrNotExist) {
// 		return false
// 	} else if err != nil && fileInfo == nil {
// 		log.Fatal(err)
// 	}
// 	return true
// }
//
// func (p WindowsPath) MkDir(mode *fs.FileMode, parents bool, existOK bool) error {
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
// 		err := os.MkdirAll(p.AsString(), *mode)
// 		if err != nil {
// 			return err
// 		}
// 	} else {
// 		err := os.Mkdir(p.AsString(), *mode)
// 		if err != nil {
// 			return err
// 		}
// 	}
//
// 	return nil
// }

// TODO:
// func IsJunction() bool {}
