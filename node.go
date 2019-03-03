package llrb

const (
	red   = true
	black = false
)

type node struct {
	key   int
	value interface{}
	left  *node
	right *node
	color bool
}

func (n *node) isRed() bool {
	if n == nil {
		return false
	}
	return n.color
}

func (n *node) isBlack() bool {
	if n == nil {
		return false
	}
	return !n.color
}
