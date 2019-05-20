package gizmo

import (
	"fmt"
	"strconv"
	"strings"
)

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

func process(value string, t *Traversal) *Traversal {
	t.value = t.value + value + "()."
	return t
}

func processString(value string, t *Traversal, s string) *Traversal {
	t.value = t.value + value + "('" + s + "')."
	return t
}

func processInt(value string, t *Traversal, i int) *Traversal {
	t.value = t.value + value + "(" + strconv.Itoa(i) + ")."
	return t
}

func processStringSlice(value string, t *Traversal, name ...string) *Traversal {
	if len(name) == 0 {
		t.value = t.value + value + "()."
	} else {
		t.value = t.value + value + "('" + name[0] + "')."
	}
	return t
}

func processInterfaceOptionalSlice(value string, t *Traversal, i ...interface{}) *Traversal {
	if len(i) > 0 {
		if _, ok := i[0].(*Traversal); ok {
			t.value = t.value + value + "(" + i[0].(*Traversal).String() + ")."
		} else {
			t.value = t.value + value + "(" + i[0].(string) + ")."
		}
	} else {
		t.value = t.value + value + "()."
	}
	return t
}

// Match step is a map step provides a more declarative form of graph querying based on pattern matching.
func processInterfaceSlice(value string, t *Traversal, i ...interface{}) *Traversal {
	if len(i) > 0 {
		var arr []string
		for _, v := range i {
			if _, ok := v.(*Traversal); ok {
				arr = append(arr, v.(*Traversal).String())
			}
		}
		t.value = t.value + value + "(" + strings.Join(arr, ",") + ")."
	} else {
		t.value = t.value + value + "()."
	}
	return t
}

func processIntSlice(value string, t *Traversal, i ...int) *Traversal {
	if len(i) > 0 {
		t.value = t.value + value + "(" + strconv.Itoa(i[0]) + ")."
	} else {
		t.value = t.value + value + "()."
	}
	return t
}

func processInterface(value string, t *Traversal, i interface{}) *Traversal {
	if _, ok := i.(*Traversal); ok {
		t.value = t.value + value + "(" + i.(*Traversal).String() + ")."
	} else {
		t.value = t.value + value + "(" + i.(string) + ")."
	}
	return t
}

func processKeyVakues(value string, k interface{}, v interface{}, t *Traversal) *Traversal {
	if _, ok := k.(string); ok {
		k = "'" + k.(string) + "'"
	}

	if _, ok := v.(string); ok {
		v = "'" + v.(string) + "'"
	}

	t.value = t.value + fmt.Sprintf("%v(%v, %v).", value, k, v)
	return t
}

func processStringAndStringSlice(value string, t *Traversal, name string, values ...string) *Traversal {
	if len(values) == 0 {
		t.value = t.value + value + "('" + name + "')."
	} else {
		t.value = t.value + value + "('" + name + "','" + strings.Join(values, "','") + "')."
	}
	return t
}

func processFromTo(value string, t *Traversal, from int, to int) *Traversal {
	t.value = t.value + value + "(" + strconv.Itoa(from) + "," + strconv.Itoa(to) + ")."
	return t
}
