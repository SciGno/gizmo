package gizmo

import (
	"fmt"
	"strconv"
	"strings"
)

// Traversal struct
type Traversal struct {
	value string
}

// GraphTraversal function
func graphTraversal(name ...string) *Traversal {
	if len(name) == 0 {
		return &Traversal{""}
	}
	return &Traversal{name[0] + "."}
}

// New function
func (b *Traversal) New(name ...string) *Traversal {
	return graphTraversal(name...)
}

// String function
func (b *Traversal) String() string {
	if b.value[len(b.value)-1:] == "." {
		b.value = b.value[:len(b.value)-1]
	}
	return b.value
}

// V function
func (b *Traversal) V(name ...string) *Traversal {
	if len(name) == 0 {
		b.value = b.value + "V()."
	} else {
		b.value = b.value + "V('" + name[0] + "')."
	}
	return b
}

// AddV step is a map/sideEffect. A vertex is added from a Traversal g using addV.
// A previously created vertex label must be specified. Property key-value pairs may be optionally specified.
func (b *Traversal) AddV(name ...string) *Traversal {
	if len(name) == 0 {
		b.value = b.value + "addV()."
	} else {
		b.value = b.value + "addV('" + name[0] + "')."
	}
	return b
}

// Where returns a new gremlin Traversal
func (b *Traversal) Where(i interface{}) *Traversal {
	if _, ok := i.(*Traversal); ok {
		b.value = b.value + "where(" + i.(*Traversal).String() + ")."
	} else {
		b.value = b.value + "where(" + i.(string) + ")."
	}
	return b
}

// And step can take an arbitrary number of Traversals to filter the returned objects using a boolean AND function.
// The related boolean OR function can be done with the or() step.
func (b *Traversal) And(i interface{}) *Traversal {
	if _, ok := i.(*Traversal); ok {
		b.value = b.value + "and(" + i.(*Traversal).String() + ")."
	} else {
		b.value = b.value + "and('" + i.(string) + "')."
	}
	return b
}

// FlatMap step is a general step. This step creates a map of the Traversal objects and streams the map to the next Traversal step.
func (b *Traversal) FlatMap(i interface{}) *Traversal {
	if _, ok := i.(*Traversal); ok {
		b.value = b.value + "flatMap(" + i.(*Traversal).String() + ")."
	} else {
		b.value = b.value + "flatMap(" + i.(string) + ")."
	}
	return b
}

// AddE is a map/sideEffect. An edge is added from a Traversal g using addE between two existing vertices.
// A previously created edge label must be specified.
func (b *Traversal) AddE(e string) *Traversal {
	b.value = b.value + "addE('" + e + "')."
	return b
}

// Property step is a sideEffect. When a vertex is added from a Traversal g using addV, properties can be added with property().
// For a previously created vertex, properties can be added. A previously created vertex label must be specified.
// Property key-value pairs are specified.
func (b *Traversal) Property(k interface{}, v interface{}) *Traversal {

	if _, ok := k.(string); ok {
		k = "'" + k.(string) + "'"
	}

	if _, ok := v.(string); ok {
		v = "'" + v.(string) + "'"
	}

	b.value = b.value + fmt.Sprintf("property(%v, %v).", k, v)
	return b
}

// Limit returns a new gremlin Traversal
func (b *Traversal) Limit(i int) *Traversal {
	b.value = b.value + "limit(" + strconv.Itoa(i) + ")."
	return b
}

// Aggregate returns a new gremlin Traversal
func (b *Traversal) Aggregate(s string) *Traversal {
	b.value = b.value + "aggregate('" + s + "')."
	return b
}

// ValueMap The valueMap() step yields a map representation of the properties of an element.
func (b *Traversal) ValueMap() *Traversal {
	b.value = b.value + "valueMap()."
	return b
}

// Dedup returns a new gremlin Traversal
func (b *Traversal) Dedup() *Traversal {
	b.value = b.value + "dedup()."
	return b
}

// Read returns a new gremlin Traversal
func (b *Traversal) Read() *Traversal {
	b.value = b.value + "read()."
	return b
}

// Iterate returns a new gremlin Traversal
func (b *Traversal) Iterate() *Traversal {
	b.value = b.value + "iterate()."
	return b
}

// Of returns a new gremlin Traversal
func (b *Traversal) Of() *Traversal {
	b.value = b.value + "of()."
	return b
}

// Fold returns a new gremlin Traversal
func (b *Traversal) Fold() *Traversal {
	b.value = b.value + "fold()."
	return b
}

// GroupCount returns a new gremlin Traversal
func (b *Traversal) GroupCount() *Traversal {
	b.value = b.value + "groupCount()."
	return b
}

// Group returns a new gremlin Traversal
func (b *Traversal) Group() *Traversal {
	b.value = b.value + "group()."
	return b
}

////////////////
// Vertex Steps
///////////////

// Out step moves the Traversal to the outgoing adjacent vertices, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (b *Traversal) Out(name ...string) *Traversal {
	if len(name) == 0 {
		b.value = b.value + "out()."
	} else {
		b.value = b.value + "out('" + name[0] + "')."
	}
	return b
}

// Has step is a filter step. It is the most common step used for graph Traversals,
// since this step narrows the query to find particular vertices or edges with certain property values.
func (b *Traversal) Has(name string, values ...string) *Traversal {
	if len(values) == 0 {
		b.value = b.value + "has('" + name + "')."
	} else {
		b.value = b.value + "has('" + name + "','" + strings.Join(values, "','") + "')."
	}
	return b
}

// OutE step moves the Traversal to the outgoing incident edges, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (b *Traversal) OutE(name ...string) *Traversal {
	if len(name) == 0 {
		b.value = b.value + "outE()."
	} else {
		b.value = b.value + "outE('" + name[0] + "')."
	}
	return b
}

// OutV The out() step moves the Traversal to the outgoing adjacent vertices,
// given the edge labels. Specifying no edge label will traverse all incident edges.
func (b *Traversal) OutV() *Traversal {
	b.value = b.value + "outV()."
	return b
}

// In step moves the Traversal to the incoming adjacent vertices, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (b *Traversal) In(name ...string) *Traversal {
	if len(name) == 0 {
		b.value = b.value + "in()."
	} else {
		b.value = b.value + "in('" + name[0] + "')."
	}
	return b
}

// InE step moves the Traversal to the incoming incident edges, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (b *Traversal) InE(name ...string) *Traversal {
	if len(name) == 0 {
		b.value = b.value + "inE()."
	} else {
		b.value = b.value + "inE('" + name[0] + "')."
	}
	return b
}

// InV step moves the Traversal to the incoming vertices.
func (b *Traversal) InV() *Traversal {
	b.value = b.value + "inV()."
	return b
}

// Both step moves the Traversal to both the outgoing and the incoming adjacent vertices, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (b *Traversal) Both(name ...string) *Traversal {
	if len(name) == 0 {
		b.value = b.value + "both()."
	} else {
		b.value = b.value + "both('" + name[0] + "')."
	}
	return b
}

// BothE step moves the Traversal to both the outgoing and the incoming incident edges, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (b *Traversal) BothE(name ...string) *Traversal {
	if len(name) == 0 {
		b.value = b.value + "bothE()."
	} else {
		b.value = b.value + "bothE('" + name[0] + "')."
	}
	return b
}

// BothV step moves the Traversal to both the outgoing and the incoming vertices.
func (b *Traversal) BothV() *Traversal {
	b.value = b.value + "bothV()."
	return b
}

// OtherV step moves the Traversal to the vertex that was not the vertex that was moved from.
func (b *Traversal) OtherV() *Traversal {
	b.value = b.value + "otherV()."
	return b
}

// Match step is a map step provides a more declarative form of graph querying based on pattern matching.
func (b *Traversal) Match(i ...interface{}) *Traversal {
	if len(i) > 0 {
		var arr []string
		for _, v := range i {
			if _, ok := v.(*Traversal); ok {
				arr = append(arr, v.(*Traversal).String())
			}
		}
		b.value = b.value + "match(" + strings.Join(arr, ",") + ")."
	} else {
		b.value = b.value + "match()."
	}
	return b
}

////////////////
// System
///////////////

// Exists returns a new gremlin Traversal
func (b *Traversal) Exists() *Traversal {
	b.value = b.value + "exists()."
	return b
}

// Create returns a new gremlin Traversal
func (b *Traversal) Create() *Traversal {
	b.value = b.value + "create()."
	return b
}

// Drop returns a new gremlin Traversal
func (b *Traversal) Drop() *Traversal {
	b.value = b.value + "drop()."
	return b
}

// Graphs discover what graphs currently exist using this command.
func (b *Traversal) Graphs() *Traversal {
	b.value = b.value + "graphs()."
	return b
}

// Set returns a new gremlin Traversal
func (b *Traversal) Set(s string) *Traversal {
	b.value = b.value + "set(\"" + s + "\")."
	return b
}

// Replication returns a new gremlin Traversal
func (b *Traversal) Replication(s string) *Traversal {
	b.value = b.value + "replication(\"" + s + "\")."
	return b
}

// SystemReplication returns a new gremlin Traversal
func (b *Traversal) SystemReplication(s string) *Traversal {
	b.value = b.value + "systemReplication(\"" + s + "\")."
	return b
}

////////////////
// Predicates
///////////////

// Eq predicate answers the question: Is the incoming object equal to the provided object?
func (b *Traversal) Eq(i interface{}) *Traversal {
	if _, ok := i.(*Traversal); ok {
		b.value = b.value + "eq(" + i.(*Traversal).String() + ")."
	} else {
		b.value = b.value + "eq(" + i.(string) + ")."
	}
	return b
}

// Neg predicate answers the question: Is the incoming object equal to the provided object?
func (b *Traversal) Neg(i interface{}) *Traversal {
	if _, ok := i.(*Traversal); ok {
		b.value = b.value + "neg(" + i.(*Traversal).String() + ")."
	} else {
		b.value = b.value + "neg(" + i.(string) + ")."
	}
	return b
}

// Lt predicate answers the question: Is the incoming number less than the provided number?
func (b *Traversal) Lt(i int) *Traversal {
	b.value = b.value + "lt(" + strconv.Itoa(i) + ")."
	return b
}

// Lte predicate answers the question: Is the incoming number less than or equal to the provided number?
func (b *Traversal) Lte(i int) *Traversal {
	b.value = b.value + "lte(" + strconv.Itoa(i) + ")."
	return b
}

// Gt  predicate answers the question: Is the incoming number greater than the provided number?
func (b *Traversal) Gt(i int) *Traversal {
	b.value = b.value + "gt(" + strconv.Itoa(i) + ")."
	return b
}

// Gte predicate answers the question: Is the incoming number greater than or equal to the provided number?
func (b *Traversal) Gte(i int) *Traversal {
	b.value = b.value + "gte(" + strconv.Itoa(i) + ")."
	return b
}

// Inside predicate answers the question: Is the incoming number greater than the first provided number and less than the second?
func (b *Traversal) Inside(from int, to int) *Traversal {
	b.value = b.value + "inside(" + strconv.Itoa(from) + "," + strconv.Itoa(to) + ")."
	return b
}

// Outside predicate answers the question: Is the incoming number less than the first provided number or greater than the second?
func (b *Traversal) Outside(from int, to int) *Traversal {
	b.value = b.value + "outside(" + strconv.Itoa(from) + "," + strconv.Itoa(to) + ")."
	return b
}

// Between predicate answers the question: Is the incoming number less than the first provided number or greater than the second?
func (b *Traversal) Between(from int, to int) *Traversal {
	b.value = b.value + "between(" + strconv.Itoa(from) + "," + strconv.Itoa(to) + ")."
	return b
}

// Within predicate answers the question: Is the incoming object in the array of provided objects?
func (b *Traversal) Within(name ...string) *Traversal {
	b.value = b.value + "within('" + strings.Join(name, "','") + "')."
	return b
}

// Without predicate answers the question: Is the incoming object not in the array of provided objects?
func (b *Traversal) Without(name ...string) *Traversal {
	b.value = b.value + "without('" + strings.Join(name, "','") + "')."
	return b
}

////////////////
// Step-modulators
///////////////

// As step is a step modulator, a helper step for another Traversal step.
func (b *Traversal) As(s string) *Traversal {
	b.value = b.value + "as('" + s + "')."
	return b
}

// By step is a step modulator that can be used to modify sort order from a previous step.
func (b *Traversal) By(i interface{}) *Traversal {
	if _, ok := i.(*Traversal); ok {
		b.value = b.value + "by(" + i.(*Traversal).String() + ")."
	} else {
		b.value = b.value + "by(" + i.(string) + ")."
	}
	return b
}

// From step is a step modulator, a helper step for another Traversal step.
// Its main use is to designate the outgoing vertex for an addE() step. Generally, a to() step is paired with a from() step.
func (b *Traversal) From(i interface{}) *Traversal {
	if _, ok := i.(*Traversal); ok {
		b.value = b.value + "from(" + i.(*Traversal).String() + ")."
	} else {
		b.value = b.value + "from(" + i.(string) + ")."
	}
	return b
}

// Option step modulator is used in conjunction with branch()
// or choose(), and specifies the returned values based on values found at that point in a Traversal.
func (b *Traversal) Option(s string) *Traversal {
	b.value = b.value + "option('" + s + "')."
	return b
}

// Times step is a step modulator, a helper step for another Traversal step.
// Its main use is to repeat a repeat() step for the specified number of times.
func (b *Traversal) Times(n int) *Traversal {
	b.value = b.value + "times(" + strconv.Itoa(n) + ")."
	return b
}

// To returns a new gremlin Traversal
func (b *Traversal) To(s string) *Traversal {
	b.value = b.value + "to('" + s + "')."
	return b
}

// Next  step is a terminal step. It returns the next number of steps, based on a supplied integer value.
func (b *Traversal) Next(i ...int) *Traversal {
	if len(i) > 0 {
		b.value = b.value + "next(" + strconv.Itoa(i[0]) + ")."
	} else {
		b.value = b.value + "next()."
	}
	return b
}

// HasNext step is a terminal step. It determines whether or not there are available results from a traversal,
// returning a Boolean value of true or false.
func (b *Traversal) HasNext() *Traversal {
	b.value = b.value + "hasNext()."
	return b
}

// Emit step is a step modulator, a helper step for another Traversal step.
// Its main use is to emit either incoming Traversals before a repeat() step, or emit outgoing Traversals after a repeat() step.
// The emission sends a copy of the current objects to the next step in the query.
// A predicate or Traversal can be used in an emit() step to cause the emission only if the predicate or Traversal is true.
func (b *Traversal) Emit(i ...interface{}) *Traversal {
	if len(i) > 0 {
		if _, ok := i[0].(*Traversal); ok {
			b.value = b.value + "emit(" + i[0].(*Traversal).String() + ")."
		} else {
			b.value = b.value + "emit(" + i[0].(string) + ")."
		}
	} else {
		b.value = b.value + "emit()."
	}
	return b
}

// Until step is a step modulator, a helper step for another Traversal step.
// Its main use is to turn a repeat() step into a do-while loop (if used after the repeat() step) or a while-do loop (if used before the repeat() step).
// A predicate or Traversal can be used in an until() step to cause the loop to complete only if the predicate or Traversal is true.
func (b *Traversal) Until(i ...interface{}) *Traversal {
	if len(i) > 0 {
		if _, ok := i[0].(*Traversal); ok {
			b.value = b.value + "until(" + i[0].(*Traversal).String() + ")."
		} else {
			b.value = b.value + "until(" + i[0].(string) + ")."
		}
	} else {
		b.value = b.value + "until()."
	}
	return b
}

////////////////
// Helpers
///////////////

// TernaryOp step is a sideEffect. When a vertex is added from a Traversal g using addV, properties can be added with property().
// For a previously created vertex, properties can be added. A previously created vertex label must be specified.
// Property key-value pairs are specified.
func (b *Traversal) TernaryOp(condition interface{}, v1 interface{}, v2 interface{}) *Traversal {

	if _, ok := condition.(string); ok {
		condition = "'" + condition.(string) + "'"
	}

	if _, ok := v1.(string); ok {
		v1 = "'" + v1.(string) + "'"
	}

	if _, ok := v2.(string); ok {
		v2 = "'" + v2.(string) + "'"
	}

	b.value = b.value + fmt.Sprintf("%v ? %v : %v", condition, v1, v2)
	return b
}

// Raw function is a helper function that can used in case a step is not supported
// or not available in this package. the string is appended as is without modification
func (b *Traversal) Raw(s string) *Traversal {
	b.value = b.value + s
	return b
}

// Append function adds the provided string to the end of the traversal query.
func (b *Traversal) Append(s string) *Traversal {
	b.value = b.value + s
	return b
}

// AddLine function adds a new line with the provided string to the traversal query.
func (b *Traversal) AddLine(s string) *Traversal {
	b.value = b.value + "\n" + s + "."
	return b
}
