/*
	Package avl provides an implementation of AVL Tree.
*/
package avl

// Tree implements an AVL tree.
type Tree struct {
	root  *node
	count int
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
