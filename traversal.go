package gizmo

// Vars type represents a traversal variable
type Vars string

// Traversal struct
type Traversal interface {
	AddV(...interface{}) Traversal
	AddE(...interface{}) Traversal
	Aggregate(...interface{}) Traversal
	And(...interface{}) Traversal
	As(...interface{}) Traversal
	// Between(...interface{}) Traversal
	Both(...interface{}) Traversal
	BothE(...interface{}) Traversal
	BothV() Traversal
	Branch(...interface{}) Traversal
	By(...interface{}) Traversal
	Choose(...interface{}) Traversal
	Coalesce(...interface{}) Traversal
	Create() Traversal
	Count(...interface{}) Traversal
	// CrLf() Traversal
	Dedup() Traversal
	Drop() Traversal
	Emit(...interface{}) Traversal
	Eq(...interface{}) Traversal
	Exists() Traversal
	FlatMap(...interface{}) Traversal
	Filter(...interface{}) Traversal
	Fold() Traversal
	From(...interface{}) Traversal
	getLastStep() *step  // TODO: remove
	getFirstStep() *step // TODO: remove
	Graphs() Traversal
	Gt(...interface{}) Traversal
	Gte(...interface{}) Traversal
	Group() Traversal
	GroupCount() Traversal
	Has(...interface{}) Traversal
	HasLabel(...interface{}) Traversal
	HasNext() Traversal
	In(...interface{}) Traversal
	InE(...interface{}) Traversal
	Inside(...interface{}) Traversal
	Is(...interface{}) Traversal
	InV() Traversal
	Iterate() Traversal
	Label() Traversal
	Limit(...interface{}) Traversal
	Local(...interface{}) Traversal
	Lt(...interface{}) Traversal
	Lte(...interface{}) Traversal
	Map(...interface{}) Traversal
	Match(...interface{}) Traversal
	Neg(...interface{}) Traversal
	Next(...interface{}) Traversal
	// New(...string) Traversal
	Of() Traversal
	Option(...interface{}) Traversal
	Or(...interface{}) Traversal
	OtherV() Traversal
	Order() Traversal
	Out(...interface{}) Traversal
	OutE(...interface{}) Traversal
	Outside(...interface{}) Traversal
	OutV() Traversal
	Path() Traversal
	Project(...interface{}) Traversal
	Property(...interface{}) Traversal
	Properties(...interface{}) Traversal
	Read() Traversal
	String() string
	Select(...interface{}) Traversal
	Set(...interface{}) Traversal
	// TernaryOp(interface{}, interface{}, interface{}) Traversal
	Times(...interface{}) Traversal
	To(...interface{}) Traversal
	ToBulkSet() Traversal
	ToList() Traversal
	ToSet() Traversal
	Unfold() Traversal
	Until(...interface{}) Traversal
	V(...interface{}) Traversal
	ValueMap(...interface{}) Traversal
	Values(...interface{}) Traversal
	Where(...interface{}) Traversal
	Within(...interface{}) Traversal
	Without(...interface{}) Traversal
}

// G returns a Step and starts the query building process
func G(n ...string) Traversal {
	return NewGraph(n...)
}

// NewGraph returns a Step and starts the query building process
func NewGraph(n ...string) Traversal {
	var s step
	if len(n) > 0 {
		s = step{n[0], nil, "", "", &s, nil, &s}
	} else {
		s = step{"g", nil, "", "", &s, nil, &s}
	}
	return &s
}

// VAR is used to start a Traversal with a variable in the form v=
func VAR(n string) Traversal {
	return startStep(n, "=", "", nil)
}

// VAL is used to start a Traversal with a variable in the form v=
func VAL(n string) Traversal {
	return startStep(n, "", "", nil)
}

// Append adds step 's' to the end of step 'd'
// TODO: remove
func Append(d Traversal, s Traversal) {
	d.getLastStep().next = s.getFirstStep()
	d.getFirstStep().tail = s.getFirstStep()
}

func appendStep(s *step, name string, open string, close string, prev *step, next *step, params ...interface{}) Traversal {
	st := &step{name, params, open, close, prev, next, nil}
	st.head = s.head
	if s.head == s {
		s.tail.next = st
	} else {
		s.next = st
	}
	s.head.tail = st
	return st
}

func startStep(name string, open string, close string, params ...interface{}) Traversal {
	st := &step{name, params, open, close, nil, nil, nil}
	st.head = st
	st.tail = st
	return st
}

// TODO: remove
func (s *step) getFirstStep() *step {
	return s.head
}

// TODO: remove
func (s *step) getLastStep() *step {
	return s.tail
}

// V function
func (s *step) V(v ...interface{}) Traversal {
	return appendStep(s, "V", "(", ")", nil, nil, v...)
}

// AddE is a map/sideEffect. An edge is added from a Traversal g using addE between two existing vertices.
// A previously created edge label must be specified.
func (s *step) AddE(v ...interface{}) Traversal {
	return appendStep(s, "addE", "(", ")", nil, nil, v...)
}

// AddV step is a map/sideEffect. A vertex is added from a Traversal g using addV.
// A previously created vertex label must be specified. Property key-value pairs may be optionally specified.
func (s *step) AddV(v ...interface{}) Traversal {
	return appendStep(s, "addV", "(", ")", nil, nil, v...)
}

// Path A traverser is transformed as it moves through a series of steps within a traversal.
// The history of the traverser is realized by examining its path with path()-step (map).
func (s *step) Path() Traversal {
	return appendStep(s, "project", "(", ")", nil, nil, nil)
}

// Project step is a map step that projects the current object into a map keyed by provided labels.
// It is similar to select() step.
func (s *step) Project(v ...interface{}) Traversal {
	return appendStep(s, "project", "(", ")", nil, nil, v...)
}

// Select step is a map step that selects labeled steps designated with as() steps.
// This step is typically used to select particular properties from objects earlier in the traversal.
func (s *step) Select(v ...interface{}) Traversal {
	return appendStep(s, "select", "(", ")", nil, nil, v...)
}

// Where returns a new gremlin Traversal
func (s *step) Where(v ...interface{}) Traversal {
	return appendStep(s, "where", "(", ")", nil, nil, v...)
}

// And step can take an arbitrary number of Traversals to filter the returned objects using a boolean AND function.
// The related boolean OR function can be done with the or() step.
func (s *step) And(v ...interface{}) Traversal {
	return appendStep(s, "and", "(", ")", nil, nil, v...)
}

// FlatMap step is a general step. This step creates a map of the Traversal objects and streams the map to the next Traversal step.
func (s *step) FlatMap(v ...interface{}) Traversal {
	return appendStep(s, "flatMap", "(", ")", nil, nil, v...)
}

// Filter step is a general step. Map the traverser to either true or false, where false will not pass the traverser to the next step.
func (s *step) Filter(v ...interface{}) Traversal {
	return appendStep(s, "filter", "{", "}", nil, nil, v...)
}

// Branch step is a general step. Split the traverser to all the traversals indexed by the M token.
func (s *step) Branch(v ...interface{}) Traversal {
	return appendStep(s, "branch", "(", ")", nil, nil, v...)
}

// Choose step is a branch step that can be used to route a traverser to a particular traversal branch, similar to if-then-else logic.
func (s *step) Choose(v ...interface{}) Traversal {
	return appendStep(s, "choose", "(", ")", nil, nil, v...)
}

// Coalesce step evaluates the provided traversals in order and returns the first traversal that emits at least one element..
func (s *step) Coalesce(v ...interface{}) Traversal {
	return appendStep(s, "coalesce", "(", ")", nil, nil, v...)
}

// SideEffect step is a general step. Perform some operation on the traverser and pass it to the next step.
// func (s *step) SideEffect(v ...interface{}) Traversal {
// 	return appendStep(s, "sideEffect", "(", ")", nil, nil, v...)
// }

// Property step in a traversal has the same beahviour as in the Graph API
func (s *step) Property(v ...interface{}) Traversal {
	return appendStep(s, "property", "(", ")", nil, nil, v...)
}

// Properties step retrieves the properties of the specified element.
func (s *step) Properties(v ...interface{}) Traversal {
	return appendStep(s, "properties", "(", ")", nil, nil, v...)
}

// Limit returns a new gremlin Traversal
func (s *step) Limit(v ...interface{}) Traversal {
	return appendStep(s, "limit", "(", ")", nil, nil, v...)
}

// Aggregate step (sideEffect) is used to aggregate all the objects at a particular point of traversal
// into a Collection. The step uses eager evaluation in that no objects continue on until
// all previous objects have been fully aggregated (as opposed to store() which lazily fills a collection).
// The eager evaluation nature is crucial in situations where everything at a particular point is required for future computation.
func (s *step) Aggregate(v ...interface{}) Traversal {
	return appendStep(s, "aggregate", "(", ")", nil, nil, v...)
}

// ValueMap The valueMap() step yields a map representation of the properties of an element.
func (s *step) ValueMap(v ...interface{}) Traversal {
	return appendStep(s, "valueMap", "(", ")", nil, nil, v...)
}

// Order step is a map step that sorts returned objects given a certain criteria.
func (s *step) Order() Traversal {
	return appendStep(s, "order", "(", ")", nil, nil, nil)
}

// Value The valueMap() step yields a map representation of the properties of an element.
func (s *step) Value() Traversal {
	return appendStep(s, "value", "(", ")", nil, nil, nil)
}

// Values step is a map step that extracts the values of either all the properties or specified properties for an element.
func (s *step) Values(v ...interface{}) Traversal {
	return appendStep(s, "values", "(", ")", nil, nil, v...)
}

// Label step is a map step that extracts the specified labels.
func (s *step) Label() Traversal {
	return appendStep(s, "value", "(", ")", nil, nil, nil)
}

// Local step is a branch step that allows the action within the step to be applied to the element within the local() step.
func (s *step) Local(v ...interface{}) Traversal {
	return appendStep(s, "local", "(", ")", nil, nil, v...)
}

// Dedup returns a new gremlin Traversal
func (s *step) Dedup() Traversal {
	return appendStep(s, "dedup", "(", ")", nil, nil, nil)
}

// Read returns a new gremlin Traversal
func (s *step) Read() Traversal {
	return appendStep(s, "read", "(", ")", nil, nil, nil)
}

// Range step is a filter step that allows only the specified number of objects to pass onto the next step.
func (s *step) Range(v ...interface{}) Traversal {
	return appendStep(s, "range", "(", ")", nil, nil, v...)
}

// Count returns a new gremlin Traversal
func (s *step) Count(v ...interface{}) Traversal {
	return appendStep(s, "count", "(", ")", nil, nil, v...)
}

// Of returns a new gremlin Traversal
func (s *step) Of() Traversal {
	return appendStep(s, "of", "(", ")", nil, nil, nil)
}

// Fold There are situations when the traversal stream needs a "barrier" to aggregate all the objects
// and emit a computation that is a function of the aggregate. The fold()-step (map) is one particular instance of this.
// Please see unfold()-step for the inverse functionality.
func (s *step) Fold() Traversal {
	return appendStep(s, "fold", "(", ")", nil, nil, nil)
}

// Unfold If the object reaching unfold() (flatMap) is an iterator, iterable, or map, then it is unrolled into a linear form.
// If not, then the object is simply emitted. Please see fold() step for the inverse behavior.
func (s *step) Unfold() Traversal {
	return appendStep(s, "unfold", "(", ")", nil, nil, nil)
}

// GroupCount returns a new gremlin Traversal
func (s *step) GroupCount() Traversal {
	return appendStep(s, "groupCount", "(", ")", nil, nil, nil)
}

// Group returns a new gremlin Traversal
func (s *step) Group() Traversal {
	return appendStep(s, "group", "(", ")", nil, nil, nil)
}

////////////////
// Vertex Steps
///////////////

// Out step moves the Traversal to the outgoing adjacent vertices, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (s *step) Out(v ...interface{}) Traversal {
	return appendStep(s, "out", "(", ")", nil, nil, v...)
}

// Has step is a filter step. It is the most common step used for graph Traversals,
// since this step narrows the query to find particular vertices or edges with certain property values.
func (s *step) Has(v ...interface{}) Traversal {
	return appendStep(s, "has", "(", ")", nil, nil, v...)
}

// HasLabel step is a filter step.
func (s *step) HasLabel(v ...interface{}) Traversal {
	return appendStep(s, "hasLabel", "(", ")", nil, nil, v...)
}

// OutE step moves the Traversal to the outgoing incident edges, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (s *step) OutE(v ...interface{}) Traversal {
	return appendStep(s, "outE", "(", ")", nil, nil, v...)
}

// OutV The out() step moves the Traversal to the outgoing adjacent vertices,
// given the edge labels. Specifying no edge label will traverse all incident edges.
func (s *step) OutV() Traversal {
	return appendStep(s, "outV", "(", ")", nil, nil, nil)
}

// In step moves the Traversal to the incoming adjacent vertices, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (s *step) In(v ...interface{}) Traversal {
	return appendStep(s, "in", "(", ")", nil, nil, v...)
}

// InE step moves the Traversal to the incoming incident edges, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (s *step) InE(v ...interface{}) Traversal {
	return appendStep(s, "inE", "(", ")", nil, nil, v...)
}

// InV step moves the Traversal to the incoming vertices.
func (s *step) InV() Traversal {
	return appendStep(s, "inV", "(", ")", nil, nil, nil)
}

// Both step moves the Traversal to both the outgoing and the incoming adjacent vertices, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (s *step) Both(v ...interface{}) Traversal {
	return appendStep(s, "both", "(", ")", nil, nil, v...)
}

// BothE step moves the Traversal to both the outgoing and the incoming incident edges, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func (s *step) BothE(v ...interface{}) Traversal {
	return appendStep(s, "bothE", "(", ")", nil, nil, v...)
}

// BothV step moves the Traversal to both the outgoing and the incoming vertices.
func (s *step) BothV() Traversal {
	return appendStep(s, "bothV", "(", ")", nil, nil, nil)
}

// OtherV step moves the Traversal to the vertex that was not the vertex that was moved from.
func (s *step) OtherV() Traversal {
	return appendStep(s, "otherV", "(", ")", nil, nil, nil)
}

// Map the traverser to some object of type E for the next step to process.
func (s *step) Map(v ...interface{}) Traversal {
	return appendStep(s, "map", "(", ")", nil, nil, v...)
}

// Match step is a map step provides a more declarative form of graph querying based on pattern matching.
func (s *step) Match(v ...interface{}) Traversal {
	return appendStep(s, "match", "(", ")", nil, nil, v...)
}

////////////////
// System
///////////////

// Exists returns a new gremlin Traversal
func (s *step) Exists() Traversal {
	return appendStep(s, "exists", "(", ")", nil, nil, nil)
}

// Create returns a new gremlin Traversal
func (s *step) Create() Traversal {
	return appendStep(s, "create", "(", ")", nil, nil, nil)
}

// Drop returns a new gremlin Traversal
func (s *step) Drop() Traversal {
	return appendStep(s, "drop", "(", ")", nil, nil, nil)
}

// Graphs discover what graphs currently exist using this command.
func (s *step) Graphs() Traversal {
	return appendStep(s, "graphs", "(", ")", nil, nil, nil)
}

// Set returns a new gremlin Traversal
func (s *step) Set(v ...interface{}) Traversal {
	return appendStep(s, "set", "(", ")", nil, nil, v...)
}

////////////////
// Predicates
///////////////

// Eq predicate answers the question: Is the incoming object equal to the provided object?
func (s *step) Eq(v ...interface{}) Traversal {
	return appendStep(nil, "eq", "(", ")", nil, nil, v...)
}

// Neg predicate answers the question: Is the incoming object equal to the provided object?
func (s *step) Neg(v ...interface{}) Traversal {
	return appendStep(nil, "neg", "(", ")", nil, nil, v...)
}

// Lt predicate answers the question: Is the incoming number less than the provided number?
func (s *step) Lt(v ...interface{}) Traversal {
	return appendStep(nil, "lt", "(", ")", nil, nil, v...)
}

// Lte predicate answers the question: Is the incoming number less than or equal to the provided number?
func (s *step) Lte(v ...interface{}) Traversal {
	return appendStep(nil, "lte", "(", ")", nil, nil, v...)
}

// Gt  predicate answers the question: Is the incoming number greater than the provided number?
func (s *step) Gt(v ...interface{}) Traversal {
	return appendStep(nil, "gt", "(", ")", nil, nil, v...)
}

// Gte predicate answers the question: Is the incoming number greater than or equal to the provided number?
func (s *step) Gte(v ...interface{}) Traversal {
	return appendStep(nil, "gte", "(", ")", nil, nil, v...)
}

// Inside predicate answers the question: Is the incoming number greater than the first provided number and less than the second?
func (s *step) Inside(v ...interface{}) Traversal {
	return appendStep(nil, "inside", "(", ")", nil, nil, v...)
}

// Is, it's possible to filter scalar values using is()-step (filter).
func (s *step) Is(v ...interface{}) Traversal {
	return appendStep(nil, "is", "(", ")", nil, nil, v...)
}

// Outside predicate answers the question: Is the incoming number less than the first provided number or greater than the second?
func (s *step) Outside(v ...interface{}) Traversal {
	return appendStep(nil, "outside", "(", ")", nil, nil, v...)
}

// Between predicate answers the question: Is the incoming number less than the first provided number or greater than the second?
func (s *step) Between(v ...interface{}) Traversal {
	return appendStep(nil, "between", "(", ")", nil, nil, v...)
}

// Within predicate answers the question: Is the incoming object in the array of provided objects?
func (s *step) Within(v ...interface{}) Traversal {
	return appendStep(s, "within", "(", ")", nil, nil, v...)
}

// Without predicate answers the question: Is the incoming object not in the array of provided objects?
func (s *step) Without(v ...interface{}) Traversal {
	return appendStep(s, "without", "(", ")", nil, nil, v...)
}

////////////////
// Step-modulators
///////////////

// As step is a step modulator, a helper step for another Traversal step.
func (s *step) As(v ...interface{}) Traversal {
	return appendStep(s, "as", "(", ")", nil, nil, v...)
}

// By step is a step modulator that can be used to modify sort order from a previous step.
func (s *step) By(v ...interface{}) Traversal {
	return appendStep(s, "by", "(", ")", nil, nil, v...)
}

// From step is a step modulator, a helper step for another Traversal step.
// Its main use is to designate the outgoing vertex for an addE() step. Generally, a to() step is paired with a from() step.
func (s *step) From(v ...interface{}) Traversal {
	return appendStep(s, "from", "(", ")", nil, nil, v...)
}

// Option step modulator is used in conjunction with branch()
// or choose(), and specifies the returned values based on values found at that point in a Traversal.
func (s *step) Option(v ...interface{}) Traversal {
	return appendStep(s, "or", "(", ")", nil, nil, v...)
}

// Or step ensures that at least one of the provided traversals yield a result (filter). Please see and() for and-semantics.
func (s *step) Or(v ...interface{}) Traversal {
	return appendStep(s, "or", "(", ")", nil, nil, v...)
}

// Times step is a step modulator, a helper step for another Traversal step.
// Its main use is to repeat a repeat() step for the specified number of times.
func (s *step) Times(v ...interface{}) Traversal {
	return appendStep(s, "times", "(", ")", nil, nil, v...)
}

// To returns a new gremlin Traversal
func (s *step) To(v ...interface{}) Traversal {
	return appendStep(s, "to", "(", ")", nil, nil, v...)
}

// Emit step is a step modulator, a helper step for another Traversal step.
// Its main use is to emit either incoming Traversals before a repeat() step, or emit outgoing Traversals after a repeat() step.
// The emission sends a copy of the current objects to the next step in the query.
// A predicate or Traversal can be used in an emit() step to cause the emission only if the predicate or Traversal is true.
func (s *step) Emit(v ...interface{}) Traversal {
	return appendStep(s, "emit", "(", ")", nil, nil, v...)
}

// Until step is a step modulator, a helper step for another Traversal step.
// Its main use is to turn a repeat() step into a do-while loop (if used after the repeat() step) or a while-do loop (if used before the repeat() step).
// A predicate or Traversal can be used in an until() step to cause the loop to complete only if the predicate or Traversal is true.
func (s *step) Until(v ...interface{}) Traversal {
	return appendStep(s, "until", "(", ")", nil, nil, v...)
}

////////////////////////
// Terminal steps
////////////////////////

// Next  step is a terminal step. It returns the next number of steps, based on a supplied integer value.
func (s *step) Next(v ...interface{}) Traversal {
	return appendStep(s, "next", "(", ")", nil, nil, v...)
}

// HasNext step is a terminal step. It determines whether or not there are available results from a traversal,
// returning a Boolean value of true or false.
func (s *step) HasNext() Traversal {
	return appendStep(s, "hasNext", "(", ")", nil, nil, nil)
}

// ToList will return all results in a list.
func (s *step) ToList() Traversal {
	return appendStep(s, "toList", "(", ")", nil, nil, nil)
}

// ToSet will return all results in a set and thus, duplicates removed.
func (s *step) ToSet() Traversal {
	return appendStep(s, "toSet", "(", ")", nil, nil, nil)
}

// ToBulkSet will return all results in a weighted set and thus, duplicates preserved via weighting
func (s *step) ToBulkSet() Traversal {
	return appendStep(s, "toBulkSet", "(", ")", nil, nil, nil)
}

// Iterate does not exactly fit the definition of a terminal step in that it doesn’t return a result,
// but still returns a traversal - it does however behave as a terminal step in that it iterates
// the traversal and generates side effects without returning the actual result.
func (s *step) Iterate() Traversal {
	return appendStep(s, "iterate", "(", ")", nil, nil, nil)
}
