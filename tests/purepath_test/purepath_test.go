package purepath_test

import (
	"testing"

	"github.com/rewgs/go-pathlib/purepath"
	"github.com/rewgs/go-pathlib/tests/utils"
)

// Simple test creating the Path struct. Does not affect I/O beyond the use of testing T.TempDir().
func TestPurePath(t *testing.T) {
	testPath := t.TempDir()
	utils.MakeDir(testPath)

	_, err := purepath.New(testPath)
	if err != nil {
		t.Error(err)
	}
}
