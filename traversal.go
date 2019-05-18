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
	if len(name) == 0 {
		t.value = t.value + "V()."
	} else {
		t.value = t.value + "V('" + name[0] + "')."
	}
	return t
}

// AddV step is a map/sideEffect. A vertex is added from a Traversal g using addV.
// A previously created vertex label must be specified. Property key-value pairs may be optionally specified.
func (t *Traversal) AddV(name ...string) *Traversal {
	if len(name) == 0 {
		t.value = t.value + "addV()."
	} else {
		t.value = t.value + "addV('" + name[0] + "')."
	}
	return t
}

// Where returns a new gremlin Traversal
func (t *Traversal) Where(i interface{}) *Traversal {
	if _, ok := i.(*Traversal); ok {
		t.value = t.value + "where(" + i.(*Traversal).String() + ")."
	} else {
		t.value = t.value + "where(" + i.(string) + ")."
	}
	return t
}

// And step can take an arbitrary number of Traversals to filter the returned objects using a boolean AND function.
// The related boolean OR function can be done with the or() step.
func (t *Traversal) And(i interface{}) *Traversal {
	if _, ok := i.(*Traversal); ok {
		t.value = t.value + "and(" + i.(*Traversal).String() + ")."
	} else {
		t.value = t.value + "and('" + i.(string) + "')."
	}
	return t
}

// FlatMap step is a general step. This step creates a map of the Traversal objects and streams the map to the next Traversal step.
func (t *Traversal) FlatMap(i interface{}) *Traversal {
	if _, ok := i.(*Traversal); ok {
		t.value = t.value + "flatMap(" + i.(*Traversal).String() + ")."
	} else {
		t.value = t.value + "flatMap(" + i.(string) + ")."
	}
	return t
}

// AddE is a map/sideEffect. An edge is added from a Traversal g using addE between two existing vertices.
// A previously created edge label must be specified.
func (t *Traversal) AddE(e string) *Traversal {
	t.value = t.value + "addE('" + e + "')."
	return t
}

// Property step in a traversal has the same beahviour as in the Graph API
func (t *Traversal) Property(k interface{}, v interface{}) *Traversal {

	if _, ok := k.(string); ok {
		k = "'" + k.(string) + "'"
	}

	if _, ok := v.(string); ok {
		v = "'" + v.(string) + "'"
	}

	t.value = t.value + fmt.Sprintf("property(%v, %v).", k, v)
	return t
}

// Limit returns a new gremlin Traversal
func (t *Traversal) Limit(i int) *Traversal {
	t.value = t.value + "limit(" + strconv.Itoa(i) + ")."
	return t
}

// Aggregate returns a new gremlin Traversal
func (t *Traversal) Aggregate(s string) *Traversal {
	t.value = t.value + "aggregate('" + s + "')."
	return t
}

// ValueMap The valueMap() step yields a map representation of the properties of an element.
func (t *Traversal) ValueMap() *Traversal {
	t.value = t.value + "valueMap()."
	return t
}

// Dedup returns a new gremlin Traversal
func (t *Traversal) Dedup() *Traversal {
	t.value = t.value + "dedup()."
	return t
}

// Read returns a new gremlin Traversal
func (t *Traversal) Read() *Traversal {
	t.value = t.value + "read()."
	return t
}

// Iterate returns a new gremlin Traversal
func (t *Traversal) Iterate() *Traversal {
	t.value = t.value + "iterate()."
	return t
}

// Of returns a new gremlin Traversal
func (t *Traversal) Of() *Traversal {
	t.value = t.value + "of()."
	return t
}

// Fold returns a new gremlin Traversal
func (t *Traversal) Fold() *Traversal {
	t.value = t.value + "fold()."
	return t
}

// GroupCount returns a new gremlin Traversal
func (t *Traversal) GroupCount() *Traversal {
	t.value = t.value + "groupCount()."
	return t
}

// Group returns a new gremlin Traversal
func (t *Traversal) Group() *Traversal {
	t.value = t.value + "group()."
	return t
}

////////////////
// Vertex Steps
///////////////

// Out step moves the Traversal to the outgoing adjacent vertices, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (t *Traversal) Out(name ...string) *Traversal {
	if len(name) == 0 {
		t.value = t.value + "out()."
	} else {
		t.value = t.value + "out('" + name[0] + "')."
	}
	return t
}

// Has step is a filter step. It is the most common step used for graph Traversals,
// since this step narrows the query to find particular vertices or edges with certain property values.
func (t *Traversal) Has(name string, values ...string) *Traversal {
	if len(values) == 0 {
		t.value = t.value + "has('" + name + "')."
	} else {
		t.value = t.value + "has('" + name + "','" + strings.Join(values, "','") + "')."
	}
	return t
}

// OutE step moves the Traversal to the outgoing incident edges, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (t *Traversal) OutE(name ...string) *Traversal {
	if len(name) == 0 {
		t.value = t.value + "outE()."
	} else {
		t.value = t.value + "outE('" + name[0] + "')."
	}
	return t
}

// OutV The out() step moves the Traversal to the outgoing adjacent vertices,
// given the edge labels. Specifying no edge label will traverse all incident edges.
func (t *Traversal) OutV() *Traversal {
	t.value = t.value + "outV()."
	return t
}

// In step moves the Traversal to the incoming adjacent vertices, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (t *Traversal) In(name ...string) *Traversal {
	if len(name) == 0 {
		t.value = t.value + "in()."
	} else {
		t.value = t.value + "in('" + name[0] + "')."
	}
	return t
}

// InE step moves the Traversal to the incoming incident edges, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (t *Traversal) InE(name ...string) *Traversal {
	if len(name) == 0 {
		t.value = t.value + "inE()."
	} else {
		t.value = t.value + "inE('" + name[0] + "')."
	}
	return t
}

// InV step moves the Traversal to the incoming vertices.
func (t *Traversal) InV() *Traversal {
	t.value = t.value + "inV()."
	return t
}

// Both step moves the Traversal to both the outgoing and the incoming adjacent vertices, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (t *Traversal) Both(name ...string) *Traversal {
	if len(name) == 0 {
		t.value = t.value + "both()."
	} else {
		t.value = t.value + "both('" + name[0] + "')."
	}
	return t
}

// BothE step moves the Traversal to both the outgoing and the incoming incident edges, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (t *Traversal) BothE(name ...string) *Traversal {
	if len(name) == 0 {
		t.value = t.value + "bothE()."
	} else {
		t.value = t.value + "bothE('" + name[0] + "')."
	}
	return t
}

// BothV step moves the Traversal to both the outgoing and the incoming vertices.
func (t *Traversal) BothV() *Traversal {
	t.value = t.value + "bothV()."
	return t
}

// OtherV step moves the Traversal to the vertex that was not the vertex that was moved from.
func (t *Traversal) OtherV() *Traversal {
	t.value = t.value + "otherV()."
	return t
}

// Match step is a map step provides a more declarative form of graph querying based on pattern matching.
func (t *Traversal) Match(i ...interface{}) *Traversal {
	if len(i) > 0 {
		var arr []string
		for _, v := range i {
			if _, ok := v.(*Traversal); ok {
				arr = append(arr, v.(*Traversal).String())
			}
		}
		t.value = t.value + "match(" + strings.Join(arr, ",") + ")."
	} else {
		t.value = t.value + "match()."
	}
	return t
}

////////////////
// System
///////////////

// Exists returns a new gremlin Traversal
func (t *Traversal) Exists() *Traversal {
	t.value = t.value + "exists()."
	return t
}

// Create returns a new gremlin Traversal
func (t *Traversal) Create() *Traversal {
	t.value = t.value + "create()."
	return t
}

// Drop returns a new gremlin Traversal
func (t *Traversal) Drop() *Traversal {
	t.value = t.value + "drop()."
	return t
}

// Graphs discover what graphs currently exist using this command.
func (t *Traversal) Graphs() *Traversal {
	t.value = t.value + "graphs()."
	return t
}

// Set returns a new gremlin Traversal
func (t *Traversal) Set(s string) *Traversal {
	t.value = t.value + "set(\"" + s + "\")."
	return t
}

// Replication returns a new gremlin Traversal
func (t *Traversal) Replication(s string) *Traversal {
	t.value = t.value + "replication(\"" + s + "\")."
	return t
}

// SystemReplication returns a new gremlin Traversal
func (t *Traversal) SystemReplication(s string) *Traversal {
	t.value = t.value + "systemReplication(\"" + s + "\")."
	return t
}

////////////////
// Predicates
///////////////

// Eq predicate answers the question: Is the incoming object equal to the provided object?
func (t *Traversal) Eq(i interface{}) *Traversal {
	if _, ok := i.(*Traversal); ok {
		t.value = t.value + "eq(" + i.(*Traversal).String() + ")."
	} else {
		t.value = t.value + "eq(" + i.(string) + ")."
	}
	return t
}

// Neg predicate answers the question: Is the incoming object equal to the provided object?
func (t *Traversal) Neg(i interface{}) *Traversal {
	if _, ok := i.(*Traversal); ok {
		t.value = t.value + "neg(" + i.(*Traversal).String() + ")."
	} else {
		t.value = t.value + "neg(" + i.(string) + ")."
	}
	return t
}

// Lt predicate answers the question: Is the incoming number less than the provided number?
func (t *Traversal) Lt(i int) *Traversal {
	t.value = t.value + "lt(" + strconv.Itoa(i) + ")."
	return t
}

// Lte predicate answers the question: Is the incoming number less than or equal to the provided number?
func (t *Traversal) Lte(i int) *Traversal {
	t.value = t.value + "lte(" + strconv.Itoa(i) + ")."
	return t
}

// Gt  predicate answers the question: Is the incoming number greater than the provided number?
func (t *Traversal) Gt(i int) *Traversal {
	t.value = t.value + "gt(" + strconv.Itoa(i) + ")."
	return t
}

// Gte predicate answers the question: Is the incoming number greater than or equal to the provided number?
func (t *Traversal) Gte(i int) *Traversal {
	t.value = t.value + "gte(" + strconv.Itoa(i) + ")."
	return t
}

// Inside predicate answers the question: Is the incoming number greater than the first provided number and less than the second?
func (t *Traversal) Inside(from int, to int) *Traversal {
	t.value = t.value + "inside(" + strconv.Itoa(from) + "," + strconv.Itoa(to) + ")."
	return t
}

// Outside predicate answers the question: Is the incoming number less than the first provided number or greater than the second?
func (t *Traversal) Outside(from int, to int) *Traversal {
	t.value = t.value + "outside(" + strconv.Itoa(from) + "," + strconv.Itoa(to) + ")."
	return t
}

// Between predicate answers the question: Is the incoming number less than the first provided number or greater than the second?
func (t *Traversal) Between(from int, to int) *Traversal {
	t.value = t.value + "between(" + strconv.Itoa(from) + "," + strconv.Itoa(to) + ")."
	return t
}

// Within predicate answers the question: Is the incoming object in the array of provided objects?
func (t *Traversal) Within(name ...string) *Traversal {
	t.value = t.value + "within('" + strings.Join(name, "','") + "')."
	return t
}

// Without predicate answers the question: Is the incoming object not in the array of provided objects?
func (t *Traversal) Without(name ...string) *Traversal {
	t.value = t.value + "without('" + strings.Join(name, "','") + "')."
	return t
}

////////////////
// Step-modulators
///////////////

// As step is a step modulator, a helper step for another Traversal step.
func (t *Traversal) As(s string) *Traversal {
	t.value = t.value + "as('" + s + "')."
	return t
}

// By step is a step modulator that can be used to modify sort order from a previous step.
func (t *Traversal) By(i interface{}) *Traversal {
	if _, ok := i.(*Traversal); ok {
		t.value = t.value + "by(" + i.(*Traversal).String() + ")."
	} else {
		t.value = t.value + "by(" + i.(string) + ")."
	}
	return t
}

// From step is a step modulator, a helper step for another Traversal step.
// Its main use is to designate the outgoing vertex for an addE() step. Generally, a to() step is paired with a from() step.
func (t *Traversal) From(i interface{}) *Traversal {
	if _, ok := i.(*Traversal); ok {
		t.value = t.value + "from(" + i.(*Traversal).String() + ")."
	} else {
		t.value = t.value + "from(" + i.(string) + ")."
	}
	return t
}

// Option step modulator is used in conjunction with branch()
// or choose(), and specifies the returned values based on values found at that point in a Traversal.
func (t *Traversal) Option(s string) *Traversal {
	t.value = t.value + "option('" + s + "')."
	return t
}

// Times step is a step modulator, a helper step for another Traversal step.
// Its main use is to repeat a repeat() step for the specified number of times.
func (t *Traversal) Times(n int) *Traversal {
	t.value = t.value + "times(" + strconv.Itoa(n) + ")."
	return t
}

// To returns a new gremlin Traversal
func (t *Traversal) To(s string) *Traversal {
	t.value = t.value + "to('" + s + "')."
	return t
}

// Next  step is a terminal step. It returns the next number of steps, based on a supplied integer value.
func (t *Traversal) Next(i ...int) *Traversal {
	if len(i) > 0 {
		t.value = t.value + "next(" + strconv.Itoa(i[0]) + ")."
	} else {
		t.value = t.value + "next()."
	}
	return t
}

// HasNext step is a terminal step. It determines whether or not there are available results from a traversal,
// returning a Boolean value of true or false.
func (t *Traversal) HasNext() *Traversal {
	t.value = t.value + "hasNext()."
	return t
}

// Emit step is a step modulator, a helper step for another Traversal step.
// Its main use is to emit either incoming Traversals before a repeat() step, or emit outgoing Traversals after a repeat() step.
// The emission sends a copy of the current objects to the next step in the query.
// A predicate or Traversal can be used in an emit() step to cause the emission only if the predicate or Traversal is true.
func (t *Traversal) Emit(i ...interface{}) *Traversal {
	if len(i) > 0 {
		if _, ok := i[0].(*Traversal); ok {
			t.value = t.value + "emit(" + i[0].(*Traversal).String() + ")."
		} else {
			t.value = t.value + "emit(" + i[0].(string) + ")."
		}
	} else {
		t.value = t.value + "emit()."
	}
	return t
}

// Until step is a step modulator, a helper step for another Traversal step.
// Its main use is to turn a repeat() step into a do-while loop (if used after the repeat() step) or a while-do loop (if used before the repeat() step).
// A predicate or Traversal can be used in an until() step to cause the loop to complete only if the predicate or Traversal is true.
func (t *Traversal) Until(i ...interface{}) *Traversal {
	if len(i) > 0 {
		if _, ok := i[0].(*Traversal); ok {
			t.value = t.value + "until(" + i[0].(*Traversal).String() + ")."
		} else {
			t.value = t.value + "until(" + i[0].(string) + ")."
		}
	} else {
		t.value = t.value + "until()."
	}
	return t
}

////////////////
// Helpers
///////////////

// TernaryOp step is a sideEffect.
func (t *Traversal) TernaryOp(condition interface{}, v1 interface{}, v2 interface{}) *Traversal {

	if _, ok := condition.(string); ok {
		condition = "'" + condition.(string) + "'"
	}

	if _, ok := v1.(string); ok {
		v1 = "'" + v1.(string) + "'"
	}

	if _, ok := v2.(string); ok {
		v2 = "'" + v2.(string) + "'"
	}

	t.value = t.value + fmt.Sprintf("%v ? %v : %v", condition, v1, v2)
	return t
}

// Raw function is a helper function that can used in case a step is not supported
// or not available in this package. the string is appended as is without modification
func (t *Traversal) Raw(s string) *Traversal {
	t.value = t.value + s
	return t
}

// Append function adds the provided string to the end of the traversal query.
func (t *Traversal) Append(i interface{}) *Traversal {
	t.value = t.value + fmt.Sprintf("%v", i)
	return t
}

// AddLine function adds a new line with the provided string to the traversal query.
func (t *Traversal) AddLine(i interface{}) *Traversal {
	t.value = t.value + "\n" + fmt.Sprintf("%v.", i)
	return t
}
