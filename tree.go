/*
 Package llrb provides an implementation of Left-Leaning Red-Black Tree.
 Original implementation is available from http://www.cs.princeton.edu/~rs/talks/LLRB/LLRB.pdf.
*/
package llrb

import (
	"fmt"
)

const (
	lt = -1
	eq = 0
	gt = 1
)

type Tree struct {
	root  *node
	count int
}

func New() *Tree {
	return &Tree{
		root:  nil,
		count: 0,
	}
}

func (t *Tree) Count() int {
	return t.count
}

// Search returns a value with a given key.
// If no value with the key is found, returns error.
func (t *Tree) Search(key int) (interface{}, error) {
	x := t.root

	for x != nil {
		cmp := compare(key, x.key)
		switch cmp {
		case eq:
			return x.value, nil
		case lt:
			x = x.left
		case gt:
			x = x.right
		}
	}
	return nil, fmt.Errorf("not found a item with the given k. Key: %v", key)
}

// Insert a value with a given key.
// If the key has already inserted, the new value overrides old one.
func (t *Tree) Insert(key int, value interface{}) {
	t.root = t.insert(t.root, key, value)

	if t.root == nil {
		return
	}

	t.root.color = black
}

func (t *Tree) insert(n *node, key int, value interface{}) *node {
	if n == nil {
		t.count = t.count + 1
		return &node{
			key:   key,
			value: value,
			color: red,
			left:  nil,
			right: nil,
		}
	}

	cmp := compare(key, n.key)

	switch cmp {
	case eq:
		n.value = value
	case lt:
		n.left = t.insert(n.left, key, value)
	case gt:
		n.right = t.insert(n.right, key, value)
	}
	return fixup(n)
}

// Delete remove a node by a given key.
// If no value with the key found, do nothing.
func (t *Tree) Delete(key int) {
	t.root = t.delete(t.root, key)
	if t.root == nil {
		return
	}
	t.root.color = black
}

func (t *Tree) delete(n *node, key int) *node {
	if n == nil {
		return nil
	}

	if compare(key, n.key) == lt {
		if n.left.isBlack() && !n.left.left.isRed() {
			n = moveRedLeft(n)
		}
		n.left = t.delete(n.left, key)
	} else {
		if n.left.isRed() {
			n = rotateRight(n)
		}
		if n.right.isBlack() && !n.right.left.isRed() {
			n = moveRedRight(n)
		}

		if compare(key, n.key) == eq {
			t.count = t.count - 1

			if n.right == nil {
				return nil
			}

			rm := min(n.right)
			n.key = rm.key
			n.value = rm.value
			n.right = deleteMin(n.right)

		} else {
			n.right = t.delete(n.right, key)
		}
	}
	return fixup(n)
}

func deleteMin(n *node) *node {
	if n.left == nil {
		return nil
	}

	if n.left.isBlack() && !n.left.left.isRed() {
		n = moveRedLeft(n)
	}
	n.left = deleteMin(n.left)
	return fixup(n)
}

func fixup(n *node) *node {
	if n.right.isRed() {
		n = rotateLeft(n)
	}
	if n.left.isRed() && n.left.left.isRed() {
		n = rotateRight(n)
	}
	if n.left.isRed() && n.right.isRed() {
		flip(n)
	}
	return n
}

func flip(n *node) {
	n.color = !n.color
	n.left.color = !n.left.color
	n.right.color = !n.right.color
}

func rotateLeft(n *node) *node {
	var x = n.right
	n.right = x.left
	x.left = n
	x.color = n.color
	n.color = red
	return x
}

func rotateRight(n *node) *node {
	var x = n.left
	n.left = x.right
	x.right = n
	x.color = n.color
	n.color = red
	return x
}

func moveRedLeft(n *node) *node {
	flip(n)
	if n.right.left.isRed() {
		n.right = rotateRight(n.right)
		n = rotateLeft(n)
		flip(n)
	}
	return n
}

func moveRedRight(n *node) *node {
	flip(n)
	if n.left.left.isRed() {
		n = rotateRight(n)
		flip(n)
	}
	return n
}

func min(n *node) *node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func compare(k1, k2 int) int {
	if k1 < k2 {
		return lt
	}
	if k1 > k2 {
		return gt
	}
	return eq
}
