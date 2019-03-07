package avl

type node struct {
	height int
	key    int
	value  interface{}
	left   *node
	right  *node
}
