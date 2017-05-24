package mongoapi

import (
	"fmt"
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/boniface/researchgolang/mongodb/mongoconnection"
)

type Student struct{
	Name   string    `bson:"name"`
	Age   uint8    `bson:"age"`
	ID   uint8    `bson:"id"`
}

func main() {
	const (
		Host     = "10.47.2.8:27017"
		Username = ""
		Password = ""
		Database = "NewTest"
		Collection = "TestStudent"
	)

	student := Student{
		Name:	"Anthony",
		Age:	21,
		ID:	0000,
	}

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{Host},
		Username: Username,
		Password: Password,
		Database: Database,
	})

	if err != nil {
		panic(err)
	}

	defer session.Close()

	fmt.Println("\nConnected to", session.LiveServers())

	coll := session.DB(Database).C(Collection)

	now1 := time.Now()
	nano1 := now1.UnixNano()
	//fmt.Println("\nTime before the insertion of the document: ", now1)
	//fmt.Println(nano1)

	err = coll.Insert(student)

	now2 := time.Now()
	nano2 := now2.UnixNano()
	//fmt.Println("\nTime after the insertion of the document: ", now2)
	//fmt.Println(nano2)

	diff := nano2 - nano1
	fmt.Println("\nTime to insert the document: ", diff, " ns\n")

	if err != nil {
		panic(err)
	}

	fmt.Println("Document inserted successfully!")

	stud := "Anthony"

	now3 := time.Now()
	nano3 := now3.UnixNano()
	//fmt.Println("\nTime before reading the document: ", now3)
	//fmt.Println(nano3)

	age, err := coll.Find(bson.M{"name": stud}).Count()

	now4 := time.Now()
	nano4 := now4.UnixNano()
	//fmt.Println("\nTime after reading the document: ", now4)
	//fmt.Println(nano4)

	diff2 := nano4 - nano3
	fmt.Println("\nTime to read the document: ", diff2, " ns\n")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s appears %d times.\n", stud, age)

}

func CreateRecord() {

	client, err := mongoconnection.GetConnection()
	if err != nil {
		// Handle error
		println("Problem with the connection to Mongo !!")
		panic(err)
	}

}

func ReadRecord(){

	client, err := mongoconnection.GetConnection()
	if err != nil {
		// Handle error
		println("Problem with the connection to Mongo !!")
		panic(err)
	}

}

func ReadAllRecord(){

	client, err := mongoconnection.GetConnection()
	if err != nil {
		// Handle error
		println("Problem with the connection to Mongo !!")
		panic(err)
	}

}

func DeleteRecord() {

	client, err := mongoconnection.GetConnection()

	if err != nil {
		// Handle error
		println("Problem with the connection to Mongo !!")
		panic(err)
	}

}
