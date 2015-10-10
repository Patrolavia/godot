// Package godot is a simeple tool helps your generate graphviz's dot file.
//
// Caution
//
// * godot DOES NO VALIDATATION, use with cares.
// * Subgraph in graphviz might not work in the way you expected, read graphviz's manual before using it.
package godot

import "strings"

func escape(str string) string {
	return `"` + strings.Replace(str, `"`, `\"`, -1) + `"`
}

// Options represents options of graph, node or edge.
type Options map[string]string

// Node represented a node in dot file
type Node interface {
	Name() string   // node name, empty string if anonymous node like subgraph
	AsNode() string // node literal, escaped node name or {n0 n1} if subgraph
}

type node string

func (n node) Name() string {
	return string(n)
}

func (n node) AsNode() string {
	return escape(string(n))
}

// Subgraph represents subgraph in dot file.
type Subgraph interface {
	Node
	AsGroup() string
	setOptions(opts map[string]string)
	setName(name string)
}

// Graph denotes a graph/digraph tag in dot file
type Graph interface {
	Strict() bool
	Direct() bool
	AddNode(name string, opts map[string]string) Node // overwrite old node info when already exists
	AddEdge(nodes []Node, opts map[string]string)
	AutoEdge(names []string, opts map[string]string) []Node // create node before add edge
	Group(name string, sub Subgraph, opts map[string]string)
	Cluster(name string, sub Subgraph, opts map[string]string)
	String() string
}
