package main

import (
	"fmt"
	"os"

	"github.com/scigno/gizmo"
)

// Result struct
type Result []struct {
	Type  string `json:"@type"`
	Value struct {
		ID struct {
			InVertex struct {
				EmailID struct {
					Type  string `json:"@type"`
					Value string `json:"@value"`
				} `json:"emailId"`
				Label string `json:"~label"`
			} `json:"~in_vertex"`
			Label   string `json:"~label"`
			LocalID struct {
				Type  string `json:"@type"`
				Value string `json:"@value"`
			} `json:"~local_id"`
			OutVertex struct {
				UserID struct {
					Type  string `json:"@type"`
					Value string `json:"@value"`
				} `json:"userId"`
				Label string `json:"~label"`
			} `json:"~out_vertex"`
		} `json:"id"`
		InV struct {
			EmailID struct {
				Type  string `json:"@type"`
				Value string `json:"@value"`
			} `json:"emailId"`
			Label string `json:"~label"`
		} `json:"inV"`
		InVLabel string `json:"inVLabel"`
		Label    string `json:"label"`
		OutV     struct {
			UserID struct {
				Type  string `json:"@type"`
				Value string `json:"@value"`
			} `json:"userId"`
			Label string `json:"~label"`
		} `json:"outV"`
		OutVLabel string `json:"outVLabel"`
	} `json:"@value"`
}

func main() {

	graph := gizmo.Graph()
	g := graph.Traversal("g")

	n := g.New().Raw("t=").Append(g.New("g").V().Has("user", "username", "scigno").String())
	n.AddLine(g.New().TernaryOp(
		g.New("t").HasNext(),
		g.New("t").Next(),
		g.New("g").AddV("user").Property("userId", "744be509-a1cc-466d-bb10-0bb9a376da2e").Property("username", "scigno").Next(),
	).String())
	fmt.Println(n)
	fmt.Println(g)

	os.Exit(0)

	// errs := make(chan error)
	// go func(chan error) {
	// 	err := <-errs
	// 	log.Fatal("Lost connection to the database: " + err.Error())
	// }(errs) // Example of connection error handling logic

	// dialer := gremgo.NewDialer(
	// 	"ws://192.168.1.161:8182",
	// 	gremgo.SetTimeout(45),
	// ) // Returns a WebSocket dialer to connect to Gremlin Server
	// g, err := gremgo.Dial(dialer, errs) // Returns a gremgo client to interact with
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// // res, err := g.Execute( // Sends a query to Gremlin Server with bindings
	// // 	// "g.V().hasLabel('user').project('username', 'userId', 'emailAddress').by(values('username')).by(values('userId')).by(out('createdEmail').values('emailAddress').fold())",
	// // 	`lluuid = java.util.UUID.randomUUID()
	// // 	g.addV('user').property('userId', lluuid).property('username', 'scigno')
	// // 	.property('createdOn', java.time.Instant.ofEpochMilli(java.time.Instant.now().toEpochMilli()))
	// // 	.property('createdBy', lluuid)
	// // 	.property('modifiedOn', java.time.Instant.ofEpochMilli(java.time.Instant.now().toEpochMilli()))
	// // 	.property('modifiedBy', lluuid)
	// // 	.next()`,
	// // 	map[string]string{"u": "user"},
	// // 	map[string]string{"g": "MARKETBIN.g"},
	// // )
	// res, err := g.Execute( // Sends a query to Gremlin Server with bindings
	// 	// `t=g.V().has('user','username','llopez')
	// 	// t.hasNext() ? t.next() : g.addV('user').property('username','llopez').next()
	// 	// `,
	// 	`
	// 	g.V().has('user', 'username', 'llopez').as('user').outE('has').inV().hasLabel('email').as('email').select('user', 'email').fold()
	// 	`,
	// 	map[string]string{"u": "user"},
	// 	map[string]string{"g": "MARKETBIN.g"},
	// )
	// if err != nil {
	// 	fmt.Printf("%+v", err)
	// 	return
	// }

	// // result := GremlinResult{}
	// data, _ := json.MarshalIndent(res.([]interface{})[0], "", " ")
	// fmt.Println(string(data))
	// // json.Unmarshal(data, &result)
	// // fmt.Println(result[2].Value.Label)

}
