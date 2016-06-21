// Copyright 2015 Husobee Associates, LLC.  All rights reserved.
// Portions Copyright 2015 Labstack.  All rights reserved.
// Use of this source code is governed by The MIT License, which
// can be found in the LICENSE file included.

package vestigo

import "fmt"

// node - a node structure for nodes within the tree
type node struct {
	typ       ntype
	label     byte
	prefix    string
	parent    *node
	children  children
	resource  *resource
	pnames    []string
	fmtpnames []string
}

// prefix - print the prefix
func prefix(tail bool, p, on, off string) string {
	if tail {
		return fmt.Sprintf("%s%s", p, on)
	}
	return fmt.Sprintf("%s%s", p, off)
}

// printTree - Helper method to print a representation of the tree
func (n *node) PrintTree(pfx string, tail bool) {
	p := prefix(tail, pfx, "└── ", "├── ")
	fmt.Printf("%s%s, %p: type=%d, parent=%p, resource=%v\n", p, n.prefix, n, n.typ, n.parent, n.resource)

	children := n.children
	l := len(children)
	p = prefix(tail, pfx, "    ", "│   ")
	for i := 0; i < l-1; i++ {
		children[i].PrintTree(p, false)
	}
	if l > 0 {
		children[l-1].PrintTree(p, true)
	}
}

// newNode - create a new router tree node
func newNode(t ntype, pre string, p *node, c children, h *resource, pnames []string) *node {
	n := &node{
		typ:      t,
		label:    pre[0],
		prefix:   pre,
		parent:   p,
		children: c,
		// create a resource method to handler map for this node
		resource: h,
		pnames:   pnames,
	}
	for _, v := range pnames {
		n.fmtpnames = append(n.fmtpnames, "%3A"+v+"=")
	}
	return n
}

// addChild - Add a child node to this node
func (n *node) addChild(c *node) {
	n.children = append(n.children, c)
}

// findChild - find a child node of this node
func (n *node) findChild(l byte, t ntype) *node {
	for _, c := range n.children {
		if c.label == l && c.typ == t {
			return c
		}
	}
	return nil
}

// findChildWithLabel - find a child with a matching label, label being the first byte in the prefix
func (n *node) findChildWithLabel(l byte) *node {
	for _, c := range n.children {
		if c.label == l {
			return c
		}
	}
	return nil
}

// findChildWithType - find a child with a matching type
func (n *node) findChildWithType(t ntype) *node {
	for _, c := range n.children {
		if c.typ == t {
			return c
		}
	}
	return nil
}
