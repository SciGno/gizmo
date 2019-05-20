package gizmo

// Traversal struct
type Traversal struct {
	value string
}

// GraphTraversal function
func graphTraversal(name ...string) *Traversal {
	if len(name) == 0 {
		return &Traversal{"g."}
	} else if name[0] == "" {
		return &Traversal{"g."}
	}
	return &Traversal{name[0] + "."}
}

// New function returns a new traversal
func (t *Traversal) New(name ...string) *Traversal {
	if len(name) == 0 {
		return &Traversal{""}
	}
	return &Traversal{name[0] + "."}
}

// String function
func (t *Traversal) String() string {
	if t.value[len(t.value)-1:] == "." {
		t.value = t.value[:len(t.value)-1]
	}
	return t.value
}

// V function
func (t *Traversal) V(name ...string) *Traversal {
	return processStringSlice("V", t, name...)
}

// AddV step is a map/sideEffect. A vertex is added from a Traversal g using addV.
// A previously created vertex label must be specified. Property key-value pairs may be optionally specified.
func (t *Traversal) AddV(name ...string) *Traversal {
	return processStringSlice("addV", t, name...)
}

// Where returns a new gremlin Traversal
func (t *Traversal) Where(i interface{}) *Traversal {
	return processInterface("where", t, i)
}

// And step can take an arbitrary number of Traversals to filter the returned objects using a boolean AND function.
// The related boolean OR function can be done with the or() step.
func (t *Traversal) And(i interface{}) *Traversal {
	return processInterface("and", t, i)
}

// FlatMap step is a general step. This step creates a map of the Traversal objects and streams the map to the next Traversal step.
func (t *Traversal) FlatMap(i interface{}) *Traversal {
	return processInterface("flatMap", t, i)
}

// AddE is a map/sideEffect. An edge is added from a Traversal g using addE between two existing vertices.
// A previously created edge label must be specified.
func (t *Traversal) AddE(s string) *Traversal {
	return processString("addE", t, s)
}

// Property step in a traversal has the same beahviour as in the Graph API
func (t *Traversal) Property(k interface{}, v interface{}) *Traversal {
	return processKeyVakues("property", k, v, t)
}

// Limit returns a new gremlin Traversal
func (t *Traversal) Limit(i int) *Traversal {
	return processInt("limit", t, i)
}

// Aggregate returns a new gremlin Traversal
func (t *Traversal) Aggregate(s string) *Traversal {
	return processString("aggregate", t, s)
}

// ValueMap The valueMap() step yields a map representation of the properties of an element.
func (t *Traversal) ValueMap() *Traversal {
	return process("valueMap", t)
}

// Dedup returns a new gremlin Traversal
func (t *Traversal) Dedup() *Traversal {
	return process("dedup", t)
}

// Read returns a new gremlin Traversal
func (t *Traversal) Read() *Traversal {
	return process("read", t)
}

// Iterate returns a new gremlin Traversal
func (t *Traversal) Iterate() *Traversal {
	return process("iterate", t)
}

// Of returns a new gremlin Traversal
func (t *Traversal) Of() *Traversal {
	return process("of", t)
}

// Fold returns a new gremlin Traversal
func (t *Traversal) Fold() *Traversal {
	return process("fold", t)
}

// GroupCount returns a new gremlin Traversal
func (t *Traversal) GroupCount() *Traversal {
	return process("groupCount", t)
}

// Group returns a new gremlin Traversal
func (t *Traversal) Group() *Traversal {
	return process("group", t)
}

////////////////
// Vertex Steps
///////////////

// Out step moves the Traversal to the outgoing adjacent vertices, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (t *Traversal) Out(name ...string) *Traversal {
	return processStringSlice("out", t, name...)
}

// Has step is a filter step. It is the most common step used for graph Traversals,
// since this step narrows the query to find particular vertices or edges with certain property values.
func (t *Traversal) Has(name string, values ...string) *Traversal {
	return processStringAndStringSlice("has", t, name, values...)
}

// OutE step moves the Traversal to the outgoing incident edges, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (t *Traversal) OutE(name ...string) *Traversal {
	return processStringSlice("outE", t, name...)
}

// OutV The out() step moves the Traversal to the outgoing adjacent vertices,
// given the edge labels. Specifying no edge label will traverse all incident edges.
func (t *Traversal) OutV() *Traversal {
	return process("outV", t)
}

// In step moves the Traversal to the incoming adjacent vertices, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (t *Traversal) In(name ...string) *Traversal {
	return processStringSlice("in", t, name...)
}

// InE step moves the Traversal to the incoming incident edges, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (t *Traversal) InE(name ...string) *Traversal {
	return processStringSlice("inE", t, name...)
}

// InV step moves the Traversal to the incoming vertices.
func (t *Traversal) InV() *Traversal {
	return process("inV", t)
}

// Both step moves the Traversal to both the outgoing and the incoming adjacent vertices, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (t *Traversal) Both(name ...string) *Traversal {
	return processStringSlice("both", t, name...)
}

// BothE step moves the Traversal to both the outgoing and the incoming incident edges, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (t *Traversal) BothE(name ...string) *Traversal {
	return processStringSlice("bothE", t, name...)
}

// BothV step moves the Traversal to both the outgoing and the incoming vertices.
func (t *Traversal) BothV() *Traversal {
	return process("bothV", t)
}

// OtherV step moves the Traversal to the vertex that was not the vertex that was moved from.
func (t *Traversal) OtherV() *Traversal {
	return process("otherV", t)
}

// Match step is a map step provides a more declarative form of graph querying based on pattern matching.
func (t *Traversal) Match(i ...interface{}) *Traversal {
	return processInterfaceSlice("match", t, i...)
}

////////////////
// System
///////////////

// Exists returns a new gremlin Traversal
func (t *Traversal) Exists() *Traversal {
	return process("exists", t)
}

// Create returns a new gremlin Traversal
func (t *Traversal) Create() *Traversal {
	return process("create", t)
}

// Drop returns a new gremlin Traversal
func (t *Traversal) Drop() *Traversal {
	return process("drop", t)
}

// Graphs discover what graphs currently exist using this command.
func (t *Traversal) Graphs() *Traversal {
	return process("graphs", t)
}

// Set returns a new gremlin Traversal
func (t *Traversal) Set(s string) *Traversal {
	return processString("set", t, s)
}

////////////////
// Predicates
///////////////

// Eq predicate answers the question: Is the incoming object equal to the provided object?
func (t *Traversal) Eq(i interface{}) *Traversal {
	return processInterface("eq", t, i)
}

// Neg predicate answers the question: Is the incoming object equal to the provided object?
func (t *Traversal) Neg(i interface{}) *Traversal {
	return processInterface("neg", t, i)
}

// Lt predicate answers the question: Is the incoming number less than the provided number?
func (t *Traversal) Lt(i int) *Traversal {
	return processInt("lt", t, i)
}

// Lte predicate answers the question: Is the incoming number less than or equal to the provided number?
func (t *Traversal) Lte(i int) *Traversal {
	return processInt("lte", t, i)
}

// Gt  predicate answers the question: Is the incoming number greater than the provided number?
func (t *Traversal) Gt(i int) *Traversal {
	return processInt("gt", t, i)
}

// Gte predicate answers the question: Is the incoming number greater than or equal to the provided number?
func (t *Traversal) Gte(i int) *Traversal {
	return processInt("gte", t, i)
}

// Inside predicate answers the question: Is the incoming number greater than the first provided number and less than the second?
func (t *Traversal) Inside(from int, to int) *Traversal {
	return processFromTo("inside", t, from, to)
}

// Outside predicate answers the question: Is the incoming number less than the first provided number or greater than the second?
func (t *Traversal) Outside(from int, to int) *Traversal {
	return processFromTo("Outside", t, from, to)
}

// Between predicate answers the question: Is the incoming number less than the first provided number or greater than the second?
func (t *Traversal) Between(from int, to int) *Traversal {
	return processFromTo("between", t, from, to)
}

// Within predicate answers the question: Is the incoming object in the array of provided objects?
func (t *Traversal) Within(name ...string) *Traversal {
	return processStringSlice("within", t, name...)
}

// Without predicate answers the question: Is the incoming object not in the array of provided objects?
func (t *Traversal) Without(name ...string) *Traversal {
	return processStringSlice("without", t, name...)
}

////////////////
// Step-modulators
///////////////

// As step is a step modulator, a helper step for another Traversal step.
func (t *Traversal) As(s string) *Traversal {
	return processString("as", t, s)
}

// By step is a step modulator that can be used to modify sort order from a previous step.
func (t *Traversal) By(i interface{}) *Traversal {
	return processInterface("by", t, i)
}

// From step is a step modulator, a helper step for another Traversal step.
// Its main use is to designate the outgoing vertex for an addE() step. Generally, a to() step is paired with a from() step.
func (t *Traversal) From(i interface{}) *Traversal {
	return processInterface("from", t, i)
}

// Option step modulator is used in conjunction with branch()
// or choose(), and specifies the returned values based on values found at that point in a Traversal.
func (t *Traversal) Option(s string) *Traversal {
	return processString("option", t, s)
}

// Times step is a step modulator, a helper step for another Traversal step.
// Its main use is to repeat a repeat() step for the specified number of times.
func (t *Traversal) Times(n int) *Traversal {
	return processInt("times", t, n)
}

// To returns a new gremlin Traversal
func (t *Traversal) To(s string) *Traversal {
	return processString("to", t, s)
}

// Next  step is a terminal step. It returns the next number of steps, based on a supplied integer value.
func (t *Traversal) Next(i ...int) *Traversal {
	return processIntSlice("next", t, i...)
}

// HasNext step is a terminal step. It determines whether or not there are available results from a traversal,
// returning a Boolean value of true or false.
func (t *Traversal) HasNext() *Traversal {
	return process("hasNext", t)
}

// Emit step is a step modulator, a helper step for another Traversal step.
// Its main use is to emit either incoming Traversals before a repeat() step, or emit outgoing Traversals after a repeat() step.
// The emission sends a copy of the current objects to the next step in the query.
// A predicate or Traversal can be used in an emit() step to cause the emission only if the predicate or Traversal is true.
func (t *Traversal) Emit(i ...interface{}) *Traversal {
	return processInterfaceOptionalSlice("emit", t, i)
}

// Until step is a step modulator, a helper step for another Traversal step.
// Its main use is to turn a repeat() step into a do-while loop (if used after the repeat() step) or a while-do loop (if used before the repeat() step).
// A predicate or Traversal can be used in an until() step to cause the loop to complete only if the predicate or Traversal is true.
func (t *Traversal) Until(i ...interface{}) *Traversal {
	return processInterfaceOptionalSlice("until", t, i)
}
