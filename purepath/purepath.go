// purepath provides purely computation operations with I/O.
package purepath

import (
	"fmt"
	"log"
	"path"
	"runtime"
	"strings"
)

var platform string = runtime.GOOS

// A generic interfaces that represents the system's path flavour (instantiating
// it creates either a PurePosixPath or PureWindowsPath).
type PurePath interface {
	// pathlib methods:
	Anchor() string
	Drive() string
	IsAbsolute() bool
	Name() string
	Parent() PurePath
	Parts() []string
	Root() string
	Stem() string
	Suffix() string
	// TODO:
	// FullMatch() bool
	// IsRelativeTo() bool
	// IsReserved() bool
	// JoinPath() PurePath
	// Match() bool
	// Parents() []PurePath
	// RelativeTo() PurePath
	// Suffixes() []string
	// WithName() PurePath
	// WithSegments() PurePath
	// WithStem() PurePath
	// WithSuffix() PurePath

	// added to purepath for go-pathlib:
	AsString() string
}

// This forms the bases of the PurePosixPath and PureWindowsPath structs.
// Any methods that are shared between both flavors are implemented via this
// struct, whereas any methods that differ between flavors are implemented via
// their respective struct.
type Shared struct {
	Path string
}

func (p Shared) AsString() string {
	return p.Path
}

func (p Shared) Name() string {
	name := path.Base(p.Path)
	if name == "." || name == "/" {
		log.Fatalf("Could not get name from %s", p.Path)
	}
	return name
}

func (p Shared) Stem() string {
	name := p.Name()
	ext := p.Suffix()

	before, found := strings.CutSuffix(name, ext)
	if !found {
		log.Fatalf("Could not find %s in %s", ext, name)
	}
	return before
}

func (p Shared) Suffix() string {
	return path.Ext(p.Path)
}

// Takes any number of strings, separated by commas.
// Returns either a PurePosixPath or PureWindowsPath, depending on `runtime.GOOS`.
func New(pathsegments ...string) (PurePath, error) {
	switch platform {
	case "darwin":
		return NewPurePosixPath(pathsegments...), nil
	case "linux":
		return NewPurePosixPath(pathsegments...), nil
	case "posix":
		return NewPurePosixPath(pathsegments...), fmt.Errorf("Operating system not yet implemented or tested: %s\nProceed with caution.\n", platform)
	case "windows":
		return NewPureWindowsPath(pathsegments...), nil
	default:
		panic(1)
	}
}
