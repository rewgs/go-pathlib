package purepath

import (
	"log"
	"path"
	"strings"
)

// Base forms the basis for the PurePosixPath and PureWindowsPath structs.
// Any methods that are shared between both flavors are implemented via this struct, whereas
// any methods that differ between flavors are implemented via their respective struct.
type Base struct {
	Filepath string
}

func (p Base) AsString() string {
	return p.Filepath
}

func (p Base) Name() string {
	name := path.Base(p.Filepath)
	if name == "." || name == "/" {
		log.Fatalf("Could not get name from %s", p.Filepath)
	}
	return name
}

func (p Base) Stem() string {
	name := p.Name()
	ext := p.Suffix()

	before, found := strings.CutSuffix(name, ext)
	if !found {
		log.Fatalf("Could not find %s in %s", ext, name)
	}
	return before
}

func (p Base) Suffix() string {
	return path.Ext(p.Filepath)
}
