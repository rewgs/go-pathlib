package purepath_test

import (
	// "fmt"
	"path/filepath"
	"testing"

	"github.com/rewgs/go-pathlib/purepath"
)

func TestPurePosixPathAnchor(t *testing.T) {
	testPath := t.TempDir()

	path := purepath.New(testPath)

	_, ok := path.(purepath.PurePosixPath)
	if !ok {
		t.Skip()
	}

	anchor := path.Anchor()
	if anchor != "/" {
		t.Errorf("TestPurePosixPathAnchor(): Wanted: %s; got: %s\n", "/", anchor)
	}
}

func TestPurePosixPathDrive(t *testing.T) {
	testPath := t.TempDir()

	path := purepath.New(testPath)

	_, ok := path.(purepath.PurePosixPath)
	if !ok {
		t.Skip()
	}

	drive := path.Drive()
	if drive != "" {
		t.Errorf("TestPurePosixPathDrive(): Wanted: %s; got: %s\n", "", drive)
	}
}

func TestPurePosixPathIsAbsolute(t *testing.T) {
	testPath := t.TempDir()

	path := purepath.New(testPath)

	_, ok := path.(purepath.PurePosixPath)
	if !ok {
		t.Skip()
	}

	isAbsolute := path.IsAbsolute()
	if !isAbsolute {
		t.Errorf("TestPurePosixPathIsAbsolute(): Wanted: %t; got: %t\n", true, isAbsolute)
	}
}

func TestPurePosixPathParent(t *testing.T) {
	testPath := t.TempDir()

	p := purepath.New(testPath)

	// Type assertion
	// https://go.dev/tour/methods/15
	_, ok := p.(purepath.PurePosixPath)
	if !ok {
		// t.Skip()
		t.Error()
	}

	parent := purepath.NewPurePosixPath(filepath.Dir(testPath))
	if p.Parent() != parent {
		t.Errorf("TestPurePosixPathParent(): Wanted: %s; got: %s\n", filepath.Dir(testPath), parent.AsString())
	}
}

func TestPurePosixPathRoot(t *testing.T) {
	testPath := t.TempDir()

	path := purepath.New(testPath)

	_, ok := path.(purepath.PurePosixPath)
	if !ok {
		t.Skip()
	}

	root := path.Root()
	if root != "/" {
		t.Errorf("TestPurePosixPathRoot(): Wanted: %s; got: %s\n", "/", root)
	}
}
