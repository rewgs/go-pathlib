package purepath_test

import (
	"runtime"
	"testing"

	"github.com/rewgs/go-pathlib/purepath"
)

func TestPureWindowsPathAnchor(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip()
	}

	testDir := t.TempDir()
	testPath := purepath.NewPureWindowsPath(testDir)
	got := testPath.Anchor()
	want := testPath.Drive()
	if got != "/" {
		t.Errorf("TestPurePosixPathAnchor(): Wanted: %s; got: %s\n", want, got)
	}
}
