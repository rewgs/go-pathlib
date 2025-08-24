package purepath_test

import (
	"runtime"
	"testing"

	"github.com/rewgs/go-pathlib/purepath"
)

func TestPurePosixPathAnchor(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip()
	}

	testDir := t.TempDir()
	testPath := purepath.NewPurePosixPath(testDir)
	anchor := testPath.Anchor()
	if anchor != "/" {
		t.Errorf("TestPurePosixPathAnchor(): Wanted: %s; got: %s\n", "/", anchor)
	}
}
