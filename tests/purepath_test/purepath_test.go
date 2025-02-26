package purepath_test

import (
	"testing"

	"github.com/rewgs/go-pathlib/purepath"
	"github.com/rewgs/go-pathlib/tests/utils"
)

var testPath string = utils.GetTestsPath()

// Simple test creating the Path struct. Does not affect I/O  beyond utils.MakeDir().
func TestPurePath(t *testing.T) {
	utils.MakeDir(testPath)

	_, err := purepath.New(testPath)
	if err != nil {
		t.Error(err)
	}
}

func TestPurePosixPathAnchor(t *testing.T) {
	utils.MakeDir(testPath)

	path, err := purepath.New(testPath)
	if err != nil {
		t.Error(err)
	}

	_, ok := path.(purepath.PurePosixPath)
	if !ok {
		t.Skip()
	}

	if path.Anchor() != "/" {
		t.Errorf("TestPurePosixPathAnchor(): Wanted: %s; got: %s\n", "/", path.Anchor())
	}
}

func TestPureWindowsPathAnchor(t *testing.T) {
	utils.MakeDir(testPath)

	path, err := purepath.New(testPath)
	if err != nil {
		t.Error(err)
	}

	_, ok := path.(purepath.PureWindowsPath)
	if !ok {
		t.Skip()
	}

	if path.Anchor() == "" {
		t.Error()
	}
}
