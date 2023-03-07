package web

import (
	"net/http"
	"strings"
)

type HandlerBasedOnTree struct {
	root *Node
}

func (h *HandlerBasedOnTree) ServerHTTP(c *Context) {
	handler, found := h.findRouter(c.R.URL.Path)
	if !found {
		c.W.WriteHeader(http.StatusNotFound)
		_, _ = c.W.Write([]byte("Not Found"))
		return
	}
	handler(c)
}

func (h *HandlerBasedOnTree) Route(method string, pattern string,
	handlerFunc func(ctx *Context)) {
	paths := strings.Split(strings.Trim(pattern, "/"), "/")

	cur := h.root
	for index, path := range paths {
		mathChild, ok := cur.findMatchChild(path)
		if ok {
			cur = mathChild
		} else {
			h.createSubTree(cur, paths[index:], handlerFunc)
			return
		}
	}
}
func (h *HandlerBasedOnTree) createSubTree(root *Node,
	paths []string, handlerFunc handlerFunc) {
	cur := root
	for _, path := range paths {
		nn := newNode(path)
		cur.children = append(cur.children, nn)
		cur = nn
	}
	cur.handler = handlerFunc
}

func (h *HandlerBasedOnTree) findRouter(path string) (handlerFunc, bool) {
	paths := strings.Split(strings.Trim(path, "/"), "/")
	cur := h.root
	for _, p := range paths {
		matchChild, found := h.findMatchChild(cur, p)
		if !found {
			return nil, false
		}
		cur = matchChild
	}
	if cur.handler == nil {
		return nil, false
	}
	return cur.handler, true
}
func (h *HandlerBasedOnTree) findMatchChild(root *Node,
	path string) (*Node, bool) {
	var wildcardNode *Node
	for _, child := range root.children {
		if child.path == path && child.path != "*" {
			return child, true
		}
		if child.path == "*" {
			wildcardNode = child
		}
	}
	return wildcardNode, wildcardNode != nil
}

type Node struct {
	path string

	children  []*Node
	handler   handlerFunc
	matchFunc handlerFunc
}

func (h *Node) findMatchChild(path string) (*Node, bool) {
	for _, child := range h.children {
		if child.path == path {
			return child, true
		}
	}
	return nil, false
}

func newNode(path string) *Node {
	return &Node{
		path:     path,
		children: make([]*Node, 0, 8),
	}
}
