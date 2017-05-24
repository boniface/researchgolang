package mongoconnection

import "gopkg.in/mgo.v2"

func GetConnection() ( *mgo.Session, error){
	const (
		Host     = "10.47.2.8:27017"
		Username = ""
		Password = ""
		Database = "NewTest"
		Collection = "TestStudent"
	)

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{Host},
		Username: Username,
		Password: Password,
		Database: Database,
	})

	if err != nil {
		println("Problem with the connection to Mongo !!")
		panic(err)
	}

	defer session.Close()

	return session, nil
}