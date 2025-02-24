package path

// Path: A representation of a filepath.
type Path interface {
	Exists() bool
}
