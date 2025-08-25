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
	filepath string
}

func (p base) AsString() string {
	return p.filepath
}

func (p base) Name() string {
	name := path.Base(p.filepath)
	if name == "." || name == "/" {
		log.Fatalf("Could not get name from %s", p.filepath)
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
	return path.Ext(p.filepath)
}

// TODO: Impelement similar functionality: "If the original path doesn’t have a name, ValueError is raised."
// TODO:
func (p base) WithName(name string) PurePath {
	panic("PurePath.WithName() not yet implemented.")
}

// TODO:
func (p base) FullMatch() bool {
	panic("PurePath.FullMatch() not yet implemented.")
}

// TODO:
func (p base) IsRelativeTo() bool {
	panic("PurePath.IsRelativeTo() not yet implemented.")
}

// TODO:
func (p base) IsReserved() bool {
	panic("PurePath.IsReserved() not yet implemented.")
}

// TODO:
func (p base) Match() bool {
	panic("PurePath.Match() not yet implemented.")
}

// TODO:
func (p base) Parents() []PurePath {
	panic("PurePath.Parents() not yet implemented.")
}

// TODO:
func (p base) RelativeTo() PurePath {
	panic("PurePath.RelativeTo() not yet implemented.")
}

// TODO:
func (p base) Suffixes() []string {
	panic("PurePath.Suffixes() not yet implemented.")
}

// TODO:
func (p base) WithSegments() PurePath {
	panic("PurePath.WithSegments() not yet implemented.")
}

// TODO:
func (p base) WithStem() PurePath {
	panic("PurePath.WithStem() not yet implemented.")
}

// TODO:
func (p base) WithSuffix() PurePath {
	panic("PurePath.WithSuffix() not yet implemented.")
}
