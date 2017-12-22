package main

import (
    "fmt"

    goredis "github.com/go-redis/redis"
    //redigo "github.com/garyburd/redigo/redis"
)

func testGoRedis() {
    fmt.Println("====== Test go-redis ======")

    //Connect
    client := goredis.NewClient(&goredis.Options{
        Addr:     "172.16.101.128:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    pong, err := client.Ping().Result()

    if err != nil {
    	fmt.Println("goredis ping error:", err)
    	return
	} else {
		fmt.Println(pong, err)
	}

	//Test set
	err = client.Set("test", "test set", 0).Err()

	if err != nil {
		fmt.Println("set failed! error:", err)
		return
	} else {
		fmt.Println("set success!")
	}

	//Test get
	data, err := client.Get("test").Result()

	if err != nil {
		fmt.Println("get failed! error:", err)
		return
	} else {
		fmt.Println("get data:", data)
	}
}

func testRedigo() {
    fmt.Println("====== Test redigo ======")
}

func main() {
	fmt.Println("====== Test redis ======")
	testGoRedis()
	testRedigo()
}
