package path

type WindowsPath struct {
	base
}

func NewWindowsPath(pathsegments ...string) WindowsPath {
	return WindowsPath{
		newBase(pathsegments...),
	}
}
