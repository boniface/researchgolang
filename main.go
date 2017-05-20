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
	app.Get("/arango/write", arangoapi.Read)
	app.Get("/arango/delete", arangoapi.Read)
	app.Get("/arango/readall", arangoapi.Read)


	// DO FOR REDIS
	app.Get("/arango/read", arangoapi.Read)
	app.Get("/arango/write", arangoapi.Read)
	app.Get("/arango/delete", arangoapi.Read)
	app.Get("/arango/readall", arangoapi.Read)


	// DO FOR MONGO
	app.Get("/arango/read", arangoapi.Read)
	app.Get("/arango/write", arangoapi.Read)
	app.Get("/arango/delete", arangoapi.Read)
	app.Get("/arango/readall", arangoapi.Read)


	// DO FOR CASSANDRA
	app.Get("/arango/read", arangoapi.Read)
	app.Get("/arango/write", arangoapi.Read)
	app.Get("/arango/delete", arangoapi.Read)
	app.Get("/arango/readall", arangoapi.Read)

	app.Listen(":8080")

}
