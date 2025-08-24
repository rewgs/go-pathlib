package path

type PosixPath struct {
	base
}

func NewPosixPath(pathsegments ...string) PosixPath {
	return PosixPath{
		newBase(pathsegments...),
	}
}
