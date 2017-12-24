package main

import (
    "fmt"

    goredis "github.com/go-redis/redis"
    redigo "github.com/garyburd/redigo/redis"
)

func testGoRedis() {
    fmt.Println("====== Test go-redis ======")

    //Connect
    client := goredis.NewClient(&goredis.Options{
        Addr:     "172.16.101.128:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    defer client.Close()

    pong, err := client.Ping().Result()

    if err != nil {
    	fmt.Println("goredis ping error:", err)
    	return
	} else {
		fmt.Println("ping success! pong:", pong)
	}

	//Test set
	err = client.Set("test-go-redis", "test set!", 0).Err()

	if err != nil {
		fmt.Println("set failed! error:", err)
		return
	} else {
		fmt.Println("set success!")
	}

	//Test append
	ret, err := client.Append("test-go-redis", "+").Result()

	if err != nil {
		fmt.Println("append failed! error:", err)
		return
	} else {
		fmt.Println("append success:", ret)
	}

	//Test get
	data, err := client.Get("test-go-redis").Result()

	if err != nil {
		fmt.Println("get failed! error:", err)
		return
	} else {
		fmt.Println("get success:", data)
	}

	/**********************
	 * There are lots of other commands
	 * Check "https://redis.io/commands" to get further information
	 **********************/
}

func testRedigo() {
    fmt.Println("====== Test redigo ======")

    //Connect
	client, err := redigo.Dial("tcp", "172.16.101.128:6379")

	if err != nil {
		fmt.Println("redigo dial failed! error:", err)
	}

	defer client.Close()

	//Test set
	_, err = client.Do("SET", "test-redigo", "test set!")

	if err != nil {
		fmt.Println("set failed! error:", err)
		return
	} else {
		fmt.Println("set success!")
	}

	//Test append
	reply, err := redigo.Int(client.Do("APPEND", "test-redigo", "+"))

	if err != nil {
		fmt.Println("append failed! error:", err)
		return
	} else {
		fmt.Println("append success:", reply)
	}

	//Test get
	result, err := redigo.String(client.Do("GET", "test-redigo"))

	if err != nil {
		fmt.Println("get failed! error:", err)
		return
	} else {
		fmt.Println("get success:", result)
	}
}

func main() {
	fmt.Println("====== Test redis ======")
	testGoRedis()
	testRedigo()
}
