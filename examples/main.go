package main

import (
	"fmt"

	"github.com/scigno/gizmo"
)

func main() {

	graph := gizmo.Graph()
	g := graph.Traversal("g")
	g1 := g.New("g")
	g1.V().Has("user", "username", "scigno")

	n := g.New().Raw("t=").Append(g1)
	n.AddLine(g.New().TernaryOp(
		g.New("t").HasNext(),
		g.New("t").Next(),
		g.New("g").AddV("user").Property("userId", "744be509-a1cc-466d-bb10-0bb9a376da2e").Property("username", "scigno").Next(),
	))
	fmt.Println(n)
}
