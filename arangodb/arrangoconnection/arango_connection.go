package arrangoconnection

import (
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

func GetConnection()  (driver.Client, error){
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://10.47.2.1:1031"},
	})
	if err != nil {
		// Handle error
		println("Problem with the http connection !!")
		panic(err)
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
	})
	if err != nil {
		// Handle error
		println("Problem with the connection to the driver !!")
		panic(err)
	}

	return client, err

}

