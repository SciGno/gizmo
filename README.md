# gizmo
Gremlin language query builder for Go.

This is a work in progress for now.  At the moment, this package can only be used for generating Gemlin query strings.
#
### The Logic
#### Let me know if the package is missing any Traversal steps.  Which I'm sure it is.
The idea is to have a starting poin and then append a step, line, etc. to the end of the query.

Lets say that we have a query like this:
```groovy
g.V().out('created').hasNext()
```
The query has as a starting point the character 'g'.  Then, V() is appended, then out('created') and finally hasNext().
There are two ways to start a query in the package:
First:
```groovy
graph := gizmo.Graph()
g := graph.Traversal("g")
```
Second. Once we have a Graph object and a Traversal, we can use the Traversal object to return a brand new
```groovy
g.New("g")	# This returns a Traversal with g as the starting point
g.New()		# This returns a Travertsal with an empty starting point.
g.New("__")	# This returns a Travertsal with "__" as the starting point.
```
Now that we know how to begin the query, we can make more complex ones.
Here are some examples.

```groovy
graph := gizmo.Graph()
```
We create a Graph object

```groovy
g := graph.Traversal("g")
```
The Graph object generates a new Traversal with a starting point of 'g'

```groovy
compundQuery := g.New()	
```
We create an empty Traversal where we are later going to join all the other queries

```groovy
getUser := g.New().Raw("t=").Append(g.New("g").V().Has("user", "username", "scigno"))
```
This is a more complex query. We create a new query and append a Raw("t=") to start the query.  

if we were to create a g.New("t="), the query would begin with 't=." since the traversal adds a '.' between each function.

```groovy
newQuery := g.New().TernaryOp(
	g.New("t").HasNext(),
	g.New("t").Next(),
	g.New("g").AddV("user").Property("userId", "744be509-a1cc-466d-bb10-0bb9a376da2e").Property("username", "scigno").Next(),
)
```
The TernaryOp function is not a Gremlin step, but a helper step to allow ternary compsitions in the package. Which contains the following arguments:

```groovy
g.New("t").HasNext() 		# returns t.hasNext()

g.New("t").Next()		# returns t.Next()

g.New("g").AddV("user")...	# returns g.addV('user')...
```
Finally

```groovy
compundQuery.Append(getUser).AddLine(newQuery)
```
The Append and NewLine are also package helpers which do just what they say.  Append, appends the 'getUser' Traversal to the end of  compoundQuery and the NewLine adds a new line and then the 'newQuery' Traversal.


#

### The full example source:

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
	compundQuery := g.New()
	getUser := g.New().Raw("t=").Append(g.New("g").V().Has("user", "username", "scigno"))
	newQuery := g.New().TernaryOp(
		g.New("t").HasNext(),
		g.New("t").Next(),
		g.New("g").AddV("user").Property("userId", "744be509-a1cc-466d-bb10-0bb9a376da2e").Property("username", "scigno").Next(),
	)
	compundQuery.Append(getUser).AddLine(newQuery)
	fmt.Println(compundQuery)
}
```

Output:
```bash
t=g.V().has('user','username','scigno') 
t.hasNext() ? t.next() : g.addV('user').property('userId', '744be509-a1cc-466d-bb10-0bb9a376da2e').property('username', 'scigno').next()
```
