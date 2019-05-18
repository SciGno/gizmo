# gizmo
Gremlin language builder for Go.

This is a work in progress for now.  At the moment, this package can only be used for generating Gemlin query strings.
#

Example:

```go
package main

import (
	"fmt"
	"os"

	"github.com/scigno/gizmo"
)

func main() {

	graph := gizmo.Graph()
	g := graph.Traversal("g")
	g1 := g.New("g")
	g1.V().Has("user", "username", "scigno")

	n := g.New().Raw("t=").Append(g1.String())
	n.AddLine(g.New().TernaryOp(
		g.New("t").HasNext(),
		g.New("t").Next(),
		g.New("g").AddV("user").Property("userId", "744be509-a1cc-466d-bb10-0bb9a376da2e").Property("username", "scigno").Next(),
	).String())
	fmt.Println(n)
}
```

Output:
```bash
t=g.V().has('user','username','scigno') 
t.hasNext() ? t.next() : g.addV('user').property('userId', '744be509-a1cc-466d-bb10-0bb9a376da2e').property('username', 'scigno').next()
```