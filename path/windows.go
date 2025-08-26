package path

// WindowsPath is the flavor of Path which represents Windows filesystem paths.
type WindowsPath struct {
	base
}

// NewWindowsPath returns a newly-instantiated WindowsPath.
func NewWindowsPath(pathsegments ...string) WindowsPath {
	return WindowsPath{
		newBase(pathsegments...),
	}
}
