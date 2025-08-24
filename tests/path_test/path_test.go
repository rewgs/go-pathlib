package path_test

import (
	"testing"

	"github.com/rewgs/go-pathlib/path"
)

func TestPath(t *testing.T) {
	testPath := t.TempDir()
	p := path.New(testPath)

	// Type assertion
	// https://go.dev/tour/methods/15
	_, ok := p.(path.Path)
	if !ok {
		// t.Skip()
		t.Error()
	}
}

func TestPathAsString(t *testing.T) {
	testDir := t.TempDir()
	testPath := path.New(testDir)
	if testPath.AsString() != testDir {
		t.Errorf("TestPathAsString: wanted: %s; got: %s\n", testDir, testPath.AsString())
	}
}

func TestPathExists(t *testing.T) {
	testDir := t.TempDir()
	testPath := path.New(testDir)
	if !testPath.Exists() {
		t.Errorf("TestPathExists: testPath does not exist\n")
	}
}

// NOTE: In progress. Need to finish purepath.JoinPath().
//
// func TestMkdir(t *testing.T) {
// 	testDir := t.TempDir()
// 	testPath := path.New(testDir)
//
// 	err := testPath.MkDir(nil, false, false)
// 	if err != nil {
// 	}
// }

// func TestPathIsAbsolute(t *testing.T) {
// 	testPath := t.TempDir()
//
// 	path, err := path.New(testPath)
// 	if err != nil {
// 		t.Error(err)
// 	}
//
// 	if !path.IsAbsolute() {
// 		t.Errorf("TestPathIsAbsolute(): %s is not absolute", path.AsString())
// 	}
// }
