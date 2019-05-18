package gizmo

import (
	"fmt"
	"strconv"
)

const (
	// GRYO format
	GRYO = "gryo()"
	// GRAPHSON format
	GRAPHSON = "graphson()"
	// GRAPHML format
	GRAPHML = "graphml()"
)

// GraphAPI struct
type GraphAPI struct {
	value string
}

// Graph function
func Graph() *GraphAPI {
	return &GraphAPI{"graph."}
}

// String function
func (g *GraphAPI) String() string {
	if len(g.value) > 0 {
		if g.value[len(g.value)-1:] == "." { // last character is a .
			g.value = g.value[:len(g.value)-1] // removes the last character
		}
	}
	return g.value
}

// Traversal function
func (g *GraphAPI) Traversal(name ...string) *Traversal {
	tmp := graphTraversal(name...)
	return tmp
}

// Property step is a sideEffect. When a vertex is added from a Graph g using addV, properties can be added with property().
// For a previously created vertex, properties can be added. A previously created vertex label must be specified.
// Property key-value pairs are specified.
func (g *GraphAPI) Property(k interface{}, v interface{}) *GraphAPI {

	if _, ok := k.(string); ok {
		k = "'" + k.(string) + "'"
	}

	if _, ok := v.(string); ok {
		v = "'" + v.(string) + "'"
	}

	g.value = g.value + fmt.Sprintf("property(%v, %v).", k, v)
	return g
}

// IO Graph data is written to a file or read from a file using io.
// The file to read must be located on a DSE cluster node, and the written file will be created on the
// DSE cluster node on which the command is run.
func (g *GraphAPI) IO(format string) *GraphAPI {
	g.value = g.value + "io(" + format + ")."
	return g
}

// WriteGraph method essentially calls writeVertices with the directionality of BOTH.
// The writeVertices method simply calls writeVertex which detaches each Vertex instance into a directional
// StarGraph which forms the basis for the format.
func (g *GraphAPI) WriteGraph(name string) *GraphAPI {
	g.value = g.value + "writeGraph('" + name + "')."
	return g
}

// ReadGraph method essentially calls writeVertices with the directionality of BOTH.
// The writeVertices method simply calls writeVertex which detaches each Vertex instance into a directional
// StarGraph which forms the basis for the format.
func (g *GraphAPI) ReadGraph(name string) *GraphAPI {
	g.value = g.value + "readGraph('" + name + "')."
	return g
}

// Config function
func (g *GraphAPI) Config() *GraphAPI {
	g.value = g.value + "config()."
	return g
}

// Option function
func (g *GraphAPI) Option(o string, v bool) *GraphAPI {
	g.value = g.value + fmt.Sprintf("optoin('%s', %v).", o, strconv.FormatBool(v))
	return g
}

// Open function
func (g *GraphAPI) Open() *GraphAPI {
	g.value = g.value + "open()."
	return g
}

// TX function
func (g *GraphAPI) TX() *GraphAPI {
	g.value = g.value + "tx()."
	return g
}

// Raw function is a helper function that can used in case a step is not supported
// or not available in this package. the string is appended as is without modification
func (g *GraphAPI) Raw(s string) *GraphAPI {
	g.value = g.value + s + "."
	return g
}
