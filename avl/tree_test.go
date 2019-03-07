package avl_test

import (
	"testing"

	"github.com/masa-suzu/gtree/avl"
)

type kv struct {
	k int
	v interface{}
}

func TestNewTree(t *testing.T) {

	want := 0
	got := avl.New().Count()

	if want != got {
		t.Errorf("num of nodes must be %v, got %v", want, got)
	}
}
