package purepath_test

import (
	// "fmt"
	// "path/filepath"
	"testing"

	"github.com/rewgs/go-pathlib/purepath"
	"github.com/rewgs/go-pathlib/tests/utils"
)

func TestPurePosixPathAnchor(t *testing.T) {
	// testPath := utils.GetTestsPath()
	testPath := t.TempDir()
	utils.MakeDir(testPath)

	path, err := purepath.New(testPath)
	if err != nil {
		t.Error(err)
	}

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
	// testPath := utils.GetTestsPath()
	testPath := t.TempDir()
	utils.MakeDir(testPath)

	path, err := purepath.New(testPath)
	if err != nil {
		t.Error(err)
	}

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
	// testPath := utils.GetTestsPath()
	testPath := t.TempDir()
	utils.MakeDir(testPath)

	path, err := purepath.New(testPath)
	if err != nil {
		t.Error(err)
	}

	_, ok := path.(purepath.PurePosixPath)
	if !ok {
		t.Skip()
	}

	isAbsolute := path.IsAbsolute()
	if !isAbsolute {
		t.Errorf("TestPurePosixPathIsAbsolute(): Wanted: %t; got: %t\n", true, isAbsolute)
	}
}

// FIXME:
//
// func TestPurePosixPathParent(t *testing.T) {
// 	// testPath := utils.GetTestsPath()
// 	testPath := t.TempDir()
// 	utils.MakeDir(testPath)
//
// 	path, err := purepath.New(testPath)
// 	if err != nil {
// 		t.Error(err)
// 	}
//
// 	_, ok := path.(purepath.PurePosixPath)
// 	if !ok {
// 		t.Skip()
// 	}
//
// 	parentPath := purepath.NewPurePosixPath(utils.GetHome(), ".local", "share", "go-pathlib")
// 	parent := path.Parent()
// 	if parent != parentPath {
// 		t.Errorf("TestPurePosixPathParent(): Wanted: %s; got: %s\n", filepath.Dir(testPath), parent.AsString())
// 	}
// }

// FIXME:
//
// func TestPurePosixPathParts(t *testing.T) {
// 	testPart := func(got string, want string) {
// 		if want != got {
// 			t.Errorf("TestPurePosixPathParts(): Want: %s; got: %s\n", want, got)
// 		}
// 	}
//
// 	// testPath := utils.GetTestsPath()
// 	testPath := t.TempDir()
// 	utils.MakeDir(testPath)
//
// 	path, err := purepath.New(testPath)
// 	if err != nil {
// 		t.Error(err)
// 	}
//
// 	_, ok := path.(purepath.PurePosixPath)
// 	if !ok {
// 		t.Skip()
// 	}
//
// 	parts := path.Parts()
// 	for _, part := range parts {
// 		fmt.Println(part)
// 	}
//
// 	testPart(parts[0], "")
// 	testPart(parts[1], "home")
// 	testPart(parts[2], utils.GetUsername())
// 	testPart(parts[3], ".local")
// 	testPart(parts[4], "share")
// 	testPart(parts[5], "go-pathlib")
// 	testPart(parts[6], "tests")
// }

func TestPurePosixPathRoot(t *testing.T) {
	// testPath := utils.GetTestsPath()
	testPath := t.TempDir()
	utils.MakeDir(testPath)

	path, err := purepath.New(testPath)
	if err != nil {
		t.Error(err)
	}

	_, ok := path.(purepath.PurePosixPath)
	if !ok {
		t.Skip()
	}

	root := path.Root()
	if root != "/" {
		t.Errorf("TestPurePosixPathRoot(): Wanted: %s; got: %s\n", "/", root)
	}
}
