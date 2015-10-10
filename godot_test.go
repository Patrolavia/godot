package godot

import (
	"fmt"
	"testing"
)

func TestFunctional(t *testing.T) {
	expected := `digraph "my diagram" {
  "n1" -> "n2" ;
  "n2" -> {"n9" "n10"} ;
  "n2" -> {"n5" "n6"} -> {"n7" "n8"} ;
  "n10" -> "n11" -> "n12" ;
  "n1" [color="red"];
  subgraph  {
    color="blue";
    "n3";
    "n4";
  }
  subgraph cluster_2 {
    color="blue";
    "n5";
    "n6";
  }
  subgraph cluster_name {
    rank="max";
    color="red";
    "n7";
    "n8";
  }
}`

	g := NewGraph(false, true, "my diagram", nil)
	n1 := g.AddNode("n1", map[string]string{"color": "red"})
	n2 := g.AddNode("n2", nil)
	g.AddEdge([]Node{n1, n2}, nil)
	sub := NewSubgraph("n3", "n4")
	g.Group("", sub, map[string]string{"color": "blue"})
	sub2 := NewSubgraph("n5", "n6")
	g.Cluster("", sub2, map[string]string{"color": "blue"})
	sub3 := NewSubgraph("n7", "n8")
	g.Cluster("name", sub3, map[string]string{"rank": "max", "color": "red"})

	grp := NewSubgraph("n9", "n10")
	g.AddEdge([]Node{n2, grp}, nil)
	g.AddEdge([]Node{n2, sub2, sub3}, nil)
	g.AutoEdge([]string{"n10", "n11", "n12"}, nil)

	if o := g.String(); o != expected {
		t.Errorf(`Output of godot does not match what we expected: %s`, o)
	}
}

func Example() {
	g := NewGraph(false, true, "my diagram", nil)
	n0 := g.AddNode("n0", map[string]string{"color": "red"})
	n1 := g.AddNode("n1", nil)
	g.AddEdge([]Node{n0, n1}, nil)

	sub := NewSubgraph("n2", "n3", "n4")
	g.AddEdge([]Node{n1, sub}, map[string]string{"label": "a subgraph"})
	g.Cluster("", sub, nil)

	g.AutoEdge([]string{"n3", "n4", "n5"}, nil)

	fmt.Println(g.String())
	// output: digraph "my diagram" {
	//   "n0" -> "n1" ;
	//   "n1" -> {"n2" "n3" "n4"} [label="a subgraph"];
	//   "n3" -> "n4" -> "n5" ;
	//   "n0" [color="red"];
	//   subgraph cluster_1 {
	//     "n2";
	//     "n3";
	//     "n4";
	//   }
	// }
}
