package godot

import "strings"

type group struct {
	nodes []Node
	opts  map[string]string
	name  string
}

// NewGroup creates a node group
func NewSubgraph(nodes ...string) Subgraph {
	ret := make([]Node, len(nodes))
	for k, v := range nodes {
		ret[k] = node(v)
	}
	return &group{nodes: ret}
}

func (g *group) setOptions(opts map[string]string) {
	g.opts = opts
}

func (g *group) setName(name string) {
	g.name = name
}

func (g *group) Name() string {
	return ""
}

func (g *group) Nodes() []Node {
	return g.nodes
}

func (g *group) AsNode() string {
	nodes := make([]string, len(g.nodes))
	for k, v := range g.nodes {
		nodes[k] = v.AsNode()
	}
	return `{` + strings.Join(nodes, " ") + `}`
}

func (g *group) AsGroup() string {
	ret := make([]string, 1, len(g.nodes)+3)
	ret[0] = "  subgraph " + g.name + " {"
	indent := "    "
	if opt := Options(g.opts).AsBlock(indent); opt != "" {
		ret = append(ret, opt)
	}
	for _, node := range g.nodes {
		ret = append(ret, indent+node.AsNode()+";")
	}
	ret = append(ret, "  }")
	return strings.Join(ret, "\n")
}
