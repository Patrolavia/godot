# godot - generating graphviz diagram with ease

godot helps you generating graphviz diagrams in your application.

[![GoDoc](https://godoc.org/github.com/Patrolavia/godot?status.svg)](https://godoc.org/github.com/Patrolavia/godot) [![Build Status](https://travis-ci.org/Patrolavia/godot.svg)](https://travis-ci.org/Patrolavia/godot)

## Synopsis

https://godoc.org/github.com/Patrolavia/godot/#example_

```go
g := NewGraph(false, true, "my diagram", nil)
n0 := g.AddNode("n0", map[string]string{"color": "red"})
n1 := g.AddNode("n1", nil)
g.AddEdge([]Node{n0, n1}, nil)

sub := NewSubgraph("n2", "n3", "n4")
g.AddEdge([]Node{n1, sub}, map[string]string{"label": "a subgraph"})
g.Cluster("", sub, nil)

g.AutoEdge([]string{"n3", "n4", "n5"}, nil)

fmt.Println(g.String())
```

Outputs: [svg file rendered using dot](https://raw.githubusercontent.com/Patrolavia/godot/master/example.svg)

```
digraph "my diagram" {
  "n0" -> "n1" ;
  "n1" -> {"n2" "n3" "n4"} [label="a subgraph"];
  "n3" -> "n4" -> "n5" ;
  "n0" [color="red"];
  subgraph cluster_1 {
    "n2";
    "n3";
    "n4";
  }
}
```

## License

Any version of MIT, GPL or LGPL.
