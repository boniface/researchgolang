package redisapi

import (
	"fmt"
	"github.com/boniface/researchgolang/redis/redisconnection"
)

func CreateRecord() {

	client, err := redisconnection.GetConnection()
	if err != nil {
		// Handle error
		println("Problem with the connection to redis !!")
		panic(err)
	}

	err = client.Set("key", "value", 0).Err()
	if err != nil {
		println("Problem with the query !!")
		panic(err)
	}

}

func ReadRecord(){

	client, err := redisconnection.GetConnection()
	if err != nil {
		// Handle error
		println("Problem with the connection to redis !!")
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		println("Problem with the query !!")
		panic(err)
	}

	fmt.Println("\nkey:", val)

}

func ReadAllRecord(){

	client, err := redisconnection.GetConnection()
	if err != nil {
		// Handle error
		println("Problem with the connection to redis !!")
		panic(err)
	}

}

func DeleteRecord() {

	client, err := redisconnection.GetConnection()

	if err != nil {
		// Handle error
		println("Problem with the connection to redis !!")
		panic(err)
	}

}