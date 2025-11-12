package path

// WindowsPath is the flavor of Path which represents Windows filesystem paths.
type WindowsPath struct {
	*base
}

// NewWindowsPath returns a newly-instantiated WindowsPath.
func NewWindowsPath(pathsegments ...string) *WindowsPath {
	return &WindowsPath{
		newBase(pathsegments...),
	}
}

// Return true if the path points to a junction, and false for any other type of file.
// Currently only Windows supports junctions.
func (p *WindowsPath) IsJunction() bool {
	panic("WindowsPath.IsJunction() not yet implemented.")
}
