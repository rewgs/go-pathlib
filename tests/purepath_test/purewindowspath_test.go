package purepath_test

import (
	"testing"

	"github.com/rewgs/go-pathlib/purepath"
)

func TestPureWindowsPathAnchor(t *testing.T) {
	testPath := t.TempDir()

	path, err := purepath.New(testPath)
	if err != nil {
		t.Error(err)
	}

	_, ok := path.(purepath.PureWindowsPath)
	if !ok {
		t.Skip()
	}

	anchor := path.Anchor()
	if anchor == "" {
		t.Errorf("TestPureWindowsPathAnchor(): Wanted: %s; got: %s\n", "C:", anchor)
	}
}
