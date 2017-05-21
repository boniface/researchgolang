package cassandraconnection

import "github.com/gocql/gocql"

func GetConnection() ( *gocql.Session, error){
	// connect to the cluster
	cluster := gocql.NewCluster("10.47.2.151", "10.47.2.8", "10.47.2.1")
	cluster.Keyspace = "example"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	return session, err


}
