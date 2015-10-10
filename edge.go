package godot

import "strings"

type edge struct {
	nodes []Node
	opts  map[string]string
}

func (e edge) export(op string) string {
	nodes := make([]string, len(e.nodes))
	for k, v := range e.nodes {
		nodes[k] = v.AsNode()
	}
	return strings.Join(nodes, " "+op+" ") + " " + Options(e.opts).AsLine() + ";"
}
