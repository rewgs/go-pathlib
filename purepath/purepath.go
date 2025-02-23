// purepath provides purely computation operations with I/O.
package purepath

import (
	"fmt"
	// "log"
	"runtime"

	"github.com/rewgs/go-pathlib/purepath/posix"
	"github.com/rewgs/go-pathlib/purepath/windows"
)

// A generic interfaces that represents the system's path flavour (instantiating
// it creates either a PurePosixPath or PureWindowsPath).
type PurePath interface {
	Drive() string
	Parts() []string
	Root() string
}

// Takes any number of strings, separated by commas.
// Returns either a PurePosixPath or PureWindowsPath, depending on `runtime.GOOS`.
func New(pathsegments ...string) (PurePath, error) {
	switch platform := runtime.GOOS; platform {
	case "darwin":
		return posix.NewFromSlice(pathsegments), nil
	case "linux":
		return posix.NewFromSlice(pathsegments), nil
	case "posix":
		return posix.NewFromSlice(pathsegments), fmt.Errorf("NOTE: %s has not yet been implemented or tested. Proceed with caution.", platform)
	case "windows":
		return windows.NewFromSlice(pathsegments), nil
	default:
		panic(1)
	}
}
