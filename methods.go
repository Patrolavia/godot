package godot

import (
	"fmt"
	"strconv"
	"strings"
)

type graph struct {
	strict bool
	tag    string
	name   string
	opts   map[string]string
	nodes  map[string]map[string]string
	edges  []edge
	groups []Subgraph
}

func NewGraph(strict, directed bool, name string, opts map[string]string) Graph {
	tag := "graph"
	if directed {
		tag = "digraph"
	}

	return &graph{
		strict,
		tag,
		name,
		opts,
		make(map[string]map[string]string),
		make([]edge, 0),
		make([]Subgraph, 0),
	}
}

func (g *graph) Strict() bool {
	return g.strict
}

func (g *graph) Direct() bool {
	return g.tag == "digraph"
}

func (g *graph) AddNode(name string, opt map[string]string) Node {
	g.nodes[name] = opt
	return node(name)
}

func (g *graph) AddEdge(nodes []Node, opt map[string]string) {
	g.edges = append(g.edges, edge{nodes, opt})
}

func (g *graph) AutoEdge(names []string, opt map[string]string) []Node {
	nodes := make([]Node, len(names))
	for k, v := range names {
		nodes[k] = node(v)
	}
	g.AddEdge(nodes, opt)
	return nodes
}

func (g *graph) Group(name string, sub Subgraph, opts map[string]string) {
	sub.setOptions(opts)
	sub.setName(name)
	g.groups = append(g.groups, sub)
}

func (g *graph) Cluster(name string, sub Subgraph, opts map[string]string) {
	g.groups = append(g.groups, sub)
	sub.setOptions(opts)
	if name == "" {
		name = strconv.Itoa(len(g.groups))
	}
	sub.setName(fmt.Sprintf("cluster_%s", name))
}

func (g *graph) String() string {
	ret := make([]string, 1, len(g.nodes)+len(g.edges)+len(g.groups)+3)
	strict := ""
	if g.strict {
		strict = "strict "
	}
	ret[0] = fmt.Sprintf("%s%s %s {", strict, g.tag, escape(g.name))
	if opts := Options(g.opts).AsBlock("  "); opts != "" {
		ret = append(ret, opts)
	}
	nodes_processed := make(map[string]bool)
	mark := func(n Node) {
		nodes_processed[n.Name()] = true
	}
	need := func(n Node, opt map[string]string) bool {
		name := n.Name()
		return name != "" && (opt != nil || !nodes_processed[name])
	}
	edgeop := "--"
	if g.Direct() {
		edgeop = "->"
	}

	// process edges and record used nodes
	for _, edge := range g.edges {
		ret = append(ret, "  " + edge.export(edgeop))
		for _, n := range edge.nodes {
			mark(n)
		}
	}

	// process nodes, skip all nodes which introduced in edges but without custom defination
	for n, o := range g.nodes {
		n := node(n)
		if need(n, o) {
			ret = append(ret, fmt.Sprintf("  %s %s;", n.AsNode(), Options(o).AsLine()))
		}
	}

	// process groups
	for _, g := range g.groups {
		ret = append(ret, g.AsGroup())
	}

	return strings.Join(ret, "\n") + "\n}"
}
