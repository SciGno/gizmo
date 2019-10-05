package main

import (
	"fmt"
	"time"

	"github.com/scigno/gizmo"
)

func main() {

	fmt.Print("If product for user does not exist, creat it and then create an edge from the user to the product.\nIf product exists, return the product.\n")
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
	fmt.Print(g.String(), "\n\n")

	g1 := gizmo.NewGraph()
	g1.V(1).Out().Values("name")
	fmt.Print(g1.String(), "\n\n")

	v := gizmo.VAR("t")
	g2 := gizmo.G()
	g2.V().AddE("test")
	fmt.Print(v.String()+g2.String(), "\n\n")

	vals := []string{"one", "two"}
	ids := []float32{1, 2}
	g4 := gizmo.G()
	g4.V().Has("user", "userId", "1968410f-a76c-4827-a0cc-4126dcd95590").OutE("has").InV().HasLabel("attribute").Order().By("name")
	g4.Or(gizmo.Has("values", gizmo.Within(vals)), gizmo.Has("ids", gizmo.Within(ids)))
	fmt.Print(g4, "\n\n")

	g5 := gizmo.G("x")
	fmt.Print(g5, "\n\n")
}
