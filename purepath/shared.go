package purepath

import (
	"log"
	"path"
	"strings"
)

// This forms the bases of the PurePosixPath and PureWindowsPath structs.
// Any methods that are shared between both flavors are implemented via this
// struct, whereas any methods that differ between flavors are implemented via
// their respective struct.
type Shared struct {
	Path string
}

func (p *Shared) AsString() string {
	return p.Path
}

func (p *Shared) Name() string {
	name := path.Base(p.Path)
	if name == "." || name == "/" {
		log.Fatalf("Could not get name from %s", p.Path)
	}
	return name
}

func (p *Shared) Stem() string {
	name := p.Name()
	ext := p.Suffix()

	before, found := strings.CutSuffix(name, ext)
	if !found {
		log.Fatalf("Could not find %s in %s", ext, name)
	}
	return before
}

func (p *Shared) Suffix() string {
	return path.Ext(p.Path)
}
