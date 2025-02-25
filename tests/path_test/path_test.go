package path_test

import (
	"testing"

	"github.com/rewgs/go-pathlib/path"
	"github.com/rewgs/go-pathlib/tests/utils"
)

// Simple test creating the Path struct. Does not affect I/O  beyond utils.MakeDir().
func TestPath(t *testing.T) {
	testsPath := utils.GetTestsPath()
	utils.MakeDir(testsPath)

	_, err := path.New(testsPath)
	if err != nil {
		t.Error(err)
	}
}
