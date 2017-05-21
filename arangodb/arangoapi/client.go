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
	}

	// Open "examples_books" database
	db, err := client.Database(nil, "person_db")
	if err != nil {
		// Handle error
	}

	var person  domain.Person
	person.Personid = rand.String(10)
	person.Age=25
	person.Firstname="John"
	person.Lastname="Doe"

	ctx := context.Background()
	col, err := db.Collection(ctx, "person_collection")
	if err != nil {
		// handle error
	}
	if err != nil {
		// Handle error
	}


	meta, err := col.CreateDocument(nil, person)
	if err != nil {
		// Handle error
	}
	println(" The Out Put is ",meta.Key)


}

func ReadRecord(){

	client, err := arrangoconnection.GetConnection()

	if err != nil {
		// Handle error
	}

	client.Database(nil, "person_db")




}

func DeleteRecord() {

}
