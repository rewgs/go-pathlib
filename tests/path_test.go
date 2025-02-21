package path_test

import (
	"log"
	"os"
	"testing"

	path "github.com/rewgs/go-pathlib"
	// "github.com/rewgs/go-pathlib/tests/utils"
)

// func TestPathExists(t *testing.T) {
// 	p, err := path.New(utils.GetTestsPath())
// }

func TestPathAsString(t *testing.T) {
	home := func() string {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		return home
	}

	p, err := path.New(home())
	if err != nil {
		log.Fatal(err)
	}

	want := "/Users/rewgs"
	got := p.AsString()
	if want != got {
		t.Errorf("path.Parent(): want: %s; got: %s", want, got)
	}
}

func TestPathParent(t *testing.T) {
	home := func() string {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		return home
	}

	p, err := path.New(home())
	if err != nil {
		log.Fatal(err)
	}

	want := "/Users"
	got := p.Parent().AsString()
	if want != got {
		t.Errorf("path.Parent(): want: %s; got: %s", want, got)
	}
}

// Not sure if this is working
// func TestPathParents(t *testing.T) {
// 	fmt.Println("Running TestPathParents()")
//
// 	home := func() string {
// 		home, err := os.UserHomeDir()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		return home
// 	}
//
// 	// /Users/rewgs
// 	p, err := path.New(home())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	// ["/", "/Users"]
// 	parents := p.Parents()
// 	if parents[0].AsString() != p.Root {
// 		t.Error()
// 	}
// }
