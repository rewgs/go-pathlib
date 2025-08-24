package path

import (
	"github.com/rewgs/go-pathlib/purepath"
)

type base struct {
	purepath.PurePath
}

// newBase() returns a base with an embedded PurePath struct matching the current runtime.GOOS.
func newBase(pathsegments ...string) base {
	p := purepath.New(pathsegments...)
	b := base{p}
	return b
}

func (b base) someThing() {
}
