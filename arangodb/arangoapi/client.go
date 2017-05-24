package arangoapi

import (
	"github.com/boniface/researchgolang/arangodb/arrangoconnection"
	"github.com/boniface/researchgolang/arangodb/domain"
	"github.com/boniface/researchgolang/util/rand"
	"golang.org/x/net/context"

)

func CreateRecord()  {
	client, err := arrangoconnection.GetConnection()
	if err != nil {
		// Handle error
		println("Problem with the connection to arango !!")
		panic(err)
	}

	// Open "examples_books" database
	db, err := client.Database(nil, "person_db")
	if err != nil {
		// Handle error
		println("Problem with the connection to the database !!")
		panic(err)
	}

	var person  domain.Person
	person.Personid = rand.String(10)
	person.Age = 25
	person.Firstname = "John"
	person.Lastname = "Doe"

	ctx := context.Background()

	col, err := db.Collection(ctx, "person_collection")
	if err != nil {
		// Handle error
		println("Problem with the connection to the collection")
		panic(err)
	}

	meta, err := col.CreateDocument(nil, person)
	if err != nil {
		// Handle error
		println("Problem with the creation of the document !!")
		panic(err)
	}

	println(" The Out Put is ", meta.Key)

}

func ReadRecord(){

	client, err := arrangoconnection.GetConnection()

	if err != nil {
		// Handle error
		println("Problem with the connection to arango !!")
		panic(err)
	}

	db, err := client.Database(nil, "person_db")
	if err != nil {
		// Handle error
		println("Problem with the connection to the database !!")
		panic(err)
	}

	ctx := context.Background()

	col, err := db.Collection(ctx, "person_collection")
	if err != nil {
		// Handle error
		println("Problem with the connection to the collection")
		panic(err)
	}

	meta, err := col.ReadDocument(nil, "John", )
	if err != nil {
		// Handle error
		println("Problem with the connection to the document !!")
		panic(err)
	}

	println("The Out Put is ", meta.Key)

}

func DeleteRecord() {

	client, err := arrangoconnection.GetConnection()

	if err != nil {
		// Handle error
		println("Problem with the connection to arango !!")
		panic(err)
	}

	db, err := client.Database(nil, "person_db")
	if err != nil {
		// Handle error
		println("Problem with the connection to the database !!")
		panic(err)
	}

	ctx := context.Background()

	col, err := db.Collection(ctx, "person_collection")
	if err != nil {
		// Handle error
		println("Problem with the connection to the collection")
		panic(err)
	}

	meta, err := col.RemoveDocument(nil, "John")
	if err != nil {
		// Handle error
		println("Problem with the connection to the document !!")
		panic(err)
	}

	println("Record deleted ", meta.Key)

}
