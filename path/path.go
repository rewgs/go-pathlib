package path

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type Path struct {
	Root      string
	path      string
	separator string
}

func New(path string) (*Path, error) {
	p := Path{path: path}

	// TODO: Implement and test other operating systems listed here:
	// https://go.dev/doc/install/source#environment
	switch os := runtime.GOOS; os {
	case "darwin":
		p.separator = "/"
		p.Root = "/"
	case "linux":
		p.Root = "/"
		p.separator = "/"
	case "windows":
		p.Root = "C:"
		p.separator = "\\"
	default:
		return nil, fmt.Errorf("operating system not yet implemented or tested: %s", os)
	}

	return &p, nil
}

// TODO:
// - `follow_symlinks` argument.
// - Check for permission-related errors in cases where file exists.
//
// Return true if the path points to an existing file or directory.
// This function normally follows symlinks; to check if a symlink exists, add the argument follow_symlinks=False
func (p *Path) Exists() (bool, error) {
	fileInfo, err := os.Stat(p.path)
	if errors.Is(err, os.ErrNotExist) {
		return false, err
	} else if err != nil && fileInfo == nil {
		// This is a little sloppy. Need to specifically handle other errors.
		return false, err
	}
	return true, nil
}

func (p *Path) Parent() *Path {
	return &Path{
		path:      filepath.Dir(p.path),
		separator: p.separator,
	}
}

// FIXME: This is just causing an infinite loop because `parent := p.Parent()`
// isn't assigning the next parent's parent in the next iteration.
func (p *Path) Parents() (parents []*Path) {
	parent := p.Parent()
	for {
		if parent.AsString() == p.Root {
			parents = append(parents, parent)
			return
		}
		parents = append(parents, parent)
		parent = p.Parent()
	}
}

func (p *Path) PrintParents() {
	for _, parent := range p.Parents() {
		fmt.Println(parent.AsString())
	}
}

// TODO:
// func  (p *Path) Resolve(strict bool) *Path {}

func (p *Path) AsString() string {
	return p.path
}
