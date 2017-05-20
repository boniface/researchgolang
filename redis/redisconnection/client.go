package redisconnection

import (
	"fmt"
	"time"
	"github.com/go-redis/redis"
)

func main () {
	client := redis.NewClient(&redis.Options{
		Addr:     "10.47.2.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	now1 := time.Now()
	nano1 := now1.UnixNano()
	//fmt.Println("\nTime before writing a key: ", now1)
	//fmt.Println(nano1)

	err := client.Set("key", "value", 0).Err()

	now2 := time.Now()
	nano2 := now2.UnixNano()
	//fmt.Println("\nTime after writing a key: ", now2)
	//fmt.Println(nano2)

	diff := nano2 - nano1
	fmt.Println("\nTime for writing a key: ", diff, " ns\n")

	if err != nil {
		panic(err)
	}

	now3 := time.Now()
	nano3 := now3.UnixNano()
	//fmt.Println("\nTime before reading a key: ", now3)
	//fmt.Println(nano3)

	val, err := client.Get("key").Result()

	now4 := time.Now()
	nano4 := now4.UnixNano()
	//fmt.Println("\nTime after reading a key: ", now4)
	//fmt.Println(nano4)

	diff2 := nano4 - nano3
	fmt.Println("\nTime for reading a key: ", diff2, " ns\n")

	if err != nil {
		panic(err)
	}

	fmt.Println("\nkey:", val)

	/*val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}*/
	// Output: key value
	// key2 does not exists
}