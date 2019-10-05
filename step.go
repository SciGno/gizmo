package gizmo

import (
	"fmt"
	"strconv"
	"strings"
)

// Step - implements Traversal
type step struct {
	name   string
	params []interface{}
	open   string
	close  string
	head   *step
	next   *step
	tail   *step
}

func (s *step) String() string {

	// this is the complete representation of the Step
	var myValue string
	// represent the list of parameters in their string output
	list := []string{}
	// logger.Info("Processing: ", s.name)
	// logger.Info("Number of params: ", len(s.params))

	for _, v := range s.params {
		// logger.Info("Processing param: ", k)
		// logger.Infof("Param type: %T", v)
		switch v.(type) {
		case []string:
			for _, i := range v.([]string) {
				list = append(list, "'"+i+"'")
			}
		case []int:
			for _, i := range v.([]int) {
				list = append(list, strconv.Itoa(i))
			}
		case []float32:
			for _, i := range v.([]float32) {
				list = append(list, fmt.Sprintf("%g", i))
			}
		case []float64:
			for _, i := range v.([]float64) {
				list = append(list, fmt.Sprintf("%g", i))
			}
		case []interface{}:
			// logger.Info("Type is []interface{}...")
			// logger.Info("Interface: ", v)
			for _, i := range v.([]interface{}) {
				// logger.Info("Interface value: ", i)
				switch i.(type) {
				case string:
					list = append(list, "'"+i.(string)+"'")
				default:
					list = append(list, fmt.Sprintf("%v", i))
				}
			}
		case string:
			// logger.Info("Processing String parameter")
			list = append(list, "'"+v.(string)+"'")
		case *step:
			// logger.Info("Processing Step parameter")
			t := v.(*step)
			// logger.Info(s.name+"'s parameter name: ", t.name)
			list = append(list, t.head.String())
		case nil:
			// logger.Info("Processing Nil parameter")
		default:
			// logger.Info("Processing default...")
			// logger.Infof("Param type: %T", v)
			list = append(list, fmt.Sprintf("%v", v))
		}
	}

	myValue = s.name + s.open + strings.Join(list, ", ") + s.close
	if s.next != nil {
		// logger.Info(s.name + " Step has next...")
		// if s.open == "v" {
		// 	myValue = myValue + s.next.String()
		// }
		if s.open == "=" {
			myValue = myValue + s.next.String()
		} else {
			myValue = myValue + "." + s.next.String()
		}
	}
	return myValue
}

// V function
func V(v ...interface{}) Traversal {
	return startStep("V", "(", ")", v...)
}

// AddE is a map/sideEffect. An edge is added from a Traversal g using addE between two existing vertices.
// A previously created edge label must be specified.
func AddE(v ...interface{}) Traversal {
	return startStep("addE", "(", ")", v...)
}

// AddV step is a map/sideEffect. A vertex is added from a Traversal g using addV.
// A previously created vertex label must be specified. Property key-value pairs may be optionally specified.
func AddV(v ...interface{}) Traversal {
	return startStep("addV", "(", ")", v...)
}

// Path A traverser is transformed as it moves through a series of steps within a traversal.
// The history of the traverser is realized by examining its path with path()-step (map).
func Path() Traversal {
	return startStep("project", "(", ")", nil)
}

// Project step is a map step that projects the current object into a map keyed by provided labels.
// It is similar to select() step.
func Project(v ...interface{}) Traversal {
	return startStep("project", "(", ")", v...)
}

// Select step is a map step that selects labeled steps designated with as() steps.
// This step is typically used to select particular properties from objects earlier in the traversal.
func Select(v ...interface{}) Traversal {
	return startStep("select", "(", ")", v...)
}

// Where returns a new gremlin Traversal
func Where(v ...interface{}) Traversal {
	return startStep("where", "(", ")", v...)
}

// And step can take an arbitrary number of Traversals to filter the returned objects using a boolean AND function.
// The related boolean OR function can be done with the or() step.
func And(v ...interface{}) Traversal {
	return startStep("and", "(", ")", v...)
}

// GroupCount returns a new gremlin Traversal
func GroupCount() Traversal {
	return startStep("groupCount", "(", ")")
}

// Group returns a new gremlin Traversal
func Group() Traversal {
	return startStep("group", "(", ")")
}

// Unfold function
func Unfold(v ...interface{}) Traversal {
	return startStep("unfold", "(", ")", v...)
}

// Values step is a map step that extracts the values of either all the properties or specified properties for an element.
func Values(v ...interface{}) Traversal {
	return startStep("values", "(", ")", v...)
}

// Label step is a map step that extracts the specified labels.
func Label() Traversal {
	return startStep("value", "(", ")", nil)
}

// Branch step is a general step. Split the traverser to all the traversals indexed by the M token.
func Branch(v ...interface{}) Traversal {
	return startStep("branch", "(", ")", v...)
}

// Value The valueMap() step yields a map representation of the properties of an element.
func Value() Traversal {
	return startStep("value", "(", ")", nil)
}

// Order step is a map step that sorts returned objects given a certain criteria.
func Order() Traversal {
	return startStep("order", "(", ")", nil)
}

// ValueMap The valueMap() step yields a map representation of the properties of an element.
func ValueMap(v ...interface{}) Traversal {
	return startStep("valueMap", "(", ")", v...)
}

// Aggregate step (sideEffect) is used to aggregate all the objects at a particular point of traversal
// into a Collection. The step uses eager evaluation in that no objects continue on until
// all previous objects have been fully aggregated (as opposed to store() which lazily fills a collection).
// The eager evaluation nature is crucial in situations where everything at a particular point is required for future computation.
func Aggregate(v ...interface{}) Traversal {
	return startStep("aggregate", "(", ")", v...)
}

// Limit returns a new gremlin Traversal
func Limit(v ...interface{}) Traversal {
	return startStep("limit", "(", ")", v...)
}

// Local step is a branch step that allows the action within the step to be applied to the element within the local() step.
func Local(v ...interface{}) Traversal {
	return startStep("local", "(", ")", v...)
}

// Properties step retrieves the properties of the specified element.
func Properties(v ...interface{}) Traversal {
	return startStep("properties", "(", ")", v...)
}

// Property step in a traversal has the same beahviour as in the Graph API
func Property(v ...interface{}) Traversal {
	return startStep("property", "(", ")", v...)
}

// Coalesce step evaluates the provided traversals in order and returns the first traversal that emits at least one element..
func Coalesce(v ...interface{}) Traversal {
	return startStep("coalesce", "(", ")", v...)
}

// Choose step is a branch step that can be used to route a traverser to a particular traversal branch, similar to if-then-else logic.
func Choose(v ...interface{}) Traversal {
	return startStep("choose", "(", ")", v...)
}

// FlatMap step is a general step. This step creates a map of the Traversal objects and streams the map to the next Traversal step.
func FlatMap(v ...interface{}) Traversal {
	return startStep("flatMap", "(", ")", v...)
}

// Filter step is a general step. Map the traverser to either true or false, where false will not pass the traverser to the next step.
func Filter(v ...interface{}) Traversal {
	return startStep("filter", "{", "}", v...)
}

// Dedup returns a new gremlin Traversal
func Dedup() Traversal {
	return startStep("dedup", "(", ")", nil)
}

// Range step is a filter step that allows only the specified number of objects to pass onto the next step.
func Range(v ...interface{}) Traversal {
	return startStep("range", "(", ")", v...)
}

// Read returns a new gremlin Traversal
func Read() Traversal {
	return startStep("read", "(", ")", nil)
}

// Count returns a new gremlin Traversal
func Count(v ...interface{}) Traversal {
	return startStep("count", "(", ")", v...)
}

// Of returns a new gremlin Traversal
func Of() Traversal {
	return startStep("of", "(", ")", nil)
}

// Fold There are situations when the traversal stream needs a "barrier" to aggregate all the objects
// and emit a computation that is a function of the aggregate. The fold()-step (map) is one particular instance of this.
// Please see unfold()-step for the inverse functionality.
func Fold() Traversal {
	return startStep("fold", "(", ")", nil)
}

// Out step moves the Traversal to the outgoing adjacent vertices, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func Out(v ...interface{}) Traversal {
	return startStep("out", "(", ")", v...)
}

// Has step is a filter step. It is the most common step used for graph Traversals,
// since this step narrows the query to find particular vertices or edges with certain property values.
func Has(v ...interface{}) Traversal {
	return startStep("has", "(", ")", v...)
}

// HasLabel step is a filter step.
func HasLabel(v ...interface{}) Traversal {
	return startStep("hasLabel", "(", ")", v...)
}

// OutE step moves the Traversal to the outgoing incident edges, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func OutE(v ...interface{}) Traversal {
	return startStep("outE", "(", ")", v...)
}

// OutV The out() step moves the Traversal to the outgoing adjacent vertices,
// given the edge labels. Specifying no edge label will traverse all incident edges.
func OutV() Traversal {
	return startStep("outV", "(", ")", nil)
}

// In step moves the Traversal to the incoming adjacent vertices, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func In(v ...interface{}) Traversal {
	return startStep("in", "(", ")", v...)
}

// InE step moves the Traversal to the incoming incident edges, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func InE(v ...interface{}) Traversal {
	return startStep("inE", "(", ")", v...)
}

// InV step moves the Traversal to the incoming vertices.
func InV() Traversal {
	return startStep("inV", "(", ")", nil)
}

// Both step moves the Traversal to both the outgoing and the incoming adjacent vertices, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func Both(v ...interface{}) Traversal {
	return startStep("both", "(", ")", v...)
}

// BothE step moves the Traversal to both the outgoing and the incoming incident edges, given the edge labels.
// Specifying no edge label will traverse all incident edges.
func BothE(v ...interface{}) Traversal {
	return startStep("bothE", "(", ")", v...)
}

// BothV step moves the Traversal to both the outgoing and the incoming vertices.
func BothV() Traversal {
	return startStep("bothV", "(", ")", nil)
}

// OtherV step moves the Traversal to the vertex that was not the vertex that was moved from.
func OtherV() Traversal {
	return startStep("otherV", "(", ")", nil)
}

// Map the traverser to some object of type E for the next step to process.
func Map(v ...interface{}) Traversal {
	return startStep("map", "(", ")", v...)
}

// Match step is a map step provides a more declarative form of graph querying based on pattern matching.
func Match(v ...interface{}) Traversal {
	return startStep("match", "(", ")", v...)
}

////////////////
// System
///////////////

// Exists returns a new gremlin Traversal
func Exists() Traversal {
	return startStep("exists", "(", ")", nil)
}

// Create returns a new gremlin Traversal
func Create() Traversal {
	return startStep("create", "(", ")", nil)
}

// Drop returns a new gremlin Traversal
func Drop() Traversal {
	return startStep("drop", "(", ")", nil)
}

// Graphs discover what graphs currently exist using this command.
func Graphs() Traversal {
	return startStep("graphs", "(", ")", nil)
}

// Set returns a new gremlin Traversal
func Set(v ...interface{}) Traversal {
	return startStep("set", "(", ")", v...)
}

////////////////
// Step-modulators
///////////////

// As step is a step modulator, a helper step for another Traversal step.
func As(v ...interface{}) Traversal {
	return startStep("as", "(", ")", v...)
}

// By step is a step modulator that can be used to modify sort order from a previous step.
func By(v ...interface{}) Traversal {
	return startStep("by", "(", ")", v...)
}

// From step is a step modulator, a helper step for another Traversal step.
// Its main use is to designate the outgoing vertex for an addE() step. Generally, a to() step is paired with a from() step.
func From(v ...interface{}) Traversal {
	return startStep("from", "(", ")", v...)
}

// Or step ensures that at least one of the provided traversals yield a result (filter). Please see and() for and-semantics.
func Or(v ...interface{}) Traversal {
	return startStep("or", "(", ")", v...)
}

// Option step modulator is used in conjunction with branch()
// or choose(), and specifies the returned values based on values found at that point in a Traversal.
func Option(v ...interface{}) Traversal {
	return startStep("option", "(", ")", v...)
}

// Times step is a step modulator, a helper step for another Traversal step.
// Its main use is to repeat a repeat() step for the specified number of times.
func Times(v ...interface{}) Traversal {
	return startStep("times", "(", ")", v...)
}

// To returns a new gremlin Traversal
func To(v ...interface{}) Traversal {
	return startStep("to", "(", ")", v...)
}

// Emit step is a step modulator, a helper step for another Traversal step.
// Its main use is to emit either incoming Traversals before a repeat() step, or emit outgoing Traversals after a repeat() step.
// The emission sends a copy of the current objects to the next step in the query.
// A predicate or Traversal can be used in an emit() step to cause the emission only if the predicate or Traversal is true.
func Emit(v ...interface{}) Traversal {
	return startStep("emit", "(", ")", v...)
}

// Until step is a step modulator, a helper step for another Traversal step.
// Its main use is to turn a repeat() step into a do-while loop (if used after the repeat() step) or a while-do loop (if used before the repeat() step).
// A predicate or Traversal can be used in an until() step to cause the loop to complete only if the predicate or Traversal is true.
func Until(v ...interface{}) Traversal {
	return startStep("until", "(", ")", v...)
}

////////////////////////
// Terminal steps
////////////////////////

// Next  step is a terminal step. It returns the next number of steps, based on a supplied integer value.
func Next(v ...interface{}) Traversal {
	return startStep("next", "(", ")", v...)
}

// HasNext step is a terminal step. It determines whether or not there are available results from a traversal,
// returning a Boolean value of true or false.
func HasNext() Traversal {
	return startStep("hasNext", "(", ")", nil)
}

// ToList will return all results in a list.
func ToList() Traversal {
	return startStep("toList", "(", ")", nil)
}

// ToSet will return all results in a set and thus, duplicates removed.
func ToSet() Traversal {
	return startStep("toSet", "(", ")", nil)
}

// ToBulkSet will return all results in a weighted set and thus, duplicates preserved via weighting
func ToBulkSet() Traversal {
	return startStep("toBulkSet", "(", ")", nil)
}

// Iterate does not exactly fit the definition of a terminal step in that it doesn’t return a result,
// but still returns a traversal - it does however behave as a terminal step in that it iterates
// the traversal and generates side effects without returning the actual result.
func Iterate() Traversal {
	return startStep("iterate", "(", ")", nil)
}

////////////////
// Predicates
///////////////

// Eq predicate answers the question: Is the incoming object equal to the provided object?
func Eq(v ...interface{}) Traversal {
	return startStep("eq", "(", ")", v...)
}

// Neg predicate answers the question: Is the incoming object equal to the provided object?
func Neg(v ...interface{}) Traversal {
	return startStep("neg", "(", ")", v...)
}

// Lt predicate answers the question: Is the incoming number less than the provided number?
func Lt(v ...interface{}) Traversal {
	return startStep("lt", "(", ")", v...)
}

// Lte predicate answers the question: Is the incoming number less than or equal to the provided number?
func Lte(v ...interface{}) Traversal {
	return startStep("lte", "(", ")", v...)
}

// Gt  predicate answers the question: Is the incoming number greater than the provided number?
func Gt(v ...interface{}) Traversal {
	return startStep("gt", "(", ")", v...)
}

// Gte predicate answers the question: Is the incoming number greater than or equal to the provided number?
func Gte(v ...interface{}) Traversal {
	return startStep("gte", "(", ")", v...)
}

// Inside predicate answers the question: Is the incoming number greater than the first provided number and less than the second?
func Inside(v ...interface{}) Traversal {
	return startStep("inside", "(", ")", v...)
}

// Is it's possible to filter scalar values using is()-step (filter).
func Is(v ...interface{}) Traversal {
	return startStep("is", "(", ")", v...)
}

// Outside predicate answers the question: Is the incoming number less than the first provided number or greater than the second?
func Outside(v ...interface{}) Traversal {
	return startStep("outside", "(", ")", v...)
}

// Between predicate answers the question: Is the incoming number less than the first provided number or greater than the second?
func Between(v ...interface{}) Traversal {
	return startStep("between", "(", ")", v...)
}

// Within predicate answers the question: Is the incoming object in the array of provided objects?
func Within(v ...interface{}) Traversal {
	return startStep("within", "(", ")", v...)
}

// Without predicate answers the question: Is the incoming object not in the array of provided objects?
func Without(v ...interface{}) Traversal {
	return startStep("without", "(", ")", v...)
}
