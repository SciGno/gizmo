package main

import (
	"fmt"
	"time"

	"github.com/scigno/gizmo"
)

func main() {

	// graph := gizmo.Graph()
	// g := graph.Traversal("g")
	// compundQuery := g.New()
	// getUser := g.New().Raw("t=").Append(g.New("g").V().Has("user", "username", "scigno"))
	// newQuery := g.New().TernaryOp(
	// 	g.New("t").HasNext(),
	// 	g.New("t").Next(),
	// 	g.New("g").AddV("user").Property("userId", "744be509-a1cc-466d-bb10-0bb9a376da2e").Property("username", "scigno").Next(),
	// )
	// compundQuery.Append(getUser).AddLine(newQuery)
	// fmt.Println(compundQuery)

	graph := gizmo.Graph()
	g := graph.Traversal("g")
	query := g.New()
	getUser := g.New().Raw("u=").Append(g.New("g").V().Has("user", "userId", "7031dc5e-a555-4988-bbae-f9786bb52f61").Next())
	getToken := g.New().Raw("t=").Append(g.New("g").V().Has("accessToken", "accessTokenId", "7031dc5e-a555-4988-bbae-f9786bb52f61").Next())
	addEdge := g.New("g").V(g.Var("u")).AddE("has").To(g.Var("t")).Next()
	addToken := g.New("g").AddV("refreshToken").
		Property("refreshTokenId", "7031dc5e-a555-4988-bbae-f9786bb52f61").
		Property("refreshTokenString", "7031dc5e-a555-4988-bbae-f9786bb52f61").
		Property("createdBy", "7031dc5e-a555-4988-bbae-f9786bb52f61").
		Property("modifiedBy", "7031dc5e-a555-4988-bbae-f9786bb52f61").
		Property("createdOn", time.Now().UnixNano()/1000000).
		Property("modifiedOn", time.Now().UnixNano()/1000000)
	query.Append(addToken).AddLine(getUser).AddLine(getToken).AddLine(addEdge)
	fmt.Println(query)

	g2 := graph.Traversal("g")
	g2.V(g2.Raw("u").String())
	fmt.Println(g2)

	g4 := graph.Traversal("g")
	g4.V(g4.New().Raw("u").String())
	fmt.Println(g4)

	g3 := graph.Traversal("g")
	g3.V(g3.Var("u"))
	fmt.Println(g3)

	g5 := graph.Traversal("g")
	getFailures := g5.New().Raw("f=").V().Has("user", "username", "llopez").Properties("loginFailures").Next().Value()
	setFailures := g5.New("g").V().Has("user", "username", "llopez").Property("loginFailures", g5.Var("f+1"))
	compoundQuery := g5.New().Append(getFailures).AddLine(setFailures)
	fmt.Println(compoundQuery)

	g6 := graph.Traversal("g")
	g6.V().Choose(g.New().Values("userId", "username"), g.New().V().Has("user", "username", "scigno"))
	fmt.Println(g6)

	g7 := graph.Traversal("g")
	compundQuery := g7.New()
	gU := g7.New().Raw("t=").Append(g7.New("g").V().Has("user", "username", "llopez"))
	newQuery := g7.New().TernaryOp(
		g7.New("t").HasNext(),
		g7.New("t").Next(),
		g7.New("g").AddV("user").Property("userId", "7031dc5e-a555-4988-bbae-f9786bb52f61").
			Property("username", "llopez").
			Property("password", "llopez").
			Property("createdBy", "7031dc5e-a555-4988-bbae-f9786bb52f61").
			Property("modifiedBy", "7031dc5e-a555-4988-bbae-f9786bb52f61").
			Property("createdOn", time.Now().UnixNano()/1000000).
			Property("modifiedOn", time.Now().UnixNano()/1000000),
	)
	compundQuery.Append(gU).AddLine(newQuery)
	fmt.Println(compundQuery)
}
