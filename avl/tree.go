/*
	Package avl provides an implementation of AVL Tree.
*/
package avl

import (
	"fmt"
	"math"
)

const (
	lt = -1
	eq = 0
	gt = 1
)

type kvp struct {
	key   int
	value interface{}
}

// Tree implements an AVL tree.
type Tree struct {
	root       *node
	count      int
	needUpdate bool
	max        *kvp
}

// New returns a reference to an empty Tree.
func New() *Tree {
	return &Tree{
		root:  nil,
		count: 0,
	}
}

// Count returns num of nodes.
func (t *Tree) Count() int {
	return t.count
}

// Insert a value with a given key.
// If the same key has already inserted, the new value overrides old one.
func (t *Tree) Insert(key int, value interface{}) {
	t.root = t.insert(t.root, key, value)
}

func (t *Tree) insert(n *node, key int, value interface{}) *node {
	if n == nil {
		t.needUpdate = true
		t.count++
		return &node{
			height: 1,
			key:    key,
			value:  value,
		}
	}

	cmp := compare(key, n.key)
	switch cmp {
	case lt:
		n.left = t.insert(n.left, key, value)
		return t.balanceLeft(n)
	case eq:
		n.value = value
		return n
	case gt:
		t.needUpdate = false
		n.right = t.insert(n.right, key, value)
		return t.balanceRight(n)
	}

	return nil
}

func (t *Tree) balanceLeft(n *node) *node {
	if !t.needUpdate {
		return n
	}

	h := height(n)
	if bias(n) == 2 {
		if bias(n.left) >= 0 {
			n = rotateRight(n)
		} else {
			n = rotateLeftRight(n)
		}
	} else {
		modifyHeight(n)
	}
	t.needUpdate = h != height(n)
	return n
}

func (t *Tree) balanceRight(n *node) *node {
	if !t.needUpdate {
		return n
	}

	h := height(n)
	if bias(n) == 2 {
		if bias(n.right) <= 0 {
			n = rotateLeft(n)
		} else {
			n = rotateRightLeft(n)
		}
	} else {
		modifyHeight(n)
	}
	t.needUpdate = h != height(n)
	return n
}

func rotateLeftRight(n *node) *node {
	n.left = rotateLeft(n.left)
	return rotateRight(n)
}
func rotateRightLeft(n *node) *node {
	n.right = rotateLeft(n.right)
	return rotateLeft(n)
}

func height(n *node) int {
	if n == nil {
		return 0
	}

	return n.height
}

func bias(n *node) int {
	return height(n.left) - height(n.right)
}

func modifyHeight(n *node) {
	n.height = 1 + int(math.Max(float64(height(n.left)), float64(height(n.right))))
}

func rotateLeft(v *node) *node {
	u := v.right
	n := u.left
	u.left = v
	v.right = n
	modifyHeight(u.left)
	modifyHeight(u)
	return u
}

func rotateRight(u *node) *node {
	v := u.left
	n := v.right
	v.right = u
	u.left = n
	modifyHeight(v.right)
	modifyHeight(v)
	return v
}

// Search returns a value associated with a given key.
// If no value is found by the key, returns nil with an error.
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
	return nil, fmt.Errorf("found no value by key '%v'", key)
}

// Delete remove a node by a given key.
// If the key does not found, do nothing.
func (t *Tree) Delete(key int) {
	t.root = t.delete(t.root, key)
}

func (t *Tree) delete(n *node, key int) *node {
	if n == nil {
		t.needUpdate = false
		return nil
	}

	cmp := compare(key, n.key)
	switch cmp {
	case lt:
		n.left = t.delete(n.left, key)
		return t.balanceRight(n)
	case gt:
		n.right = t.delete(n.right, key)
		return t.balanceLeft(n)
	case eq:
		t.count--
		if n.left == nil {
			t.needUpdate = true
			return n.right
		} else {
			n.left = t.deleteMax(n.left)
			n.key = t.max.key
			n.value = t.max.value
			return t.balanceRight(n)
		}
	}
	panic("unknown switch case")

}

func (t *Tree) deleteMax(n *node) *node {
	if n.right != nil {
		n.right = t.deleteMax(n.right)
		return t.balanceLeft(n)
	}

	t.needUpdate = true
	t.max = &kvp{
		key:   n.key,
		value: n.value,
	}
	return n.left
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
