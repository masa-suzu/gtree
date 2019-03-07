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

func TestInsert(t *testing.T) {

	tests := []struct {
		name string
		want []kv
	}{
		{
			name: "integers",
			want: []kv{
				{k: 2, v: 100},
				{k: 1, v: 200},
			},
		},
		{
			name: "strings",
			want: []kv{
				{k: 1, v: "200"},
				{k: 2, v: "100"},
			},
		},
		{
			name: "integers_and_strings",
			want: []kv{
				{k: 4, v: 100},
				{k: 3, v: 200},
				{k: 1, v: "100"},
				{k: 2, v: "200"},
			},
		},
	}

	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := avl.New()
			for _, kv := range tt.want {
				tree.Insert(kv.k, kv.v)
			}
			assertTree(t, tree, tt.want)
		})
	}
}

func assertTree(t *testing.T, tree *avl.Tree, kvs []kv) {
	if len(kvs) != tree.Count() {
		t.Errorf("num of nodes must be %v, got %v", len(kvs), tree.Count())
	}

	for _, kv := range kvs {
		got, err := tree.Search(kv.k)

		if err != nil {
			t.Errorf("got an error '%v'", err)
		}
		if kv.v != got {
			t.Errorf("want %v, got %v", kv.v, got)
		}
	}
}

func TestInsert_with_SameKeys(t *testing.T) {

	tests := []struct {
		name   string
		kvs    []kv
		unique int
		want   interface{}
	}{
		{
			name: "Want_Last_Inserted",
			kvs: []kv{
				{k: 1, v: 300},
				{k: 1, v: 100},
				{k: 1, v: 200},
			},
			unique: 1,
			want:   200,
		},
	}

	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := avl.New()

			for i := 0; i < len(tt.kvs); i++ {
				tree.Insert(tt.kvs[i].k, tt.kvs[i].v)
			}

			if tt.unique != tree.Count() {
				t.Errorf("num of nodes must be %v, got %v", tt.unique, tree.Count())
			}

			got, err := tree.Search(tt.kvs[0].k)

			if err != nil {
				t.Errorf("got an error '%v'", err)
			}
			if tt.want != got {
				t.Errorf("want %v, got %v", tt.want, got)
			}
		})
	}
}

func TestSearch_by_InvalidKey(t *testing.T) {

	in := kv{k: 1, v: 100}

	tree := avl.New()
	tree.Insert(in.k, in.v)

	got, err := tree.Search(100)

	if err == nil {
		t.Errorf("got no error, want '%v'", err)
	}

	if got != nil {
		t.Errorf("got %v, want %v", got, nil)
	}

}
