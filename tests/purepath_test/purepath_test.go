package purepath_test

import (
	"testing"

	"github.com/rewgs/go-pathlib/purepath"
)

// Simple test creating the Path struct. Does not affect I/O beyond the use of testing T.TempDir().
// func TestPurePath(t *testing.T) {
// 	testPath := t.TempDir()
//
// 	_, err := purepath.PurePath(testPath)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

func TestPathIsAbsolute(t *testing.T) {
	testPath := t.TempDir()

	p := purepath.New(testPath)

	if !p.IsAbsolute() {
		t.Errorf("TestPathIsAbsolute(): %s is not absolute", p.AsString())
	}
}
