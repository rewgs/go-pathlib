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

// TODO: Ensure that filepath arg is a valid path string.
//
// newBase returns a base with a guaranteed valid filepath field.
// func newBase(filepath string) base {
// }

// AsString returns the Path as a string.
func (b *base) AsString() string {
	return b.filepath
}

// Name returns a string representing the final path component, excluding the drive and root, if any.
func (b *base) Name() string {
	name := path.Base(b.filepath)
	if name == "." || name == "/" {
		log.Fatalf("Could not get name from %s", b.filepath)
	}
	return name
}

// Stem returns the final path component, without its suffix.
func (b *base) Stem() string {
	name := b.Name()
	ext := b.Suffix()

	before, found := strings.CutSuffix(name, ext)
	if !found {
		log.Fatalf("Could not find %s in %s", ext, name)
	}
	return before
}

// Suffix returns the last dot-separated portion of the final component, if any.
func (b *base) Suffix() string {
	return path.Ext(b.filepath)
}

// TODO: Implement similar functionality: "If the original path doesn’t have a name, ValueError is raised."
// TODO:
//
// WithName returns a new PurePath with the name changed.
func (b *base) WithName(name string) PurePath {
	panic("PurePath.WithName() not yet implemented.")
}

// TODO:
//
// FullMatch matches this path against the provided glob-style pattern. Returns true if matching is successful, false otherwise.
// Case-sensitivity follows platform defaults.
func (b *base) FullMatch(pattern string, caseSensitive bool) bool {
	panic("PurePath.FullMatch() not yet implemented.")
}

// TODO:
//
// IsRelativeTo returns whether or not this path is relative to the other path.
// This method is string-based; it neither accesses the filesystem nor treats additionalPaths specially (if supplied, they are joined with other).
func (b *base) IsRelativeTo(other string, additionalPaths *[]string) bool {
	panic("PurePath.IsRelativeTo() not yet implemented.")
}

// TODO:
//
// Match matches this path against the provided non-recursive glob-style pattern. Returns true if matching is successful, false otherwise.
// This function is similar to PurePath.FullMatch(), but empty patterns aren't allowed.
func (b *base) Match(pattern string, caseSensitive bool) bool {
	panic("PurePath.Match() not yet implemented.")
}

// TODO:
//
// Parents returns an immutable sequence providing access to th elogical ancestors of the path.
func (b *base) Parents() []PurePath {
	panic("PurePath.Parents() not yet implemented.")
}

// TODO: Implement similar functionality:
// - "If it’s impossible, ValueError is raised."
// - "In all other cases, such as the paths referencing different drives, ValueError is raised."
// TODO:
//
// RelativeTo computes a version of this PurePath relative to the path represented by other.
// When walkUp is false, the path must start with other.
// When walkUp is true, n additional paths may be added to form the relative path.
// Warning: this function is part of PurePath and works with strings. It does not check or access the underlying file structure.
// This can impact the walkUp option as it assumes that no symlinks are present in the path; call path.Resolve() first if necessary to resolve symlinks.
func (b *base) RelativeTo(other string, walkUp bool, additionalPaths *[]string) PurePath {
	panic("PurePath.RelativeTo() not yet implemented.")
}

// TODO:
//
// Suffixes returns a list of the path’s suffixes, often called file extensions.
func (b *base) Suffixes() []string {
	panic("PurePath.Suffixes() not yet implemented.")
}

// WithSegments Creates a new PurePath by combining the given pathsegments.
// This function simply calls purepath.New() and is thus redundant; it is only included for compatability with Python's pathlib module.
func (b *base) WithSegments(pathsegments ...string) PurePath {
	return New(pathsegments...)
}

// TODO: Implement similar functionality: "If the original path doesn’t have a name, ValueError is raised."
// TODO:
//
// WithStem returns a new PurePath with the stem changed.
func (b *base) WithStem(stem string) PurePath {
	panic("PurePath.WithStem() not yet implemented.")
}

// TODO:
//
// WithSuffix returns a new PurePath with the suffix changed.
// If the original path doesn’t have a suffix, the new suffix is appended instead.
// If the suffix is an empty string, the original suffix is removed.
func (b *base) WithSuffix(suffix string) PurePath {
	panic("PurePath.WithSuffix() not yet implemented.")
}
