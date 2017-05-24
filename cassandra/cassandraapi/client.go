package cassandraapi

import (
	"github.com/gocql/gocql"
	"github.com/boniface/researchgolang/cassandra/cassandraconnection"
)

func CreateRecord()  {

	client, err := cassandraconnection.GetConnection()
	if err != nil {
		// Handle error
		println("Problem with the connection to cassandra !!")
		panic(err)
	}

	err = client.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
		"me", gocql.TimeUUID(), "hello world").Exec()
	if err != nil {
		println("Problem with the query !!")
		panic(err)
	}

}

func ReadRecord(){

	var id gocql.UUID
	var text string

	client, err := cassandraconnection.GetConnection()

	if err != nil {
		// Handle error
		println("Problem with the connection to cassandra !!")
		panic(err)
	}

	err = client.Query(`SELECT id, text FROM tweet WHERE timeline = ? LIMIT 1`,
		"me").Consistency(gocql.One).Scan(&id, &text)
	if err != nil {
		println("Problem with the query !!")
		panic(err)
	}

	println("Tweet:", id, text)

}

func ReadAllRecord(){

	var id gocql.UUID
	var text string

	client, err := cassandraconnection.GetConnection()
	if err != nil {
		// Handle error
		println("Problem with the connection to cassandra !!")
		panic(err)
	}

	iter := client.Query(`SELECT id, text FROM tweet WHERE timeline = ?`, "me").Iter()

	for iter.Scan(&id, &text) {
		println("Tweet:", id, text)
	}

	err = iter.Close()
	if err != nil {
		println("Problem with the query !!")
		panic(err)
	}

}

func DeleteRecord() {

	client, err := cassandraconnection.GetConnection()

	if err != nil {
		// Handle error
		println("Problem with the connection to cassandra !!")
		panic(err)
	}

	err = client.Query(`DELETE * FROM tweet WHERE text = ?`,
		"hello world").Exec()
	if err != nil {
		println("Problem with the query !!")
		panic(err)
	}

}