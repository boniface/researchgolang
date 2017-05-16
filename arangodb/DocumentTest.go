package main

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/solher/arangolite"
)

type Node struct {
	arangolite.Document
}

func main() {
	db := arangolite.New().
		LoggerOptions(false, false, false).
		Connect("http://10.47.2.151:1031", "_system", "root", "rootPassword")

	now1 := time.Now()
	nano1 := now1.UnixNano()
	//fmt.Println("\nTime before the creation of the new database: ", now1)
	//fmt.Println(nano1)

	// creation of the database
	_, _ = db.Run(&arangolite.CreateDatabase{
		Name: "boniface",
		Users: []map[string]interface{}{
			{"username": "root", "passwd": "rootPassword"},
			{"username": "user", "passwd": "password"},
		},
	})

	now2 := time.Now()
	nano2 := now2.UnixNano()
	//fmt.Println("\nTime after the creation of the database: ", now2)
	//fmt.Println(nano2)

	diff := nano2 - nano1
	fmt.Println("\nTime to create the new database: ", diff, " ns\n")

	db.SwitchDatabase("boniface").SwitchUser("user", "password")

	_, _ = db.Run(&arangolite.CreateCollection{Name: "nodes"})

	key := "48765564346"

	q := arangolite.NewQuery(`
    FOR n
    IN nodes
    FILTER n._key == %s
    RETURN n
  `, key).Cache(true).BatchSize(500) // The caching feature is unavailable prior to ArangoDB 2.7

	// The Run method returns all the query results of every batches
	// available in the cursor as a slice of byte.
	r, _ := db.Run(q)

	nodes := []Node{}
	json.Unmarshal(r, &nodes)

	// The RunAsync method returns a Result struct allowing to handle batches as they
	// are retrieved from the database.
	async, _ := db.RunAsync(q)

	nodes = []Node{}
	decoder := json.NewDecoder(async.Buffer())

	for async.HasMore() {
		batch := []Node{}
		decoder.Decode(&batch)
		nodes = append(nodes, batch...)
	}

	fmt.Printf("%v", nodes)
}

// OUTPUT EXAMPLE:
// [
//   {
//     "_id": "nodes/48765564346",
//     "_rev": "48765564346",
//     "_key": "48765564346"
//   }
// ]