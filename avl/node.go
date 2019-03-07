package avl

type node struct {
	key   int
	value interface{}
	left  *node
	right *node
}
