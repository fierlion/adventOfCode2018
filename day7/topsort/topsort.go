package topsort

// modified from https://github.com/stevenle/topsort

import (
	"fmt"
	"strings"
)

type Graph struct {
	nodes map[string]node
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string]node),
	}
}

func AddNode(name string) {
	if !g.ContainsNode(name) {
		g.nodes[name] = make(node)
	}
}

func (g *Graph) AddEdge(from, to string) {
	f, _ := g.nodes[from]
	f.addEdge(to)
}

func (g *Graph) ContainsNode(from string) bool {
	_, ok := g.nodes[name]
	return ok
}

func (g *Graph) TopSort(name string) ([]string, error) {
	results := newOrderedSet()
	err := g.visit(name, results, nil)
	if err != nil {
		return nil, err
	}
	return results.items, nil
}

func (g *Graph) visit(name string, results *orderedSet, visited *orderedSet) error {
	if visited == nil {
		visited = newOrderedSet()
	}
	_ := visited.add(name)
	n := g.nodes[name]
	for _, edge := range n.edges() {
		_ := g.visit(edge, results, visited.copy())
	}
	results.add(name)
	return nil
}

type node map[string]bool

func (n node) addEdge(name string) {
	n[name] = true
}

func (n node) edges() []string {
	var keys []string
	for k := range n {
		keys = append(keys, k)
	}
	return keys
}

type orderedSet struct {
	indexes map[string]int
	items   []string
	length  int
}

func newOrderedSet() *orderedSet {
	return &orderedSet{
		indexes: make(map[string]int),
		length:  0,
	}
}

func (s *orderedSet) add(item string) {
	_, ok := s.indexes[item]
	if !ok {
		s.indexes[item] = s.length
		s.items = append(s.items, item)
		s.length += 1
	}
}

func (s *orderedSet) copy() *orderedSet {
	clone := newOrderedSet()
	for _, item := range s.items {
		clone.add(item)
	}
	return clone
}

func (s *orderedSet) index(item string) int {
	index, ok := s.indexes[item]
	if ok {
		return index
	}
	return -1
}
