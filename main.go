package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"github.com/boniface/researchgolang/arangodb/arangoapi"
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
	app.Get("/redis/read", arangoapi.Read)
	app.Get("/redis/write", arangoapi.Read)
	app.Get("/redis/write/:number", arangoapi.Read)
	app.Get("/redis/delete", arangoapi.Read)
	app.Get("/redis/readall", arangoapi.Read)


	// DO FOR MONGO
	app.Get("/mongo/read", arangoapi.Read)
	app.Get("/mongo/write", arangoapi.Read)
	app.Get("/mongo/write/:number", arangoapi.Read)
	app.Get("/mongo/delete", arangoapi.Read)
	app.Get("/mongo/readall", arangoapi.Read)


	// DO FOR CASSANDRA
	app.Get("/cassandra/read", arangoapi.Read)
	app.Get("/cassandra/write", arangoapi.Read)
	app.Get("/cassandra/write/:number", arangoapi.Read)
	app.Get("/cassandra/delete", arangoapi.Read)
	app.Get("/cassandra/readall", arangoapi.Read)

	app.Listen(":8080")

}
