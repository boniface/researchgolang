package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"github.com/boniface/researchgolang/arangodb/arangoapi"
	"github.com/boniface/researchgolang/redis/redisapi"
	"github.com/boniface/researchgolang/mongodb/mongoapi"
	"github.com/boniface/researchgolang/cassandra/cassandraapi"
)


func main() {
	app := iris.New()
	// output startup banner and error logs on os.Stdout
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())


        // ARANGODB
	app.Get("/arango/read", arangoapi.Read)
	app.Get("/arango/write", arangoapi.Create)
	app.Get("/arango/write/:number", arangoapi.CreateInNumbers)
	app.Get("/arango/delete", arangoapi.Delete)
	app.Get("/arango/readall", arangoapi.ReadAll)


	// DO FOR REDIS
	app.Get("/redis/read", redisapi.Read)
	app.Get("/redis/write", redisapi.Create)
	app.Get("/redis/write/:number", redisapi.CreateInNumbers)
	app.Get("/redis/delete", redisapi.Delete)
	app.Get("/redis/readall", redisapi.ReadAll)


	// DO FOR MONGO
	app.Get("/mongo/read", mongoapi.Read)
	app.Get("/mongo/write", mongoapi.Create)
	app.Get("/mongo/write/:number", mongoapi.CreateInNumbers)
	app.Get("/mongo/delete", mongoapi.Delete)
	app.Get("/mongo/readall", mongoapi.ReadAll)


	// DO FOR CASSANDRA
	app.Get("/cassandra/read", cassandraapi.Read)
	app.Get("/cassandra/write", cassandraapi.Create)
	app.Get("/cassandra/write/:number", cassandraapi.CreateInNumbers)
	app.Get("/cassandra/delete", cassandraapi.Delete)
	app.Get("/cassandra/readall", cassandraapi.ReadAll)

	app.Listen(":8080")

}
