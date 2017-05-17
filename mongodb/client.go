package main

import (
	"fmt"
	"time"
	"gopkg.in/mgo.v2"
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
}
