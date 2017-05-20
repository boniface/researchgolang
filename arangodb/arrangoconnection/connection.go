package arrangoconnection

import (
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

func GetConnection()  (driver.Client, error){
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
	})
	if err != nil {
		// Handle error
	}
	client, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
	})
	if err != nil {
		// Handle error
	}

	return client, err

}

