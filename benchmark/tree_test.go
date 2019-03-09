package benchmark

import (
	"testing"

	"github.com/masa-suzu/gtree/avl"
	"github.com/masa-suzu/gtree/llrb"
)

type kvs interface {
	Search(key int) (interface{}, error)
	Insert(key int, value interface{})
	Delete(key int)
	Count() int
}

func Benchmark_Ascending_10000_avl(b *testing.B) {
	tree := avl.New()
	ascending(b, tree, 10000)
}

func Benchmark_Descending_10000_avl(b *testing.B) {
	tree := avl.New()
	descending(b, tree, 10000)
}

func Benchmark_Ascending_100000_avl(b *testing.B) {
	tree := avl.New()
	ascending(b, tree, 100000)
}

func Benchmark_Descending_100000_avl(b *testing.B) {
	tree := avl.New()
	descending(b, tree, 100000)
}

func Benchmark_Ascending_200000_avl(b *testing.B) {
	tree := avl.New()
	ascending(b, tree, 200000)
}

func Benchmark_Descending_200000_avl(b *testing.B) {
	tree := avl.New()
	descending(b, tree, 200000)
}

func Benchmark_Ascending_400000_avl(b *testing.B) {
	tree := avl.New()
	ascending(b, tree, 400000)
}
func Benchmark_Descending_400000_avl(b *testing.B) {
	tree := avl.New()
	descending(b, tree, 400000)
}

func Benchmark_Ascending_1000_llrb(b *testing.B) {
	tree := llrb.New()
	ascending(b, tree, 1000)
}

func Benchmark_Descending_1000_llrb(b *testing.B) {
	tree := llrb.New()
	descending(b, tree, 1000)
}

func Benchmark_Ascending_10000_llrb(b *testing.B) {
	tree := llrb.New()
	ascending(b, tree, 10000)
}

func Benchmark_Descending_10000_llrb(b *testing.B) {
	tree := llrb.New()
	descending(b, tree, 10000)
}

func Benchmark_Ascending_100000_llrb(b *testing.B) {
	tree := llrb.New()
	ascending(b, tree, 100000)
}

func Benchmark_Descending_100000_llrb(b *testing.B) {
	tree := llrb.New()
	descending(b, tree, 100000)
}

func Benchmark_Ascending_200000_llrb(b *testing.B) {
	tree := llrb.New()
	ascending(b, tree, 200000)
}

func Benchmark_Descending_200000_llrb(b *testing.B) {
	tree := llrb.New()
	descending(b, tree, 200000)
}

func Benchmark_Ascending_400000_llrb(b *testing.B) {
	tree := llrb.New()
	ascending(b, tree, 400000)
}

func ascending(b *testing.B, tree kvs, n int) {
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

func descending(b *testing.B, tree kvs, n int) {
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

func assertNumOfTree(b *testing.B, tree kvs, want int) {
	if want != tree.Count() {
		b.Errorf("num of nodes must be %v, got %v", want, tree.Count())
	}
}
