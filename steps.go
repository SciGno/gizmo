package gizmo

// // V function
// func V(name ...interface{}) Traversal {
// 	t := NewTraversal("")
// 	return t.V(name...)
// }

// // AddV step is a map/sideEffect. A vertex is added from a Traversal g using addV.
// // A previously created vertex label must be specified. Property key-value pairs may be optionally specified.
// func AddV(name ...string) Traversal {
// 	t := NewTraversal("")
// 	return t.AddV(name...)
// }

// // Project step is a map step that projects the current object into a map keyed by provided labels.
// // It is similar to select() step.
// func Project(name ...string) Traversal {
// 	t := NewTraversal("")
// 	return t.Project(name...)
// }

// // Select step is a map step that selects labeled steps designated with as() steps.
// // This step is typically used to select particular properties from objects earlier in the traversal.
// func Select(name ...string) Traversal {
// 	t := NewTraversal("")
// 	return t.Select(name...)
// }

// // Where returns a new gremlin Traversal
// func Where(i interface{}) Traversal {
// 	t := NewTraversal("")
// 	return t.Where(i)
// }

// // And step can take an arbitrary number of Traversals to filter the returned objects using a boolean AND function.
// // The related boolean OR function can be done with the or() step.
// func And(i interface{}) Traversal {
// 	t := NewTraversal("")
// 	return t.And(i)
// }

// // FlatMap step is a general step. This step creates a map of the Traversal objects and streams the map to the next Traversal step.
// func FlatMap(i interface{}) Traversal {
// 	t := NewTraversal("")
// 	return t.FlatMap(i)
// }

// // Filter step is a general step. Map the traverser to either true or false, where false will not pass the traverser to the next step.
// func Filter(i interface{}) Traversal {
// 	t := NewTraversal("")
// 	return t.Filter(i)
// }

// // AddE is a map/sideEffect. An edge is added from a Traversal g using addE between two existing vertices.
// // A previously created edge label must be specified.
// func AddE(s string) Traversal {
// 	t := NewTraversal("")
// 	return t.AddE(s)
// }

// // Property step in a traversal has the same beahviour as in the Graph API
// func Property(k interface{}, v interface{}) Traversal {
// 	t := NewTraversal("")
// 	return t.Property(k, v)
// }

// // Properties step retrieves the properties of the specified element.
// func Properties(name ...string) Traversal {
// 	t := NewTraversal("")
// 	return t.Properties(name...)
// }

// // Limit returns a new gremlin Traversal
// func Limit(i int) Traversal {
// 	t := NewTraversal("")
// 	return t.Limit(i)
// }

// // Aggregate step (sideEffect) is used to aggregate all the objects at a particular point of traversal
// // into a Collection. The step uses eager evaluation in that no objects continue on until
// // all previous objects have been fully aggregated (as opposed to store() which lazily fills a collection).
// // The eager evaluation nature is crucial in situations where everything at a particular point is required for future computation.
// func Aggregate(s string) Traversal {
// 	t := NewTraversal("")
// 	return t.Aggregate(s)
// }

// // ValueMap The valueMap() step yields a map representation of the properties of an element.
// func ValueMap() Traversal {
// 	t := NewTraversal("")
// 	return t.ValueMap()
// }

// // Order step is a map step that sorts returned objects given a certain criteria.
// func Order() Traversal {
// 	t := NewTraversal("")
// 	return t.Order()
// }

// // Label step is a map step that extracts the specified labels.
// func Label() Traversal {
// 	t := NewTraversal("")
// 	return t.Label()
// }

// // Dedup returns a new gremlin Traversal
// func Dedup() Traversal {
// 	t := NewTraversal("")
// 	return t.Dedup()
// }

// // Read returns a new gremlin Traversal
// func Read() Traversal {
// 	t := NewTraversal("")
// 	return t.Read()
// }

// // Of returns a new gremlin Traversal
// func Of() Traversal {
// 	t := NewTraversal("")
// 	return t.Of()
// }

// // Fold There are situations when the traversal stream needs a "barrier" to aggregate all the objects
// // and emit a computation that is a function of the aggregate. The fold()-step (map) is one particular instance of this.
// // Please see unfold()-step for the inverse functionality.
// func Fold() Traversal {
// 	t := NewTraversal("")
// 	return t.Fold()
// }

// // Unfold If the object reaching unfold() (flatMap) is an iterator, iterable, or map, then it is unrolled into a linear form.
// // If not, then the object is simply emitted. Please see fold() step for the inverse behavior.
// func Unfold() Traversal {
// 	t := NewTraversal("")
// 	return t.Unfold()
// }

// // GroupCount returns a new gremlin Traversal
// func GroupCount() Traversal {
// 	t := NewTraversal("")
// 	return t.GroupCount()
// }

// // Group returns a new gremlin Traversal
// func Group() Traversal {
// 	t := NewTraversal("")
// 	return t.Group()
// }

// ////////////////
// // Vertex Steps
// ///////////////

// // Out step moves the Traversal to the outgoing adjacent vertices, given the edge labels.
// // Specifying no edge label will traverse all incident edges.
// func Out(name ...string) Traversal {
// 	t := NewTraversal("")
// 	return t.Out(name...)
// }

// // Has step is a filter step. It is the most common step used for graph Traversals,
// // since this step narrows the query to find particular vertices or edges with certain property values.
// func Has(name string, values ...interface{}) Traversal {
// 	t := NewTraversal("")
// 	return t.Has(name, values...)
// }

// // HasLabel step is a filter step.
// func HasLabel(values ...string) Traversal {
// 	t := NewTraversal("")
// 	return t.HasLabel(values...)
// }

// // OutE step moves the Traversal to the outgoing incident edges, given the edge labels.
// // Specifying no edge label will traverse all incident edges.
// func OutE(name ...string) Traversal {
// 	t := NewTraversal("")
// 	return t.OutE(name...)
// }

// // OutV The out() step moves the Traversal to the outgoing adjacent vertices,
// // given the edge labels. Specifying no edge label will traverse all incident edges.
// func OutV() Traversal {
// 	t := NewTraversal("")
// 	return t.OutV()
// }

// // In step moves the Traversal to the incoming adjacent vertices, given the edge labels.
// // Specifying no edge label will traverse all incident edges.
// func In(name ...string) Traversal {
// 	t := NewTraversal("")
// 	return t.In(name...)
// }

// // InE step moves the Traversal to the incoming incident edges, given the edge labels.
// // Specifying no edge label will traverse all incident edges.
// func InE(name ...string) Traversal {
// 	t := NewTraversal("")
// 	return t.InE(name...)
// }

// // InV step moves the Traversal to the incoming vertices.
// func InV() Traversal {
// 	t := NewTraversal("")
// 	return t.InV()
// }

// // Both step moves the Traversal to both the outgoing and the incoming adjacent vertices, given the edge labels.
// // Specifying no edge label will traverse all incident edges.
// func Both(name ...string) Traversal {
// 	t := NewTraversal("")
// 	return t.Both(name...)
// }

// // BothE step moves the Traversal to both the outgoing and the incoming incident edges, given the edge labels.
// // Specifying no edge label will traverse all incident edges.
// func BothE(name ...string) Traversal {
// 	t := NewTraversal("")
// 	return t.BothE(name...)
// }

// // BothV step moves the Traversal to both the outgoing and the incoming vertices.
// func BothV() Traversal {
// 	t := NewTraversal("")
// 	return t.BothV()
// }

// // OtherV step moves the Traversal to the vertex that was not the vertex that was moved from.
// func OtherV() Traversal {
// 	t := NewTraversal("")
// 	return t.OtherV()
// }

// // Match step is a map step provides a more declarative form of graph querying based on pattern matching.
// func Match(i ...interface{}) Traversal {
// 	t := NewTraversal("")
// 	return t.Match(i...)
// }

// ////////////////
// // System
// ///////////////

// // Exists returns a new gremlin Traversal
// func Exists() Traversal {
// 	t := NewTraversal("")
// 	return t.Exists()
// }

// // Create returns a new gremlin Traversal
// func Create() Traversal {
// 	t := NewTraversal("")
// 	return t.Create()
// }

// // Drop returns a new gremlin Traversal
// func Drop() Traversal {
// 	t := NewTraversal("")
// 	return t.Drop()
// }

// // Graphs discover what graphs currently exist using this command.
// func Graphs() Traversal {
// 	t := NewTraversal("")
// 	return t.Graphs()
// }

// // Set returns a new gremlin Traversal
// func Set(s string) Traversal {
// 	t := NewTraversal("")
// 	return t.Set(s)
// }

// // SideEffect step is a general step. Perform some operation on the traverser and pass it to the next step.
// func SideEffect(i interface{}) Traversal {
// 	t := NewTraversal("")
// 	return t.SideEffect(i)
// }

// // Branch step is a general step. Split the traverser to all the traversals indexed by the M token.
// func Branch(i interface{}) Traversal {
// 	t := NewTraversal("")
// 	return t.Branch(i)
// }

// // Choose step is a general step.
// func Choose(i ...interface{}) Traversal {
// 	t := NewTraversal("")
// 	return t.Choose(i...)
// }

// // Coalesce step evaluates the provided traversals in order and returns the first traversal that emits at least one element..
// func Coalesce(i ...interface{}) Traversal {
// 	t := NewTraversal("")
// 	return t.Coalesce(i...)
// }

// ////////////////
// // Predicates
// ///////////////

// // Eq predicate answers the question: Is the incoming object equal to the provided object?
// func Eq(i interface{}) Traversal {
// 	t := NewTraversal("")
// 	return t.Eq(i)
// }

// // Neg predicate answers the question: Is the incoming object equal to the provided object?
// func Neg(i interface{}) Traversal {
// 	t := NewTraversal("")
// 	return t.Neg(i)
// }

// // Lt predicate answers the question: Is the incoming number less than the provided number?
// func Lt(i int) Traversal {
// 	t := NewTraversal("")
// 	return t.Lt(i)
// }

// // Lte predicate answers the question: Is the incoming number less than or equal to the provided number?
// func Lte(i int) Traversal {
// 	t := NewTraversal("")
// 	return t.Lte(i)
// }

// // Gt  predicate answers the question: Is the incoming number greater than the provided number?
// func Gt(i int) Traversal {
// 	t := NewTraversal("")
// 	return t.Gt(i)
// }

// // Gte predicate answers the question: Is the incoming number greater than or equal to the provided number?
// func Gte(i int) Traversal {
// 	t := NewTraversal("")
// 	return t.Gte(i)
// }

// // Inside predicate answers the question: Is the incoming number greater than the first provided number and less than the second?
// func Inside(from int, to int) Traversal {
// 	t := NewTraversal("")
// 	return t.Inside(from, to)
// }

// // Outside predicate answers the question: Is the incoming number less than the first provided number or greater than the second?
// func Outside(from int, to int) Traversal {
// 	t := NewTraversal("")
// 	return t.Outside(from, to)
// }

// // Between predicate answers the question: Is the incoming number less than the first provided number or greater than the second?
// func Between(from int, to int) Traversal {
// 	t := NewTraversal("")
// 	return t.Between(from, to)
// }

// // Within predicate answers the question: Is the incoming object in the array of provided objects?
// func Within(name ...string) Traversal {
// 	t := NewTraversal("")
// 	return t.Within(name...)
// }

// // Without predicate answers the question: Is the incoming object not in the array of provided objects?
// func Without(name ...string) Traversal {
// 	t := NewTraversal("")
// 	return t.Without(name...)
// }

// ////////////////
// // Step-modulators
// ///////////////

// // As step is a step modulator, a helper step for another Traversal step.
// func As(s string) Traversal {
// 	t := NewTraversal("")
// 	return t.As(s)
// }

// // By step is a step modulator that can be used to modify sort order from a previous step.
// func By(i interface{}) Traversal {
// 	t := NewTraversal("")
// 	return t.By(i)
// }

// // From step is a step modulator, a helper step for another Traversal step.
// // Its main use is to designate the outgoing vertex for an addE() step. Generally, a to() step is paired with a from() step.
// func From(i interface{}) Traversal {
// 	t := NewTraversal("")
// 	return t.From(i)
// }

// // Option step modulator is used in conjunction with branch()
// // or choose(), and specifies the returned values based on values found at that point in a Traversal.
// func Option(s string) Traversal {
// 	t := NewTraversal("")
// 	return t.Option(s)
// }

// // Times step is a step modulator, a helper step for another Traversal step.
// // Its main use is to repeat a repeat() step for the specified number of times.
// func Times(n int) Traversal {
// 	t := NewTraversal("")
// 	return t.Times(n)
// }

// // To returns a new gremlin Traversal
// func To(s interface{}) Traversal {
// 	t := NewTraversal("")
// 	return t.To(s)
// }

// // Emit step is a step modulator, a helper step for another Traversal step.
// // Its main use is to emit either incoming Traversals before a repeat() step, or emit outgoing Traversals after a repeat() step.
// // The emission sends a copy of the current objects to the next step in the query.
// // A predicate or Traversal can be used in an emit() step to cause the emission only if the predicate or Traversal is true.
// func Emit(i ...interface{}) Traversal {
// 	t := NewTraversal("")
// 	return t.Emit(i...)
// }

// // Until step is a step modulator, a helper step for another Traversal step.
// // Its main use is to turn a repeat() step into a do-while loop (if used after the repeat() step) or a while-do loop (if used before the repeat() step).
// // A predicate or Traversal can be used in an until() step to cause the loop to complete only if the predicate or Traversal is true.
// func Until(i ...interface{}) Traversal {
// 	t := NewTraversal("")
// 	return t.Until(i...)
// }

// // Values step is a map step that extracts the values of either all the properties or specified properties for an element.
// func Values(s ...string) Traversal {
// 	t := NewTraversal("")
// 	return t.Values(s...)
// }

// //////////////////
// // Other Helpers
// /////////////////

// // TernaryOp step is a sideEffect.
// func TernaryOp(condition interface{}, v1 interface{}, v2 interface{}) Traversal {
// 	t := NewTraversal("")
// 	return t.TernaryOp(condition, v1, v2)
// }

// // Raw function is a helper function that can used in case a step is not supported
// // or not available in this package. The string is appended to the end as is without modification
// func Raw(s string) Traversal {
// 	t := NewTraversal("")
// 	return t.Raw(s)
// }

// // Append function adds the provided string to the end of the traversal query.
// func Append(i interface{}) Traversal {
// 	t := NewTraversal("")
// 	return t.Append(i)
// }

// // AddLine function adds a new line with the provided string to the traversal query.
// // func AddLine(i interface{}) Traversal {
// // 	t := NewTraversal("")
// // 	return t.AddLine(i)
// // }

// ////////////////////////
// // Terminal steps
// ////////////////////////

// // Next  step is a terminal step. It returns the next number of steps, based on a supplied integer value.
// func Next(i ...int) Traversal {
// 	t := NewTraversal("")
// 	return t.Next(i...)
// }

// // HasNext step is a terminal step. It determines whether or not there are available results from a traversal,
// // returning a Boolean value of true or false.
// func HasNext() Traversal {
// 	t := NewTraversal("")
// 	return t.HasNext()
// }

// // ToList will return all results in a list.
// func ToList() Traversal {
// 	t := NewTraversal("")
// 	return t.ToList()
// }

// // ToSet will return all results in a set and thus, duplicates removed.
// func ToSet() Traversal {
// 	t := NewTraversal("")
// 	return t.ToSet()
// }

// // ToBulkSet will return all results in a weighted set and thus, duplicates preserved via weighting
// func ToBulkSet() Traversal {
// 	t := NewTraversal("")
// 	return t.ToBulkSet()
// }

// // Iterate does not exactly fit the definition of a terminal step in that it doesn’t return a result,
// // but still returns a traversal - it does however behave as a terminal step in that it iterates
// // the traversal and generates side effects without returning the actual result.
// func Iterate() Traversal {
// 	t := NewTraversal("")
// 	return t.Iterate()
// }
