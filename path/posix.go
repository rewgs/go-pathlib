package path

// PosixPath is the flavor of Path which represents non-Windows filesystem paths.
type PosixPath struct {
	*base
}

// NewPosixPath returns a newly-instantiated PosixPath.
func NewPosixPath(pathsegments ...string) *PosixPath {
	return &PosixPath{
		newBase(pathsegments...),
	}
}
