
package cassandraapi
import (
"fmt"
"log"
"time"
"github.com/gocql/gocql"
)

func main() {
	// connect to the cluster
	cluster := gocql.NewCluster("10.47.2.151", "10.47.2.8", "10.47.2.1")
	cluster.Keyspace = "example"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	now1 := time.Now()
	nano1 := now1.UnixNano()
	//fmt.Println("\nTime before writing the Tweet: ", now1)
	//fmt.Println(nano1)

	// insert a tweet
	if err := session.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
		"me", gocql.TimeUUID(), "hello world").Exec(); err != nil {
		log.Fatal(err)
	}

	now2 := time.Now()
	nano2 := now2.UnixNano()
	//fmt.Println("\nTime after writing the Tweet: ", now2)
	//fmt.Println(nano2)

	diff := nano2 - nano1
	fmt.Println("\nTime for writing a Tweet: ", diff, " ns\n")

	// insert thousand tweets
	/*i := 1

	now5 := time.Now()
	nano5 := now5.UnixNano()
	//fmt.Println("\nTime before writing Tweets: ", now5)
	//fmt.Println(nano5)

	for i <= 1000 {
		session.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
			"me", gocql.TimeUUID(), "hello world").Exec()
		i = i + 1
	}

	now6 := time.Now()
	nano6 := now6.UnixNano()
	//fmt.Println("\nTime after writing Tweets: ", now6)
	//fmt.Println(nano6)

	diff3 := nano6 - nano5
	fmt.Println("\nTime for writing thousand Tweets: ", diff3, " ns\n")*/

	var id gocql.UUID
	var text string

	/* Search for a specific set of records whose 'timeline' column matches
	 * the value 'me'. The secondary index that we created earlier will be
	 * used for optimizing the search */

	now3 := time.Now()
	nano3 := now3.UnixNano()
	//fmt.Println("\nTime before reading the Tweet: ", now3)
	//fmt.Println(nano3)

	if err := session.Query(`SELECT id, text FROM tweet WHERE timeline = ? LIMIT 1`,
		"me").Consistency(gocql.One).Scan(&id, &text); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Tweet:", id, text)

	now4 := time.Now()
	nano4 := now4.UnixNano()
	//fmt.Println("\nTime before reading the Tweet: ", now4)
	//fmt.Println(nano4)

	diff2 := nano4 - nano3
	fmt.Println("\nTime for reading a Tweet: ", diff2, " ns\n")

	// list all tweets
	iter := session.Query(`SELECT id, text FROM tweet WHERE timeline = ?`, "me").Iter()

	for iter.Scan(&id, &text) {
		fmt.Println("Tweet:", id, text)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}