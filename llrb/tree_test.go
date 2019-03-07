package llrb_test

import (
	"bytes"
	"testing"

	"github.com/masa-suzu/gtree/llrb"
)

type kv struct {
	k int
	v interface{}
}

func TestNewTree(t *testing.T) {

	want := 0
	got := llrb.New().Count()

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
			tree := llrb.New()
			for _, kv := range tt.want {
				tree.Insert(kv.k, kv.v)
			}
			assertTree(t, tree, tt.want)
		})
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
			tree := llrb.New()

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

	tree := llrb.New()
	tree.Insert(in.k, in.v)

	got, err := tree.Search(100)

	if err == nil {
		t.Errorf("got no error, want '%v'", err)
	}

	if got != nil {
		t.Errorf("got %v, want %v", got, nil)
	}

}

func TestDelete(t *testing.T) {
	tests := []struct {
		name    string
		deleted []kv
		want    []kv
	}{
		{
			name: "ascending",
			deleted: []kv{
				{k: 1, v: 200},
				{k: 2, v: 2400},
				{k: 3, v: 2040},
			},
			want: []kv{
				{k: 4, v: 100},
				{k: 5, v: 100},
			},
		},
		{
			name: "descending",
			deleted: []kv{
				{k: 5, v: 200},
				{k: 4, v: 2400},
				{k: 3, v: 2040},
			},
			want: []kv{
				{k: 2, v: 100},
				{k: 1, v: 100},
			},
		},
		{
			name: "random-ordering",
			deleted: []kv{
				{k: 6, v: 200},
				{k: 10, v: 2400},
				{k: 1, v: 2040},
				{k: 9, v: 2040},
				{k: 8, v: 2040},
				{k: 2, v: 2040},
			},
			want: []kv{
				{k: 4, v: 100},
				{k: 11, v: 100},
			},
		},
	}

	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := llrb.New()

			// insert all key-value pairs
			for _, kv := range tt.deleted {
				tree.Insert(kv.k, kv.v)
			}

			for _, kv := range tt.want {
				tree.Insert(kv.k, kv.v)
			}

			// delete nodes of tt.delete
			for _, kv := range tt.deleted {
				tree.Delete(kv.k)
			}
			assertTree(t, tree, tt.want)
		})
	}
}

func TestDelete_withSameKeys(t *testing.T) {
	tests := []struct {
		name     string
		inserted []kv
		deleted  []int
		want     []kv
	}{
		{
			name: "Insert-1to3-Delete-2",
			inserted: []kv{
				{k: 1, v: 300},
				{k: 2, v: 100},
				{k: 3, v: 200},
			},

			deleted: []int{
				2,
				2,
			},
			want: []kv{
				{k: 1, v: 300},
				{k: 3, v: 200},
			},
		},
	}

	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := llrb.New()
			for _, kv := range tt.inserted {
				tree.Insert(kv.k, kv.v)
			}

			for _, k := range tt.deleted {
				tree.Delete(k)
			}

			assertTree(t, tree, tt.want)
		})
	}

}

func TestWalk(t *testing.T) {
	tests := []struct {
		name     string
		inserted []kv
		want     []interface{}
	}{
		{
			name:     "empty",
			inserted: []kv{},
			want:     []interface{}{},
		},

		{
			name: "ascending",
			inserted: []kv{
				{k: 1, v: 100},
				{k: 2, v: 200},
				{k: 3, v: 300},
			},
			want: []interface{}{
				100, 200, 300,
			},
		},
		{
			name: "descending",
			inserted: []kv{
				{k: 5, v: 500},
				{k: 4, v: 400},
				{k: 3, v: 300},
			},
			want: []interface{}{
				300, 400, 500,
			},
		},
		{
			name: "random-ordering",
			inserted: []kv{
				{k: 6, v: 600},
				{k: 10, v: nil},
				{k: 1, v: 100},
				{k: 9, v: 900},
				{k: 8, v: 800},
				{k: 2, v: 200},
			},
			want: []interface{}{
				100, 200, 600, 800, 900, nil,
			},
		},
	}

	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := llrb.New()
			for _, kv := range tt.inserted {
				tree.Insert(kv.k, kv.v)
			}
			got := []interface{}{}
			for v := range tree.Walk() {
				got = append(got, v)
			}

			if len(tt.want) != len(got) {
				t.Errorf("num of nodes must be %v, got %v", len(tt.want), len(got))
			}

			for i, v := range tt.want {
				if v != got[i] {
					t.Errorf("want %v, got %v", v, got[i])
				}
			}
		})
	}
}

func TestToHTML(t *testing.T) {
	tests := []struct {
		name     string
		inserted []kv
		want     string
	}{
		{
			name:     "zero-node",
			inserted: []kv{},
			want: `<div class="tree">
</div>
`,
		},
		{
			name: "two-nodes",
			inserted: []kv{
				{k: 2, v: 600},
				{k: 1, v: 600},
			},
			want: `<div class="tree">
  <ul>
    <li>
      <black href="#">2/600</black>
      <ul>
        <li>
          <red href="#">1/600</red>
        </li>
      </ul>
    </li>
  </ul>
</div>
`,
		},
		{
			name: "six-nodes",
			inserted: []kv{
				{k: 10, v: 600},
				{k: 20, v: 600},
				{k: 30, v: 600},
				{k: 40, v: 600},
				{k: 50, v: 600},
				{k: 25, v: 600},
			},
			want: `<div class="tree">
  <ul>
    <li>
      <black href="#">40/600</black>
      <ul>
        <li>
          <red href="#">20/600</red>
          <ul>
            <li>
              <black href="#">10/600</black>
            </li>
            <li>
              <black href="#">30/600</black>
              <ul>
                <li>
                  <red href="#">25/600</red>
                </li>
              </ul>
            </li>
          </ul>
        </li>
        <li>
          <black href="#">50/600</black>
        </li>
      </ul>
    </li>
  </ul>
</div>
`,
		},
	}

	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := llrb.New()
			for _, kv := range tt.inserted {
				tree.Insert(kv.k, kv.v)
			}

			w := &bytes.Buffer{}
			tree.ToHTML(w)
			got := w.String()
			if tt.want != got {
				t.Errorf("\nwant\n%v\ngot\n%v", tt.want, got)
			}
		})
	}
}

func assertTree(t *testing.T, tree *llrb.Tree, kvs []kv) {
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

func Benchmark_Ascending_1000(b *testing.B) {
	tree := llrb.New()
	ascending(b, tree, 1000)
}

func Benchmark_Descending_1000(b *testing.B) {
	tree := llrb.New()
	descending(b, tree, 1000)
}

func Benchmark_Ascending_10000(b *testing.B) {
	tree := llrb.New()
	ascending(b, tree, 10000)
}

func Benchmark_Descending_10000(b *testing.B) {
	tree := llrb.New()
	descending(b, tree, 10000)
}

func Benchmark_Ascending_100000(b *testing.B) {
	tree := llrb.New()
	ascending(b, tree, 100000)
}

func Benchmark_Descending_100000(b *testing.B) {
	tree := llrb.New()
	descending(b, tree, 100000)
}

func Benchmark_Ascending_200000(b *testing.B) {
	tree := llrb.New()
	ascending(b, tree, 200000)
}

func Benchmark_Descending_200000(b *testing.B) {
	tree := llrb.New()
	descending(b, tree, 200000)
}

func Benchmark_Ascending_400000(b *testing.B) {
	tree := llrb.New()
	ascending(b, tree, 400000)
}
func Benchmark_Descending_400000(b *testing.B) {
	tree := llrb.New()
	descending(b, tree, 400000)
}

func ascending(b *testing.B, tree *llrb.Tree, n int) {
	for i := n; i > 0; i-- {
		tree.Insert(i, i)
	}

	assertNumOfTree(b, tree, n)

	for i := n; i > 0; i-- {
		_, _ = tree.Search(i)
	}

	for i := n; i > 0; i-- {
		tree.Delete(i)
	}

	b.StopTimer()

	assertNumOfTree(b, tree, 0)
}

func descending(b *testing.B, tree *llrb.Tree, n int) {
	for i := n; i > 0; i-- {
		tree.Insert(i, i)
	}

	assertNumOfTree(b, tree, n)

	for i := n; i > 0; i-- {
		_, _ = tree.Search(i)
	}

	for i := n; i > 0; i-- {
		tree.Delete(i)
	}

	b.StopTimer()

	assertNumOfTree(b, tree, 0)
}

func assertNumOfTree(b *testing.B, tree *llrb.Tree, want int) {
	if want != tree.Count() {
		b.Errorf("num of nodes must be %v, got %v", want, tree.Count())
	}
}
