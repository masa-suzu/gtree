package llrb

import (
	"fmt"
	"io"
)

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

func (n *node) ToHTML(w io.Writer) {
	indent(w, []byte("<ul>\n"), 1)
	n.toHTML(w, 1)
	indent(w, []byte("</ul>\n"), 1)
}

func (n *node) toHTML(w io.Writer, numOfIndents int) {

	indent(w, []byte("<li>\n"), numOfIndents+1)

	if n.color {
		indent(w, []byte(fmt.Sprintf("<red href=\"#\">%v/%v</red>\n", n.key, n.value)), numOfIndents+2)
	} else {
		indent(w, []byte(fmt.Sprintf("<black href=\"#\">%v/%v</black>\n", n.key, n.value)), numOfIndents+2)
	}

	if n.left != nil || n.right != nil {
		indent(w, []byte("<ul>\n"), numOfIndents+2)
	}

	if n.left != nil {
		n.left.toHTML(w, numOfIndents+2)
	}
	if n.right != nil {
		n.right.toHTML(w, numOfIndents+2)
	}

	if n.left != nil || n.right != nil {
		indent(w, []byte("</ul>\n"), numOfIndents+2)
	}

	indent(w, []byte("</li>\n"), numOfIndents+1)
}

func indent(w io.Writer, v []byte, indents int) {
	for index := 0; index < indents; index++ {
		w.Write([]byte("  "))
	}
	w.Write(v)
}
