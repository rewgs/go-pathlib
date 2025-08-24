package purepath

import (
	"log"
	"path"
	"strings"
)

// base forms the basis for the PurePosixPath and PureWindowsPath structs.
// Any methods that are shared between both flavors are implemented via this struct, whereas
// any methods that differ between flavors are implemented via their respective struct.
type base struct {
	Filepath string
}

func (p base) AsString() string {
	return p.Filepath
}

func (p base) Name() string {
	name := path.Base(p.Filepath)
	if name == "." || name == "/" {
		log.Fatalf("Could not get name from %s", p.Filepath)
	}
	return name
}

func (p base) Stem() string {
	name := p.Name()
	ext := p.Suffix()

	before, found := strings.CutSuffix(name, ext)
	if !found {
		log.Fatalf("Could not find %s in %s", ext, name)
	}
	return before
}

func (p base) Suffix() string {
	return path.Ext(p.Filepath)
}
