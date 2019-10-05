# gizmo
Gremlin language query builder for Go.

This is a work in progress for now.  At the moment, this package can only be used for generating Gremlin query strings.
#
### The Logic
#### Let me know if the package is missing any Traversal steps.  Which I'm sure it is.
The basic idea is to have a starting point for a gremlin query and then append a sequence of steps to the end.
#
### The Process
Lets say that we have a query like this:
```groovy
g.V().hasLabel('user')
```
The query has as a starting point the character 'g'.  Then, V() is appended, then finally hasLabel('user').
There are two ways to start a query:

```groovy
g := gizmo.NewGraph()
or
g := gizmo.G()
```
This results in a Traversal that has 'g' as the start/head of the query.

Now that we know how to start the query, we can continue with the rest of the query
```groovy
g.V().HasLabel("user)
```

You can also pass a parameter.

```groovy
g := gizmo.NewGraph("x")
or
g := gizmo.G("x")
```
Which results in a Traversal that begins with 'x'.  The character 'g' is the default when a parameter is ommitted.


Now for a more complex example.  Let's say we want to create the query bellow.

We want to create a vertex 'product' for for a particular user if it does not exist, then create an edge from the user to the product.  But if the product exists, return the product vertex.
```groovy
g.V().has('product', 'productName', 'camera').fold().
coalesce(
	unfold(), 
	addV('product').
		property('productName', 'camera').
		property('createdBy', '12345').
		property('modifiedBy', '12345').
		property('createdOn', 1570288181201).
		property('modifiedOn', 1570288181201).
		as('p').project('p').
		V().has('user', 'userId', '12345').as('u').
		addE('has').from('u').to('p').
		select('p'))
```
The Graph object generates a new Traversal with a starting point of 'g'

```groovy
g := gizmo.NewGraph()
g.V().Has("product", "productName", "camera").Fold().
	Coalesce(
		gizmo.Unfold(),
		gizmo.AddV("product").
			Property("productName", "camera").
			Property("createdBy", "12345").
			Property("modifiedBy", "12345").
			Property("createdOn", time.Now().UnixNano()/1000000).
			Property("modifiedOn", time.Now().UnixNano()/1000000).
			As("p").Project("p").
			V().Has("user", "userId", "12345").As("u").
			AddE("has").From("u").To("p").
			Select("p"),
	)
```

In this other example, we can separate the quey in two steps since the steps are always added to the end of the Traversal

```groovy
vals := []string{"one", "two"}
ids := []float32{1, 2}
g := gizmo.G()
g.V().Has("user", "userId", "1968410f-a76c-4827-a0cc-4126dcd95590").
	OutE("has").InV().HasLabel("attribute").Order().By("name")
g.Or(gizmo.Has("values", gizmo.Within(vals)), gizmo.Has("ids", gizmo.Within(ids)))
```
This results in the following query

```groovy
g.V().has('user', 'userId', '1968410f-a76c-4827-a0cc-4126dcd95590').
	outE('has').inV().hasLabel('attribute').order().by('name').
	or(has('values', within('one', 'two')), has('ids', within(1, 2)))
)
```

You can find more examples in the example folder.