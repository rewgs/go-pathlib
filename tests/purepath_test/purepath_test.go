package purepath_test

import (
	"path/filepath"
	"testing"

	"github.com/rewgs/go-pathlib/purepath"
)

func TestPurePath(t *testing.T) {
	testDir := t.TempDir()
	testPath := purepath.New(testDir)

	// Type assertion
	// https://go.dev/tour/methods/15
	_, ok := testPath.(purepath.PurePath)
	if !ok {
		// t.Skip()
		t.Error()
	}
}
func TestPathIsAbsolute(t *testing.T) {
	testDir := t.TempDir()
	testPath := purepath.New(testDir)
	if !testPath.IsAbsolute() {
		t.Errorf("TestPathIsAbsolute(): %s is not absolute", testPath.AsString())
	}
}

func TestPurePathDrive(t *testing.T) {
	testDir := t.TempDir()
	testPath := purepath.New(testDir)
	drive := testPath.Drive()
	if drive != "" {
		t.Errorf("TestPurePathDrive(): Wanted: %s; got: %s\n", "", drive)
	}
}

func TestPurePathIsAbsolute(t *testing.T) {
	testDir := t.TempDir()
	testPath := purepath.New(testDir)
	isAbsolute := testPath.IsAbsolute()
	if !isAbsolute {
		t.Errorf("TestPurePathIsAbsolute(): Wanted: %t; got: %t\n", true, isAbsolute)
	}
}

func TestPurePathParent(t *testing.T) {
	testDir := t.TempDir()
	testPath := purepath.New(testDir)
	parentPath := purepath.New(filepath.Dir(testDir))
	if testPath.Parent() != parentPath {
		t.Errorf("TestPurePathParent(): Wanted: %s; got: %s\n", filepath.Dir(testDir), parentPath.AsString())
	}
}

func TestPurePathRoot(t *testing.T) {
	testDir := t.TempDir()
	testPath := purepath.New(testDir)
	root := testPath.Root()
	if root != "/" {
		t.Errorf("TestPurePathRoot(): Wanted: %s; got: %s\n", "/", root)
	}
}
