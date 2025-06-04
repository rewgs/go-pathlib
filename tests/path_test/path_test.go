package path_test

import (
	"testing"

	"github.com/rewgs/go-pathlib/path"
)

// Simple test creating the Path struct. Does not affect I/O  beyond utils.MakeDir().
func TestPath(t *testing.T) {
	testPath := t.TempDir()

	_, err := path.New(testPath)
	if err != nil {
		t.Error(err)
	}
}

func TestPathAsString(t *testing.T) {
	testPath := t.TempDir()

	path, err := path.New(testPath)
	if err != nil {
		t.Error(err)
	}

	if path.AsString() != testPath {
		t.Errorf("TestPathAsString: wanted: %s; got: %s\n", testPath, path.AsString())
	}
}

func TestPathIsAbsolute(t *testing.T) {
	testPath := t.TempDir()

	path, err := path.New(testPath)
	if err != nil {
		t.Error(err)
	}

	if !path.IsAbsolute() {
		t.Errorf("TestPathIsAbsolute(): %s is not absolute", path.AsString())
	}
}
